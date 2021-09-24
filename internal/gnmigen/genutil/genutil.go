// Copyright 2019 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package genutil implements helpers used in the generated Go telemetry API.
package genutil

import (
	"golang.org/x/net/context"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/internal/closer"
	"github.com/pkg/errors"
	"github.com/openconfig/goyang/pkg/yang"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"github.com/openconfig/gnmi/errlist"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/internal/testbed"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

var (
	mu    sync.Mutex
	gnmis = make(map[reservation.Device]gpb.GNMIClient)
)

// DataPoint is a value of a gNMI path at a particular time.
type DataPoint struct {
	Path *gpb.Path
	// value of the data; nil means delete
	Value *gpb.TypedValue
	// time the value was updated on the device
	Timestamp time.Time
	// time the update was received by the test
	RecvTimestamp time.Time
}

// QualifiedTypeString renders a Qualified* telemetry value to a format that's
// easier to read when debugging.
func QualifiedTypeString(value interface{}, qv *QualifiedType) string {
	// Get string for value
	var valStr string
	if v, ok := value.(ygot.ValidatedGoStruct); ok && !reflect.ValueOf(v).IsNil() {
		// Display JSON for GoStructs.
		var err error
		valStr, err = ygot.EmitJSON(v, &ygot.EmitJSONConfig{
			Format: ygot.RFC7951,
			Indent: "  ",
			RFC7951Config: &ygot.RFC7951JSONConfig{
				AppendModuleName: true,
			},
		})
		if err != nil {
			valStr = fmt.Sprintf("Displaying normally as cannot marshal to JSON (%v):\n%v", err, v)
		}
		// Add a blank line for JSON output.
		valStr = "\n" + valStr
	} else {
		if qv.IsPresent() {
			valStr = fmt.Sprintf("%v (present)", value)
		} else {
			valStr = fmt.Sprintf("%v (not present)", value)
		}
	}

	// Get string for path
	pathStr, err := ygot.PathToString(qv.GetPath())
	if err != nil {
		b, _ := prototext.Marshal(qv.GetPath())
		pathStr = string(b)
	} else {
		pathStr = fmt.Sprintf("target: %v, path: %s", qv.GetPath().GetTarget(), pathStr)
	}

	return fmt.Sprintf("\nTimestamp: %v\n%s\nVal: %s\n%s\n", qv.GetTimestamp(), pathStr, valStr, qv.GetComplianceErrors())
}

// TelemetryError stores the path, value, and error string from unsuccessfully
// unmarshalling a datapoint into a YANG schema.
type TelemetryError struct {
	Path  *gpb.Path
	Value *gpb.TypedValue
	Err   error
}

func (t *TelemetryError) String() string {
	if t == nil {
		return ""
	}
	return fmt.Sprintf("Unmarshal %v into %v: %s", t.Value, t.Path, t.Err.Error())
}

// ComplianceErrors contains the compliance errors encountered from an Unmarshal operation.
type ComplianceErrors struct {
	// PathErrors are compliance errors encountered due to an invalid schema path.
	PathErrors []*TelemetryError
	// TypeErrors are compliance errors encountered due to an invalid type.
	TypeErrors []*TelemetryError
	// ValidateErrors are compliance errors encountered while doing schema
	// validation on the unmarshalled data.
	ValidateErrors []error
}

func (c *ComplianceErrors) String() string {
	if c == nil {
		return ""
	}
	var b strings.Builder
	b.WriteString("Noncompliance Errors by category:")
	b.WriteString("\nPath Noncompliance Errors:")
	if len(c.PathErrors) != 0 {
		for _, e := range c.PathErrors {
			b.WriteString("\n\t")
			b.WriteString(e.String())
		}
	} else {
		b.WriteString(" None")
	}
	b.WriteString("\nType Noncompliance Errors:")
	if len(c.TypeErrors) != 0 {
		for _, e := range c.TypeErrors {
			b.WriteString("\n\t")
			b.WriteString(e.String())
		}
	} else {
		b.WriteString(" None")
	}
	b.WriteString("\nValue Restriction Noncompliance Errors:")
	if len(c.ValidateErrors) != 0 {
		for _, e := range c.ValidateErrors {
			b.WriteString("\n\t")
			b.WriteString(e.Error())
		}
	} else {
		b.WriteString(" None")
	}
	b.WriteString("\n")
	return b.String()
}

// prettySetRequest returns a string version of a gNMI SetRequest for human
// consumption and ignores errors. Note that the output is subject to change.
// See documentation for prototext.Format.
func prettySetRequest(setRequest *gpb.SetRequest) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "SetRequest:\n%s\n", prototext.Format(setRequest))

	writePath := func(path *gpb.Path) {
		pathStr, err := ygot.PathToString(path)
		if err != nil {
			pathStr = prototext.Format(path)
		}
		fmt.Fprintf(&buf, "%s\n", pathStr)
	}

	writeVal := func(val *gpb.TypedValue) {
		switch v := val.Value.(type) {
		case *gpb.TypedValue_JsonIetfVal:
			fmt.Fprintf(&buf, "%s\n", v.JsonIetfVal)
		default:
			fmt.Fprintf(&buf, "%s\n", prototext.Format(val))
		}
	}

	for i, path := range setRequest.Delete {
		fmt.Fprintf(&buf, "-------delete path #%d------\n", i)
		writePath(path)
	}
	for i, update := range setRequest.Replace {
		fmt.Fprintf(&buf, "-------replace path/value pair #%d------\n", i)
		writePath(update.Path)
		writeVal(update.Val)
	}
	for i, update := range setRequest.Update {
		fmt.Fprintf(&buf, "-------update path/value pair #%d------\n", i)
		writePath(update.Path)
		writeVal(update.Val)
	}
	return buf.String()
}

// Unmarshal is a wrapper to ytypes.SetNode() and ygot.Validate() that
// unmarshals a given *DataPoint to its field given the parent and parent
// schema and verifies that all data conform to the schema. Any errors due to
// the unmarshal operations above are returned in a *ComplianceErrors for the
// caller to choose whether to tolerate, while other errors will cause the test
// to error out.
// The queryPath is used to trim output Path to match length of the input path.
// If node is a leaf, the full path is returned and the queryPath can be nil.
// NOTE: The datapoints are applied in order as they are in the input slice,
// *NOT* in order of their timestamps. As such, in order to correctly support
// Collect calls, the input data must be sorted in order of timestamps.
func Unmarshal(t testing.TB, data []*DataPoint, schema *ytypes.Schema, goStructName string, structPtr ygot.GoStruct,
	queryPath *gpb.Path, isLeaf bool, reverseShadowPaths bool) (*QualifiedType, error) {
	t.Helper()

	ret := &QualifiedType{
		Path: queryPath,
	}
	if len(data) == 0 {
		return ret, nil
	}

	unmarshalledData, complianceErrs, err := unmarshal(data, schema.SchemaTree[goStructName], structPtr, schema, reverseShadowPaths)
	if err != nil {
		return ret, err
	}
	if complianceErrs != nil {
		t.Logf("INFO: noncompliant data encountered while unmarshalling: %v", complianceErrs)
	}

	ret.ComplianceErrors = complianceErrs
	if len(unmarshalledData) == 0 {
		return ret, nil
	}

	path := unmarshalledData[0].Path
	if !isLeaf {
		path = proto.Clone(unmarshalledData[0].Path).(*gpb.Path)
		path.Elem = path.Elem[:len(queryPath.Elem)]
		ret.Timestamp = LatestTimestamp(unmarshalledData)
		ret.RecvTimestamp = LatestRecvTimestamp(unmarshalledData)
	} else { // leaves only have one datapoint, so get timestamp from first element
		ret.Timestamp = unmarshalledData[0].Timestamp
		ret.RecvTimestamp = unmarshalledData[0].RecvTimestamp
	}

	ret.Path = path
	ret.Present = true
	return ret, nil
}

// unmarshal unmarshals a given slice of datapoints to its field given a
// containing GoStruct and its schema and verifies that all data conform to the
// schema. The subset of datapoints that successfully unmarshalled into the given GoStruct is returned.
// NOTE: The subset of datapoints includes datapoints that are value restriction noncompliant.
// The first error slice are internal errors, while the returned
// *ComplianceError stores the compliance errors.
func unmarshal(data []*DataPoint, structSchema *yang.Entry, structPtr ygot.GoStruct, schema *ytypes.Schema, reverseShadowPaths bool) ([]*DataPoint, *ComplianceErrors, error) {
	var unmarshalledDatapoints []*DataPoint
	var pathUnmarshalErrs []*TelemetryError
	var typeUnmarshalErrs []*TelemetryError

	errs := &errlist.List{}
	if !schema.IsValid() {
		errs.Add(errors.Errorf("input schema for generated code is invalid"))
		return nil, nil, errs.Err()
	}
	for _, dp := range data {
		// 1. Check for path compliance. Note that by unmarshalling from
		// the root, we check the entire path, including the list key.
		var gcopts []ytypes.GetOrCreateNodeOpt
		if reverseShadowPaths {
			gcopts = append(gcopts, &ytypes.PreferShadowPath{})
		}
		if _, _, err := ytypes.GetOrCreateNode(schema.RootSchema(), schema.Root, dp.Path, gcopts...); err != nil {
			pathUnmarshalErrs = append(pathUnmarshalErrs, &TelemetryError{Path: dp.Path, Value: dp.Value, Err: err})
			continue
		}
		// The structSchema passed in here is assumed to be the unzipped
		// version from the generated structs file. That schema has a single
		// root entry that all top-level schemas are connected to via their
		// parent pointers. Therefore, we must remove that first element to
		// obtain the sanitized path.
		schemaPath := util.PathStringToElements(structSchema.Path())[1:]
		relPath := util.TrimGNMIPathPrefix(dp.Path, schemaPath)
		if relPath == dp.Path { // input pointer returned by TrimGNMIPathPrefix means input prefix does not match.
			errs.Add(errors.Errorf("parent schema path %v is not a prefix of the datapoints path %v", schemaPath, dp.Path))
			continue
		}

		if dp.Value == nil {
			var dopts []ytypes.DelNodeOpt
			if reverseShadowPaths {
				dopts = append(dopts, &ytypes.PreferShadowPath{})
			}
			if err := ytypes.DeleteNode(structSchema, structPtr, relPath, dopts...); err == nil {
				unmarshalledDatapoints = append(unmarshalledDatapoints, dp)
			} else {
				errs.Add(err)
			}
		} else {
			// 2. Check for type compliance (since path should already be compliant).
			sopts := []ytypes.SetNodeOpt{&ytypes.InitMissingElements{}, &ytypes.TolerateJSONInconsistencies{}}
			if reverseShadowPaths {
				sopts = append(sopts, &ytypes.PreferShadowPath{})
			}
			if err := ytypes.SetNode(structSchema, structPtr, relPath, dp.Value, sopts...); err == nil {
				unmarshalledDatapoints = append(unmarshalledDatapoints, dp)
			} else {
				typeUnmarshalErrs = append(typeUnmarshalErrs, &TelemetryError{Path: dp.Path, Value: dp.Value, Err: err})
			}
		}
	}
	// 3. Check for value (restriction) compliance.
	validateErrs := ytypes.Validate(structSchema, structPtr)
	if pathUnmarshalErrs != nil || typeUnmarshalErrs != nil || validateErrs != nil {
		return unmarshalledDatapoints, &ComplianceErrors{PathErrors: pathUnmarshalErrs, TypeErrors: typeUnmarshalErrs, ValidateErrors: validateErrs}, errs.Err()
	}
	return unmarshalledDatapoints, nil, errs.Err()
}

// Get uses GetFn to retrieve a []*DataPoint sample for a PathStruct.
// If singleLeaf is specified, which says that the path struct corresponds to a
// non-wildcarded leaf node, then it verifies that the result is no more than 1
// datapoint.
func Get(t testing.TB, n ygot.PathStruct, singleLeaf bool) ([]*DataPoint, *gpb.Path) {
	t.Helper()
	data, path, err := get(context.Background(), n, singleLeaf)
	if err != nil {
		t.Fatalf("Get(t) at path %s: %v", path, err)
	}
	return data, path
}

func get(ctx context.Context, n ygot.PathStruct, singleLeaf bool) ([]*DataPoint, *gpb.Path, error) {
	sub, path, err := subscribe(ctx, n, gpb.SubscriptionList_ONCE)
	if err != nil {
		return nil, path, errors.Wrap(err, "cannot subscribe to gNMI client")
	}
	data, err := receiveAll(sub, false, gpb.SubscriptionList_ONCE)
	if err != nil {
		return nil, path, err
	}
	if singleLeaf && len(data) > 1 {
		return nil, path, fmt.Errorf("Got %d data points from leaf node %v, want 0 or 1: %v", len(data), path, data)
	}
	var nonMatchingLines []string
	for _, dp := range data {
		if len(dp.Path.GetElem()) < len(path.GetElem()) {
			// TODO: Catch all query-noncompliant paths.
			// It's ok to keep this temporarily, as users (i.e. generated telemetry library)
			// are unaffected. They already check for unmarshal errors, and will treat these as
			// path-noncompliant datapoints.
			pathStr, err := ygot.PathToString(path)
			if err != nil {
				b, _ := prototext.Marshal(dp.Path)
				pathStr = string(b)
			}
			nonMatchingLines = append(nonMatchingLines, fmt.Sprintf("path: %s, data: %v", pathStr, dp.Value))
		}
	}
	if nonMatchingLines != nil {
		nonMatchingLines = append([]string{"query-noncompliant datapoints encountered while unmarshalling:"}, nonMatchingLines...)
		return nil, path, errors.New(strings.Join(nonMatchingLines, "\n"))
	}
	return data, path, nil
}

// QualifiedType contains to common fields and method for the generated Qualified structs.
type QualifiedType struct {
	Present          bool              // Present indicates whether the sample value is nil.
	Path             *gpb.Path         // Path is the sample's YANG path.
	Timestamp        time.Time         // Timestamp is the sample time.
	RecvTimestamp    time.Time         // Timestamp is the time the test received the sample.
	ComplianceErrors *ComplianceErrors // ComplianceErrors contains the compliance errors encountered from an Unmarshal operation.
}

// GetPath returns the YANG query path for this value.
func (q *QualifiedType) GetPath() *gpb.Path {
	if q == nil {
		return nil
	}
	return q.Path
}

// IsPresent returns whether the value is present.
func (q *QualifiedType) IsPresent() bool {
	if q == nil {
		return false
	}
	return q.Present
}

// GetTimestamp returns the latest notification timestamp.
func (q *QualifiedType) GetTimestamp() time.Time {
	if q == nil {
		return time.Time{}
	}
	return q.Timestamp
}

// GetRecvTimestamp returns the latest timestamp when notification(s) were received.
func (q *QualifiedType) GetRecvTimestamp() time.Time {
	if q == nil {
		return time.Time{}
	}
	return q.RecvTimestamp
}

// GetComplianceErrors returns the schema compliance errors encountered while unmarshalling and validating the received data.
func (q *QualifiedType) GetComplianceErrors() *ComplianceErrors {
	if q == nil {
		return nil
	}
	return q.ComplianceErrors
}

// Predicate is a func used for conditional collection.
type Predicate func(QualifiedValue) bool

// ConvertFunc converts a datapoint to the corresponding QualifiedType.
type ConvertFunc func(*DataPoint) (QualifiedValue, error)

// QualifiedValue is an interface for generated telemetry types.
type QualifiedValue interface {
	GetPath() *gpb.Path
	IsPresent() bool
	GetRecvTimestamp() time.Time
	GetTimestamp() time.Time
}

// CollectUntil uses CollectFn to retrieve a Collection sample for a PathStruct and evaluates data against the predicate.
func CollectUntil(t testing.TB, n ygot.PathStruct, duration time.Duration, converter ConvertFunc, pred Predicate) *Collection {
	t.Helper()
	coll, path, err := collect(context.Background(), n, duration, converter, pred)
	if err != nil {
		t.Fatalf("CollectUntil(t) at path %s: %v", path, err)
	}
	return coll
}

func collect(ctx context.Context, n ygot.PathStruct, duration time.Duration, converter ConvertFunc, pred Predicate) (_ *Collection, _ *gpb.Path, rerr error) {
	cancel := func() {}
	mode := gpb.SubscriptionList_ONCE
	collectEnd := time.Now().Add(duration)
	if time.Now().Before(collectEnd) {
		ctx, cancel = context.WithDeadline(ctx, collectEnd)
		// Only cancel the context in this function if there is an error;
		// otherwise it is up to the asynchronous go routine to cancel.
		defer closer.CloseVoidOnErr(&rerr, cancel)
		mode = gpb.SubscriptionList_STREAM
	}
	sub, path, err := subscribe(ctx, n, mode)
	if err != nil {
		return nil, path, errors.Wrap(err, "cannot subscribe to gNMI client")
	}

	c := &Collection{
		data: make(chan []QualifiedValue, 1),
		err:  make(chan error, 1),
		path: path,
	}

	go func() {
		defer cancel()
		data, err := receiveUntil(sub, mode, converter, pred)
		c.data <- data
		c.err <- err
	}()

	return c, path, nil
}

// receiveUntil receives gNMI notifications until predicate is true (if setZ) or subscription times out.
func receiveUntil(sub gpb.GNMI_SubscribeClient, mode gpb.SubscriptionList_Mode, converter ConvertFunc, pred Predicate) ([]QualifiedValue, error) {
	var convertedData []QualifiedValue
	for {
		data, sync, err := receive(sub, []*DataPoint{}, true)
		if err != nil {
			return convertedData, errors.Wrap(err, "error receiving gNMI response")
		}
		if mode == gpb.SubscriptionList_ONCE && sync {
			return convertedData, nil
		}
		for _, upd := range data {
			val, err := converter(upd)
			if err != nil {
				return nil, err
			}
			convertedData = append(convertedData, val)
			if pred(val) {
				return convertedData, nil
			}
		}
	}
}

// Collection represents an ongoing collection of telemetry values.
type Collection struct {
	err  chan error
	data chan []QualifiedValue
	path *gpb.Path
}

// Await waits for the collection to finish and returns the time series of results
// and a boolean indicating whether the predicate evaluated to true.
func (c *Collection) Await(t testing.TB) ([]QualifiedValue, bool) {
	t.Helper()
	err := <-c.err
	isTimeout := false
	if err != nil {
		// if the err is gRPC timeout, then the predicate was never true
		st, ok := status.FromError(errors.Cause(err))
		if ok && st.Code() == codes.DeadlineExceeded {
			isTimeout = true
		} else {
			t.Fatal(err)
		}
	}

	return <-c.data, !isTimeout
}

func batchSet(origin string, target string, customData map[string]interface{}, req *gpb.SetRequest) (*gpb.SetResponse, error) {
	dev, opts, err := resolveBatch(target, customData)
	if err != nil {
		return nil, err
	}
	return setVals(context.Background(), dev, opts, origin, target, req)
}

func resolveBatch(target string, customData map[string]interface{}) (reservation.Device, *requestOpts, error) {
	res, err := testbed.Reservation()
	if err != nil {
		return nil, nil, err
	}
	dev, err := res.Device(target)
	if err != nil {
		return nil, nil, err
	}
	opts, err := extractRequestOpts(customData)
	if err != nil {
		return dev, nil, errors.Wrapf(err, "Error extracting request options from %v", customData)
	}
	return dev, opts, err
}

// Delete creates and makes a gNMI SetRequest delete call for the given value
// for the path specified by the path struct.
func Delete(t testing.TB, n ygot.PathStruct) *gpb.SetResponse {
	t.Helper()
	resp, path, err := set(context.Background(), n, nil, deletePath)
	if err != nil {
		t.Fatalf("Delete(t) at path %s: %v", path, err)
	}
	return resp
}

// Replace creates and makes a gNMI SetRequest replace call for the given value
// for the path specified by the path struct.
func Replace(t testing.TB, n ygot.PathStruct, val interface{}) *gpb.SetResponse {
	t.Helper()
	resp, path, err := set(context.Background(), n, val, replacePath)
	if err != nil {
		t.Fatalf("Replace(t, %v) at path %s: %v", val, path, err)
	}
	return resp
}

// Update creates and makes a gNMI SetRequest update call for the given value
// for the path specified by the path struct.
func Update(t testing.TB, n ygot.PathStruct, val interface{}) *gpb.SetResponse {
	t.Helper()
	resp, path, err := set(context.Background(), n, val, updatePath)
	if err != nil {
		t.Fatalf("Update(t, %v) at path %s: %v", val, path, err)
	}
	return resp
}

// set configures the target with the input SetRequest. The target should be
// specified in the req.Prefix.Target field of the SetRequest; this field will
// be erased before the request is forwarded to the target in the gNMI call.
func set(ctx context.Context, n ygot.PathStruct, val interface{}, op setOperation) (*gpb.SetResponse, *gpb.Path, error) {
	path, dev, opts, err := resolve(n)
	if err != nil {
		return nil, path, err
	}
	req := &gpb.SetRequest{}
	if err := populateSetRequest(req, path, val, op); err != nil {
		return nil, nil, err
	}
	response, err := setVals(ctx, dev, opts, path.GetOrigin(), path.GetTarget(), req)
	return response, path, err
}

func setVals(ctx context.Context, dev reservation.Device, opts *requestOpts, origin, target string, req *gpb.SetRequest) (*gpb.SetResponse, error) {
	// TODO: Is there any value in setting the target here?
	req.Prefix = &gpb.Path{Origin: origin}

	dut, ok := dev.(*reservation.DUT)
	if !ok {
		return nil, errors.Errorf("gNMI set cannot be called on ATEs: %v", dev)
	}
	ctx = metadata.NewOutgoingContext(ctx, opts.md)
	gnmi, err := fetchGNMI(ctx, dut, opts.client)
	if err != nil {
		return nil, err
	}

	log.V(1).Info(prettySetRequest(req))
	resp, err := gnmi.Set(ctx, req)
	log.V(1).Infof("SetResponse:\n%s", prototext.Format(resp))
	if err != nil {
		return nil, fmt.Errorf("SetRequest unsuccessful: %w", err)
	}
	return resp, nil
}

// resolve resolves a path struct to a path, device, and request options.
func resolve(n ygot.PathStruct) (*gpb.Path, reservation.Device, *requestOpts, error) {
	path, customData, errs := ygot.ResolvePath(n)
	if len(errs) > 0 {
		return nil, nil, nil, errors.Errorf("Errors resolving path struct %v: %v", n, errs)
	}
	// All paths that don't start with "meta" must be OC paths.
	if len(path.GetElem()) == 0 || path.GetElem()[0].GetName() != "meta" {
		path.Origin = "openconfig"
	}
	res, err := testbed.Reservation()
	if err != nil {
		return path, nil, nil, err
	}
	dev, err := res.Device(path.GetTarget())
	if err != nil {
		return path, nil, nil, err
	}
	opts, err := extractRequestOpts(customData)
	if err != nil {
		return path, dev, nil, errors.Wrapf(err, "Error extracting request options from %v", customData)
	}
	return path, dev, opts, err
}

// LatestTimestamp returns the latest timestamp of the input datapoints.
// If datapoints is empty, then the zero time is returned.
func LatestTimestamp(data []*DataPoint) time.Time {
	var latest time.Time
	for _, dp := range data {
		if ts := dp.Timestamp; ts.After(latest) {
			latest = ts
		}
	}
	return latest
}

// LatestRecvTimestamp returns the latest recv timestamp of the input datapoints.
// If datapoints is empty, then the zero time is returned.
func LatestRecvTimestamp(data []*DataPoint) time.Time {
	var latest time.Time
	for _, dp := range data {
		if ts := dp.RecvTimestamp; ts.After(latest) {
			latest = ts
		}
	}
	return latest
}

// BundleDatapoints splits the incoming datapoints into groups where each has
// the same common prefix path for the given query path (which may contain
// wildcards). A slice of sorted prefixes is returned so users can examine each
// group deterministically.
// The leaf input specifies that no datapoint group should have more than 1 datapoint.
// A fatal error occurs if:
// - Any path is shorter than the prefixLen.
// - leaf is true, and a particular datapoint group (i.e. common-prefix group)
//   contains more than 1 path.
func BundleDatapoints(t testing.TB, datapoints []*DataPoint, prefixLen uint, leaf bool) (map[string][]*DataPoint, []string) {
	t.Helper()
	groups, prefixes, err := bundleDatapoints(datapoints, prefixLen, leaf)
	if err != nil {
		t.Fatal(err)
	}
	return groups, prefixes
}

func bundleDatapoints(datapoints []*DataPoint, prefixLen uint, leaf bool) (map[string][]*DataPoint, []string, error) {
	groups := map[string][]*DataPoint{}
	var violations []string

	// TODO: Add fatal check for duplicate paths, as they're not allowed by GET semantics.
	for _, dp := range datapoints {
		elems := dp.Path.GetElem()
		if uint(len(elems)) < prefixLen {
			violations = append(violations, fmt.Sprintf("invalid datapoint path elems (len < %d): %v", prefixLen, elems))
			continue
		}
		prefixPath, err := ygot.PathToString(&gpb.Path{Elem: elems[:prefixLen]})
		if err != nil {
			violations = append(violations, err.Error())
			continue
		}
		groups[prefixPath] = append(groups[prefixPath], dp)
	}

	var prefixes []string
	for prefix := range groups {
		prefixes = append(prefixes, prefix)
	}
	sort.Strings(prefixes)

	if leaf {
		for prefix, dps := range groups {
			if len(dps) != 1 {
				violations = append(violations, fmt.Sprintf("got more than 1 path for leaf node %s: %v", prefix, dps))
			}
		}
	}

	if len(violations) > 0 {
		return nil, nil, fmt.Errorf("%d violations found:\n%s", len(violations), strings.Join(violations, "\n"))
	}
	return groups, prefixes, nil
}

// NewGNMI creates a new gNMI client for the specified Device.
func NewGNMI(ctx context.Context, dev reservation.Device) (gpb.GNMIClient, error) {
	switch d := dev.(type) {
	case *reservation.DUT:
		c, err := binding.Get().DialGNMI(ctx, d,
			grpc.WithBlock(),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)))
		if err != nil {
			return nil, err
		}
		return c, nil
	case *reservation.ATE:
		c, err := binding.Get().DialATEGNMI(ctx, d, grpc.WithBlock())
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	return nil, fmt.Errorf("unsupported device type: %T", dev)
}

// fetchGNMI fetches the gNMI client for the given device.
// If a GNMIClient is provided it will be just returned as is and not cached.
func fetchGNMI(ctx context.Context, dev reservation.Device, c gpb.GNMIClient) (gpb.GNMIClient, error) {
	if c != nil {
		return c, nil
	}
	mu.Lock()
	defer mu.Unlock()
	c, ok := gnmis[dev]
	if !ok {
		var err error
		c, err = NewGNMI(ctx, dev)
		if err != nil {
			return nil, err
		}
		gnmis[dev] = c
	}
	return c, nil
}

func subscribe(ctx context.Context, n ygot.PathStruct, mode gpb.SubscriptionList_Mode) (_ gpb.GNMI_SubscribeClient, _ *gpb.Path, rerr error) {
	path, dev, opts, err := resolve(n)
	if err != nil {
		return nil, path, err
	}
	ctx = metadata.NewOutgoingContext(ctx, opts.md)
	gnmi, err := fetchGNMI(ctx, dev, opts.client)
	if err != nil {
		return nil, path, err
	}
	sub, err := gnmi.Subscribe(ctx)
	if err != nil {
		return nil, path, errors.Wrap(err, "gNMI failed to Subscribe")
	}
	defer closer.Close(&rerr, sub.CloseSend, "error closing gNMI send stream")
	sr := &gpb.SubscribeRequest{
		Request: &gpb.SubscribeRequest_Subscribe{
			Subscribe: &gpb.SubscriptionList{
				Prefix: &gpb.Path{
					Origin: path.GetOrigin(),
					Target: dev.Dimensions().Name,
				},
				Subscription: []*gpb.Subscription{{
					Path: &gpb.Path{Elem: path.GetElem()},
					Mode: opts.subMode,
				}},
				Mode:     mode,
				Encoding: gpb.Encoding_PROTO,
			},
		},
	}
	if err := sub.Send(sr); err != nil {
		return nil, path, errors.Wrapf(err, "gNMI failed to Send(%+v)", sr)
	}
	// Use the target only for the subscription but exclude from the datapoint construction.
	path.Target = ""
	return sub, path, nil
}

// pathElemSlicesEqual compares whether two PathElem slices are equal.
func pathElemSlicesEqual(a, b []*gpb.PathElem) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !util.PathElemsEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

// receiveAll receives data until the context deadline is reached, or when in
// ONCE mode, a sync response is received.
func receiveAll(sub gpb.GNMI_SubscribeClient, deletesExpected bool, mode gpb.SubscriptionList_Mode) (data []*DataPoint, err error) {
	for {
		var sync bool
		data, sync, err = receive(sub, data, deletesExpected)
		if err != nil {
			// DeadlineExceeded is expected when collections are complete.
			if st, ok := status.FromError(err); ok && st.Code() == codes.DeadlineExceeded {
				break
			}
			return nil, errors.Wrapf(err, "error receiving gNMI response")
		}
		if mode == gpb.SubscriptionList_ONCE && sync {
			break
		}
	}
	return data, nil
}

// receive processes a single response from the subscription stream. If an "update" response is
// received, those points are appended to the given data and the result of that concatenation is
// the first return value, and the second return value is false. If a "sync" response is received,
// the data is returned as-is and the second return value is true. If Delete paths are present in
// the update, they are appended to the given data before the Update values. If deletesExpected
// is false, however, any deletes received will cause an error.
func receive(sub gpb.GNMI_SubscribeClient, data []*DataPoint, deletesExpected bool) ([]*DataPoint, bool, error) {
	res, err := sub.Recv()
	if err != nil {
		return data, false, err
	}
	recvTS := time.Now()

	switch v := res.Response.(type) {
	case *gpb.SubscribeResponse_Update:
		n := v.Update
		if !deletesExpected && len(n.Delete) != 0 {
			return data, false, errors.Errorf("unexpected delete updates: %v", n.Delete)
		}
		ts := time.Unix(0, n.GetTimestamp())
		newDataPoint := func(p *gpb.Path, val *gpb.TypedValue) (*DataPoint, error) {
			j, err := util.JoinPaths(n.GetPrefix(), p)
			if err != nil {
				return nil, err
			}
			// Use the target only for the subscription but exclude from the datapoint construction.
			j.Target = ""
			return &DataPoint{Path: j, Value: val, Timestamp: ts, RecvTimestamp: recvTS}, nil
		}

		// Append delete data before the update values -- per gNMI spec, they
		// should always be processed first if both update types exist in the
		// same notification.
		for _, d := range n.Delete {
			dp, err := newDataPoint(d, nil)
			if err != nil {
				return data, false, err
			}
			data = append(data, dp)
		}
		for _, u := range n.GetUpdate() {
			if u.Path == nil {
				return data, false, errors.Errorf("invalid nil path in update: %v", u)
			}
			if u.Val == nil {
				return data, false, errors.Errorf("invalid nil Val in update: %v", u)
			}
			dp, err := newDataPoint(u.Path, u.Val)
			if err != nil {
				return data, false, err
			}
			data = append(data, dp)
		}
		return data, false, nil
	case *gpb.SubscribeResponse_SyncResponse:
		return data, true, nil
	default:
		return data, false, errors.Errorf("unexpected response: %v (%T)", v, v)
	}
}
