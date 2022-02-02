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
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/internal/closer"
	"github.com/pkg/errors"
	"github.com/openconfig/goyang/pkg/yang"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"github.com/openconfig/gnmi/errlist"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/testbed"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
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
	// Sync indicates whether the received datapoint was gNMI sync response.
	Sync bool
}

func (d *DataPoint) String() string {
	return fmt.Sprintf(`Value: %s
Timestamp: %v
RecvTimestamp: %v
Path: %s`, prototext.Format(d.Value), d.Timestamp, d.RecvTimestamp, prototext.Format(d.Path))
}

// DefaultClientKey is the key for the default client in the customData map.
const DefaultClientKey = "defaultclient"

// QualifiedTypeString renders a Qualified* telemetry value to a format that's
// easier to read when debugging.
func QualifiedTypeString(value interface{}, md *Metadata) string {
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
		// TODO: Decide if value presence should be inferred by checking zero value
		if val := reflect.ValueOf(value); val.IsValid() && !val.IsZero() {
			valStr = fmt.Sprintf("%v (present)", value)
		} else {
			valStr = fmt.Sprintf("%v (not present)", value)
		}
	}

	// Get string for path
	pathStr, err := ygot.PathToString(md.Path)
	if err != nil {
		b, _ := prototext.Marshal(md.Path)
		pathStr = string(b)
	} else {
		pathStr = fmt.Sprintf("target: %v, path: %s", md.Path.GetTarget(), pathStr)
	}

	return fmt.Sprintf("\nTimestamp: %v\n%s\nVal: %s\n%s\n", md.Timestamp, pathStr, valStr, md.ComplianceErrors)
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

// MustUnmarshal calls Unmarshal and fails the calling test fatally on error.
func MustUnmarshal(t testing.TB, data []*DataPoint, schema *ytypes.Schema, goStructName string, structPtr ygot.GoStruct,
	queryPath *gpb.Path, isLeaf bool, reverseShadowPaths bool) (*Metadata, bool) {

	t.Helper()
	md, ok, err := Unmarshal(data, schema, goStructName, structPtr, queryPath, isLeaf, reverseShadowPaths)
	if err != nil {
		t.Fatal(err)
	}
	return md, ok
}

// Unmarshal is a wrapper to ytypes.SetNode() and ygot.Validate() that
// unmarshals a given *DataPoint to its field given the parent and parent
// schema and verifies that all data conform to the schema. Any errors due to
// the unmarshal operations above are returned in a *ComplianceErrors for the
// caller to choose whether to tolerate, while other errors will cause the test
// to error out. The returned bool indicates whether at least data was unmarshaled
// successfully.
// The queryPath is used to trim output Path to match length of the input path.
// If node is a leaf, the full path is returned and the queryPath can be nil.
// NOTE: The datapoints are applied in order as they are in the input slice,
// *NOT* in order of their timestamps. As such, in order to correctly support
// Collect calls, the input data must be sorted in order of timestamps.
func Unmarshal(data []*DataPoint, schema *ytypes.Schema, goStructName string, structPtr ygot.GoStruct,
	queryPath *gpb.Path, isLeaf bool, reverseShadowPaths bool) (*Metadata, bool, error) {

	ret := &Metadata{
		Path: queryPath,
	}
	if len(data) == 0 {
		return ret, false, nil
	}

	unmarshalledData, complianceErrs, err := unmarshal(data, schema.SchemaTree[goStructName], structPtr, queryPath, schema, isLeaf, reverseShadowPaths)
	ret.ComplianceErrors = complianceErrs
	if err != nil {
		return ret, false, err
	}
	if len(unmarshalledData) == 0 {
		return ret, false, nil
	}

	// Received path will be different for wildcard queries.
	path := unmarshalledData[0].Path
	if !isLeaf {
		path = proto.Clone(unmarshalledData[0].Path).(*gpb.Path)
		path.Elem = path.Elem[:len(queryPath.Elem)]
		ret.Timestamp = LatestTimestamp(unmarshalledData)
		ret.RecvTimestamp = LatestRecvTimestamp(unmarshalledData)
	} else {
		ret.Timestamp = unmarshalledData[0].Timestamp
		ret.RecvTimestamp = unmarshalledData[0].RecvTimestamp
	}
	ret.Path = path

	return ret, true, nil
}

// pathToString returns a string version of the input path for display during
// debugging.
func pathToString(path *gpb.Path) string {
	pathStr, err := ygot.PathToString(path)
	if err != nil {
		// Use Sprint instead of prototext.Format to avoid newlines.
		pathStr = fmt.Sprint(path)
	}
	return pathStr
}

// PathStructToString returns a string representing the path struct.
// Note: the output may contain an error message or invalid path;
// do not use this func outside of the generated code.
func PathStructToString(ps ygot.PathStruct) string {
	p, _, err := ygot.ResolvePath(ps)
	if err != nil {
		return fmt.Sprintf("unknown path: %v", err)
	}
	return pathToString(p)
}

// unmarshal unmarshals a given slice of datapoints to its field given a
// containing GoStruct and its schema and verifies that all data conform to the
// schema. The subset of datapoints that successfully unmarshalled into the given GoStruct is returned.
// NOTE: The subset of datapoints includes datapoints that are value restriction noncompliant.
// The second error slice are internal errors, while the returned
// *ComplianceError stores the compliance errors.
func unmarshal(data []*DataPoint, structSchema *yang.Entry, structPtr ygot.GoStruct, queryPath *gpb.Path, schema *ytypes.Schema, isLeaf, reverseShadowPaths bool) ([]*DataPoint, *ComplianceErrors, error) {
	queryPathStr := pathToString(queryPath)
	if isLeaf {
		switch {
		case len(data) > 2:
			return nil, &ComplianceErrors{PathErrors: []*TelemetryError{{
				Err: fmt.Errorf("got multiple (%d) data points for leaf node at path %s: %v", len(data), queryPathStr, data),
			}}}, nil
		case len(data) == 2 && data[1].Sync == false:
			return nil, &ComplianceErrors{PathErrors: []*TelemetryError{{
				Err: fmt.Errorf("got multiple (%d) data points for leaf node at path %s: %v", len(data), queryPathStr, data),
			}}}, nil
		}
	}

	var unmarshalledDatapoints []*DataPoint
	var pathUnmarshalErrs []*TelemetryError
	var typeUnmarshalErrs []*TelemetryError

	errs := &errlist.List{}
	if !schema.IsValid() {
		errs.Add(errors.Errorf("input schema for generated code is invalid"))
		return nil, nil, errs.Err()
	}
	// TODO: Add fatal check for duplicate paths, as they're not allowed by GET semantics.
	for _, dp := range data {
		var gcopts []ytypes.GetOrCreateNodeOpt
		if reverseShadowPaths {
			gcopts = append(gcopts, &ytypes.PreferShadowPath{})
		}
		// Sync datapoints don't contain path nor values.
		if dp.Sync {
			continue
		}

		// 1a. Check for path compliance by doing a prefix-match, since
		// the given datapoint must be a descendant of the query.
		if !util.PathMatchesQuery(dp.Path, queryPath) {
			var pathErr error
			dpPathStr := pathToString(dp.Path)
			switch {
			case len(dp.Path.Elem) == 0 && len(dp.Path.Element) > 0:
				pathErr = errors.Errorf("datapoint path uses deprecated and unsupported Element field: %s", prototext.Format(dp.Path))
			default:
				pathErr = errors.Errorf("datapoint path %q (value %v) does not match the query path %q", dpPathStr, dp.Value, queryPathStr)
			}
			pathUnmarshalErrs = append(pathUnmarshalErrs, &TelemetryError{
				Path:  dp.Path,
				Value: dp.Value,
				Err:   pathErr,
			})
			continue
		}
		// 1b. Check for path compliance: by unmarshalling from the
		// root, we check that the path, including the list key,
		// corresponds to an actual schema element.
		if _, _, err := ytypes.GetOrCreateNode(schema.RootSchema(), schema.Root, dp.Path, gcopts...); err != nil {
			pathUnmarshalErrs = append(pathUnmarshalErrs, &TelemetryError{Path: dp.Path, Value: dp.Value, Err: err})
			continue
		}
		// The structSchema passed in here is assumed to be the unzipped
		// version from the generated structs file. That schema has a single
		// root entry that all top-level schemas are connected to via their
		// parent pointers. Therefore, we must remove that first element to
		// obtain the sanitized path.
		relPath := util.TrimGNMIPathPrefix(dp.Path, util.PathStringToElements(structSchema.Path())[1:])
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
			sopts := []ytypes.SetNodeOpt{&ytypes.InitMissingElements{}, &ytypes.TolerateJSONInconsistencies{}}
			if reverseShadowPaths {
				sopts = append(sopts, &ytypes.PreferShadowPath{})
			}
			// 2. Check for type compliance (since path should already be compliant).
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

// MustGet calls Get and fails the calling test fatally on error.
func MustGet(t testing.TB, n ygot.PathStruct, subPaths ...*gpb.Path) ([]*DataPoint, *gpb.Path) {
	t.Helper()
	data, path, err := Get(context.Background(), n, subPaths...)
	if err != nil {
		t.Fatalf("Get(t) at path %s: %v", path, err)
	}
	return data, path
}

// Get does gNMI ONCE subscription for the device under n. SubPaths, if set, override the subscription paths.
func Get(ctx context.Context, n ygot.PathStruct, subPaths ...*gpb.Path) ([]*DataPoint, *gpb.Path, error) {
	sub, path, err := subscribe(ctx, n, subPaths, gpb.SubscriptionList_ONCE)
	if err != nil {
		return nil, path, errors.Wrap(err, "cannot subscribe to gNMI client")
	}
	data, err := receiveAll(sub, false, gpb.SubscriptionList_ONCE)
	if err != nil {
		return nil, path, err
	}
	return data, path, nil
}

// Metadata contains to common fields and method for the generated Qualified structs.
type Metadata struct {
	Path             *gpb.Path         // Path is the sample's YANG path.
	Timestamp        time.Time         // Timestamp is the sample time.
	RecvTimestamp    time.Time         // Timestamp is the time the test received the sample.
	ComplianceErrors *ComplianceErrors // ComplianceErrors contains the compliance errors encountered from an Unmarshal operation.
}

// GetPath returns the YANG query path for this value.
func (q *Metadata) GetPath() *gpb.Path {
	return q.Path
}

// GetTimestamp returns the latest notification timestamp.
func (q *Metadata) GetTimestamp() time.Time {
	return q.Timestamp
}

// GetRecvTimestamp returns the latest timestamp when notification(s) were received.
func (q *Metadata) GetRecvTimestamp() time.Time {
	return q.RecvTimestamp
}

// GetComplianceErrors returns the schema compliance errors encountered while unmarshalling and validating the received data.
func (q *Metadata) GetComplianceErrors() *ComplianceErrors {
	return q.ComplianceErrors
}

// QualifiedValue is an interface for generated telemetry types.
type QualifiedValue interface {
	GetPath() *gpb.Path
	GetRecvTimestamp() time.Time
	GetTimestamp() time.Time
	GetComplianceErrors() *ComplianceErrors
}

// Predicate is a func used for conditional collection.
type Predicate func(QualifiedValue) bool

// ConvertFunc converts a datapoint queried from a path to the corresponding
// Metadata.
type ConvertFunc func([]*DataPoint, *gpb.Path) (QualifiedValue, error)

// MustWatch retrieves a Collection sample for a PathStruct and evaluates data against the predicate.
func MustWatch(t testing.TB, n ygot.PathStruct, paths []*gpb.Path, duration time.Duration, isLeaf bool, converter ConvertFunc, pred Predicate) *Watcher {
	t.Helper()
	w, path, err := watch(context.Background(), n, paths, duration, isLeaf, converter, pred)
	if err != nil {
		t.Fatalf("Watch(t) at path %s: %v", path, err)
	}
	return w
}

// watch starts a gNMI subscription for the provided duration. Specifying subPaths is optional, if unset will subscribe to the path at n.
// Note: For leaves the converter and predicate are evaluated once per DataPoint. For non-leaves, they are evaluated once per notification,
// after the first sync is received.
func watch(ctx context.Context, n ygot.PathStruct, paths []*gpb.Path, duration time.Duration, isLeaf bool, converter ConvertFunc, pred Predicate) (_ *Watcher, _ *gpb.Path, rerr error) {
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
	sub, path, err := subscribe(ctx, n, paths, mode)
	if err != nil {
		return nil, path, errors.Wrap(err, "cannot subscribe to gNMI client")
	}

	c := &Watcher{
		err:  make(chan error, 1),
		path: path,
	}

	go func() {
		defer cancel()
		err := receiveUntil(sub, mode, path, isLeaf, converter, pred)
		c.err <- err
	}()

	return c, path, nil
}

// receiveUntil receives gNMI notifications until predicate is true (if set) or subscription times out.
// Note: For leaves the converter and predicate are evaluated once per DataPoint. For non-leaves,
// they are evaluated once per notification, after the first sync is received.
func receiveUntil(sub gpb.GNMI_SubscribeClient, mode gpb.SubscriptionList_Mode, path *gpb.Path, isLeaf bool, converter ConvertFunc, pred Predicate) error {
	var recvData []*DataPoint
	var hasSynced bool
	var sync bool
	var err error

	for {
		recvData, sync, err = receive(sub, recvData, true)
		if err != nil {
			return errors.Wrap(err, "error receiving gNMI response")
		}
		if mode == gpb.SubscriptionList_ONCE && sync {
			return nil
		}
		firstSync := !hasSynced && (sync || isLeaf)
		hasSynced = hasSynced || sync || isLeaf
		// Skip conversion and predicate until first sync for non-leaves.
		if !hasSynced {
			continue
		}
		var datas [][]*DataPoint
		if isLeaf {
			for _, datum := range recvData {
				// Only add a sync datapoint on the first sync, if there are no other values.
				if (len(recvData) == 1 && firstSync) || !datum.Sync {
					datas = append(datas, []*DataPoint{datum})
				}
			}
		} else {
			datas = [][]*DataPoint{recvData}
		}
		for _, data := range datas {
			val, err := converter(data, path)
			if err != nil {
				return err
			}
			if complianceErrs := val.GetComplianceErrors(); complianceErrs != nil {
				log.V(0).Infof("noncompliant data encountered during receiveUntil, ignoring value: %v", complianceErrs)
				continue
			}
			if pred(val) {
				return nil
			}
		}
		recvData = nil
	}
}

// Watcher represents an ongoing watch of telemetry values.
type Watcher struct {
	err  chan error
	path *gpb.Path
}

// Await waits for the watch to finish and returns a boolean indicating whether the predicate evaluated to true.
func (c *Watcher) Await(t testing.TB) bool {
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
	return !isTimeout
}

func batchSet(ctx context.Context, origin string, target string, customData map[string]interface{}, req *gpb.SetRequest) (*gpb.SetResponse, error) {
	dev, opts, err := resolveBatch(ctx, target, customData)
	if err != nil {
		return nil, err
	}
	return setVals(ctx, dev, opts, origin, target, req)
}

func resolveBatch(ctx context.Context, target string, customData map[string]interface{}) (binding.Device, *requestOpts, error) {
	res, err := testbed.Reservation()
	if err != nil {
		return nil, nil, err
	}
	dev, err := testbed.Device(res, target)
	if err != nil {
		return nil, nil, err
	}
	opts, err := extractRequestOpts(customData)
	if err != nil {
		return dev, nil, errors.Wrapf(err, "Error extracting request options from %v", customData)
	}
	if opts.client != nil {
		return dev, opts, nil
	}

	dc, ok := customData[DefaultClientKey]
	if !ok {
		return dev, opts, fmt.Errorf("gnmi client getter not set on root object")
	}
	client, ok := dc.(func(context.Context) (gpb.GNMIClient, error))
	if !ok {
		return dev, opts, fmt.Errorf("unexpected gnmi client getter type")
	}
	opts.client, err = client(ctx)
	if err != nil {
		return nil, nil, err
	}
	return dev, opts, nil
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
	path, dev, opts, err := resolve(ctx, n)
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

func setVals(ctx context.Context, dev binding.Device, opts *requestOpts, origin, target string, req *gpb.SetRequest) (*gpb.SetResponse, error) {
	// TODO: Is there any value in setting the target here?
	req.Prefix = &gpb.Path{Origin: origin}
	ctx = metadata.NewOutgoingContext(ctx, opts.md)

	log.V(1).Info(prettySetRequest(req))
	resp, err := opts.client.Set(ctx, req)
	log.V(1).Infof("SetResponse:\n%s", prototext.Format(resp))
	if err != nil {
		return nil, fmt.Errorf("SetRequest unsuccessful: %w", err)
	}
	return resp, nil
}

// ResolvePath resolves a path struct to a path and request options.
func ResolvePath(n ygot.PathStruct) (*gpb.Path, map[string]interface{}, error) {
	path, customData, errs := ygot.ResolvePath(n)
	if len(errs) > 0 {
		return nil, nil, errors.Errorf("Errors resolving path struct %v: %v", n, errs)
	}
	// All paths that don't start with "meta" must be OC paths.
	if len(path.GetElem()) == 0 || path.GetElem()[0].GetName() != "meta" {
		path.Origin = "openconfig"
	}
	return path, customData, nil
}

// resolve resolves a path struct to a path, device, and request options.
// The returned requestOpts contains the gnmi Client to use.
func resolve(ctx context.Context, n ygot.PathStruct) (*gpb.Path, binding.Device, *requestOpts, error) {
	path, customData, err := ResolvePath(n)
	if err != nil {
		return path, nil, nil, err
	}
	dev, opts, err := resolveBatch(ctx, path.GetTarget(), customData)
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

// BundleDatapoints splits the incoming datapoints into common-prefix groups.
//
// Each bundle is identified by a common prefix path of length prefixLen. A
// slice of sorted prefixes is returned so users can examine each group
// deterministically. If any path is longer than prefixLen, then it is stored
// in a special "/" bundle.
func BundleDatapoints(t testing.TB, datapoints []*DataPoint, prefixLen uint) (map[string][]*DataPoint, []string) {
	t.Helper()
	groups, prefixes, err := bundleDatapoints(datapoints, prefixLen)
	if err != nil {
		t.Fatal(err)
	}
	return groups, prefixes
}

func bundleDatapoints(datapoints []*DataPoint, prefixLen uint) (map[string][]*DataPoint, []string, error) {
	groups := map[string][]*DataPoint{}

	for _, dp := range datapoints {
		elems := dp.Path.GetElem()
		if uint(len(elems)) < prefixLen {
			groups["/"] = append(groups["/"], dp)
			continue
		}
		prefixPath, err := ygot.PathToString(&gpb.Path{Elem: elems[:prefixLen]})
		if err != nil {
			return nil, nil, err
		}
		groups[prefixPath] = append(groups[prefixPath], dp)
	}

	var prefixes []string
	for prefix := range groups {
		prefixes = append(prefixes, prefix)
	}
	sort.Strings(prefixes)

	return groups, prefixes, nil
}

// subscribe create a gNMI SubscribeClient. Specifying subPaths is optional, if unset will subscribe to the path at n.
func subscribe(ctx context.Context, n ygot.PathStruct, subPaths []*gpb.Path, mode gpb.SubscriptionList_Mode) (_ gpb.GNMI_SubscribeClient, _ *gpb.Path, rerr error) {
	path, dev, opts, err := resolve(ctx, n)
	if err != nil {
		return nil, path, err
	}
	if len(subPaths) == 0 {
		subPaths = []*gpb.Path{path}
	}
	ctx = metadata.NewOutgoingContext(ctx, opts.md)
	sub, err := opts.client.Subscribe(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "gNMI failed to Subscribe")
	}
	defer closer.Close(&rerr, sub.CloseSend, "error closing gNMI send stream")

	var subs []*gpb.Subscription
	for _, path := range subPaths {
		subs = append(subs, &gpb.Subscription{
			Path: &gpb.Path{
				Elem:   path.GetElem(),
				Origin: path.GetOrigin(),
			},
			Mode: opts.subMode,
		})
	}

	sr := &gpb.SubscribeRequest{
		Request: &gpb.SubscribeRequest_Subscribe{
			Subscribe: &gpb.SubscriptionList{
				Prefix: &gpb.Path{
					Target: dev.Dimensions().Name,
				},
				Subscription: subs,
				Mode:         mode,
				Encoding:     gpb.Encoding_PROTO,
			},
		},
	}
	log.V(1).Info(prototext.Format(sr))
	if err := sub.Send(sr); err != nil {
		return nil, nil, errors.Wrapf(err, "gNMI failed to Send(%+v)", sr)
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
			// Record the deprecated Element field for clearer compliance error messages.
			if elements := append(append([]string{}, n.GetPrefix().GetElement()...), p.GetElement()...); len(elements) > 0 {
				j.Element = elements
			}
			// Use the target only for the subscription but exclude from the datapoint construction.
			j.Target = ""
			return &DataPoint{Path: j, Value: val, Timestamp: ts, RecvTimestamp: recvTS}, nil
		}

		// Append delete data before the update values -- per gNMI spec, they
		// should always be processed first if both update types exist in the
		// same notification.
		for _, p := range n.Delete {
			log.V(2).Infof("Received gNMI Delete at path: %s", prototext.Format(p))
			dp, err := newDataPoint(p, nil)
			if err != nil {
				return data, false, err
			}
			log.V(2).Infof("Constructed datapoint for delete: %s", dp)
			data = append(data, dp)
		}
		for _, u := range n.GetUpdate() {
			if u.Path == nil {
				return data, false, errors.Errorf("invalid nil path in update: %v", u)
			}
			if u.Val == nil {
				return data, false, errors.Errorf("invalid nil Val in update: %v", u)
			}
			log.V(2).Infof("Received gNMI Update value %s at path: %s", prototext.Format(u.Val), prototext.Format(u.Path))
			dp, err := newDataPoint(u.Path, u.Val)
			if err != nil {
				return data, false, err
			}
			log.V(2).Infof("Constructed datapoint for update: %s", dp)
			data = append(data, dp)
		}
		return data, false, nil
	case *gpb.SubscribeResponse_SyncResponse:
		log.V(2).Infof("Received gNMI SyncResponse.")
		data = append(data, &DataPoint{
			RecvTimestamp: recvTS,
			Sync:          true,
		})
		return data, true, nil
	default:
		return data, false, errors.Errorf("unexpected response: %v (%T)", v, v)
	}
}
