// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gnmi provides an API to provides an API for querying telemetry and
// setting the state of the device via gNMI.
//
// This package provides test helpers that wrap API calls to the [ygnmi
// library]. Ondatra uses ygnmi (and by extension [ygot]) to autogenerate gNMI
// paths and value structs from an input set of YANG modules. Please read the
// [ygnmi README] to learn how ygnmi translates gNMI paths to Go API calls and
// how it handles noncompliance errors in telemetry it receives.
//
// # Available gNMI Paths
//
// Ondatra provide a combination of OpenConfig and Open Traffic Generator YANG
// modules as input to the ygnmi auto-generation.
//
// The OpenConfig paths are documented here:
//
//   - https://openconfig.net/projects/models/paths/
//   - https://openconfig.net/projects/models/schemadocs/
//
// Information on the Open Traffic Generator YANG modules is available here:
//
//   - https://github.com/open-traffic-generator/models-yang
//
// If you are using the IxNetwork (aka "ATE") API, see the [IxNetwork Telemetry]
// section below to learn which OpenConfig paths are supported by our IxNetwork
// gNMI implementation.
//
// # Best Practice: Avoid time.Sleep
//
// A test often defines a goal state based on telemetry and waits for the device to
// reach that state. It may be tempting to add a `time.Sleep` call to insert such a
// wait into the test, but sleeping is both flaky (the sleep may be too short) and
// wasteful (the sleep may be unnecessarily long). Tests can more reliably and
// efficiently wait for a particular telemetry state by using [Await] or [Watch].
//
// Example 1: To wait to an interface to be up, use [Await]:
//
//	gnmi.Await(t, dut, gnmi.OC().Interface(dp.Name()).OperStatus().State(), time.Minute, oc.Interface_OperStatus_UP)
//
// Example 2: To wait for at least 100 packets to be sent, use [Watch]. Using
// Await would not be wise here, because the counter may never register exactly
// 100.
//
//	watch := gnmi.Watch(t, dut, gnmi.OC().Interface(dp.Name()).Counters().OutPkts().State(), time.Minute, func(val *ygnmi.Value[uint64]) bool {
//		count, ok := val.Val()
//		return ok && count > 100
//	})
//	if val, ok := watch.Await(t); !ok {
//		t.Fatalf("DUT did not reach target state: got %v", val)
//	}
//
// Example 3: To wait for several interfaces to be up in parallel, use a batch
// query with [Watch]. Using multiple Await calls wouldn't be as efficient,
// because that would cause the test to wait for each interface in serial.
//
//	batch := gnmi.OCBatch()
//	for _, port := range dut.Ports() {
//		batch.AddPaths(gnmi.OC().Interface(port.Name()))
//	}
//	watch := gnmi.Watch(t, dut, batch.State(), time.Minute, func(val *ygnmi.Value[*oc.Root]) bool {
//		root, _ := val.Val()
//		for _, port := range dut.Ports() {
//			if root.GetInterface(port.Name()).GetOperStatus() != oc.Interface_OperStatus_UP {
//				return false
//			}
//		}
//		return true
//	})
//	if val, ok := watch.Await(t); !ok {
//		t.Fatalf("DUT did not reach target state: got %v", val)
//	}
//
// # IxNetwork Telemetry
//
// When using Ondatra's IxNetwork (aka "ATE") API, a test can gather statistics and
// protocol data about the IxNetwork session via gNMI. This section describes the
// available IxNetwork telemetry and the corresponding OpenConfig paths at which it
// can be queried. As the IxNetwork telemetry requires a conversion from a REST API
// to gNMI, you may encounter some necessary limitations in its behavior.
//
// To query IxNetwork Port Stats:
//
//	gnmi.OC().Interface($PortName)
//
// To query IxNetwork Port CPU Stats:
//
//	gnmi.OC().Component($PortName).Cpu()
//
// To query IxNetwork Flow Stats:
//
//	gnmi.OC().Flow($FlowName)
//
// To query IxNetwork Ingress Flow Stats:
//
//	gnmi.OC().Flow($FlowName).IngressTrackingAny()
//
// To query IxNetwork Egress Flow Stats:
//
//	gnmi.OC().Flow($FlowName).EgressTrackingAny()
//
// Learned routing data is exported under an OpenConfig "network instance" with
// the same name as the interface over which the data was learned. Specifically,
// routing data is available under gNMI paths that start
// "gnmi.OC().NetworkInstance($InterfaceName)"  where "$InterfaceName" is the
// argument to an "ate.AddInterface($InterfaceName)" call in the test.
//
// To query BGP IPv4 RIB learned routing data:
//
//	gnmi.OC().NetworkInstance($InterfaceName).
//		Protocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "0").Bgp().Rib().
//		AfiSafi(oc.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST).Ipv4Unicast().
//		Neighbor($DutIp).AdjRibInPre().Route($Prefix, $PathId)
//
// To query BGP IPv6 RIB learned routing data:
//
//	gnmi.OC().NetworkInstance($InterfaceName).
//		Protocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "0").Bgp().Rib().
//		AfiSafi(oc.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST).Ipv6Unicast().
//		Neighbor($DutIp).AdjRibInPre().Route($Prefix, $PathId)
//
// To query BGP Attr Sets learned routing data:
//
//	gnmi.OC().NetworkInstance($InterfaceName).
//		Protocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "0").Bgp().Rib().
//		AttrSet($Index)
//
// To query BGP Communities learned routing data:
//
//	gnmi.OC().NetworkInstance($InterfaceName).
//		Protocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "0").Bgp().Rib().
//		Community($Index)
//
// To query ISIS LSPs learned routing data:
//
//	gnmi.OC().NetworkInstance($InterfaceName).
//		Protocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_ISIS, "0").Isis().
//		Level($LevelNum).Lsp($LspId)
//
// To query RSVP-TE Sessions learned routing data:
//
//	gnmi.OC().NetworkInstance($InterfaceName).Mpls().SignalingProtocols().
//		RsvpTe().SessionId($LocalIndex)
//
// [ygot]: https://github.com/openconfig/ygot
// [ygnmi README]: https://github.com/openconfig/ygnmi#readme
// [ygnmi library]: https://github.com/openconfig/ygnmi
// [IxNetwork Telemetry]: https://pkg.go.dev/github.com/openconfig/ondatra/gnmi#hdr-IxNetwork_Telemetry
package gnmi

import (
	"errors"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/gnmi/oc/ocpath"
	"github.com/openconfig/ondatra/gnmi/otg/otgpath"
	"github.com/openconfig/ygnmi/ygnmi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// OC returns the root OpenConfig gNMI path.
func OC() *ocpath.RootPath {
	return ocpath.Root()
}

// OCBatch returns a new batch collection of OpenConfig gnmi paths.
func OCBatch() *ocpath.Batch {
	return new(ocpath.Batch)
}

// OTG returns the root Open Traffic Generator gNMI path.
func OTG() *otgpath.RootPath {
	return otgpath.Root()
}

// OTGBatch returns a new batch collection of Open Traffic Generator gnmi paths.
func OTGBatch() *otgpath.Batch {
	return new(otgpath.Batch)
}

var _ DeviceOrOpts = (*Opts)(nil)

// Opts contains customs options for a gNMI request.
type Opts struct {
	id              string
	initClientFn    func(context.Context) (gpb.GNMIClient, error)
	md              metadata.MD
	useGetForConfig bool
	client          gpb.GNMIClient
	opts            []ygnmi.Option
}

// NewOpts creates a set of GNMI options.
// DO NOT USE: call dut.GNMIOpts() or ate.OTG().GNMIOpts() instead.
func NewOpts(id string, useGetForConfig bool, initClient func(context.Context) (gpb.GNMIClient, error)) *Opts {
	return &Opts{
		initClientFn:    initClient,
		id:              id,
		useGetForConfig: useGetForConfig,
	}
}

// WithMetadata adds metadata to the outgoing context, overriding any set by previous calls.
func (o *Opts) WithMetadata(md metadata.MD) *Opts {
	o.md = md
	return o
}

// WithClient sets a custom gNMI client to the client, overriding any set by previous calls.
func (o *Opts) WithClient(client gpb.GNMIClient) *Opts {
	o.client = client
	return o
}

// WithYGNMIOpts adds the ygnmi Options to the client, overriding any set by previous calls.
func (o *Opts) WithYGNMIOpts(opts ...ygnmi.Option) *Opts {
	o.opts = opts
	return o
}

// GNMIOpts implements the DeviceOrOpts interface.
func (o *Opts) GNMIOpts() *Opts { return o }

// DeviceOrOpts is an interface that is ondatra.Device or a gnmi.Opts
type DeviceOrOpts interface {
	GNMIOpts() *Opts
}

// Lookup fetches the value of a ygnmi.SingletonQuery with a ONCE subscription.
func Lookup[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.SingletonQuery[T]) *ygnmi.Value[T] {
	t.Helper()
	c := newClient(t, dev, "Lookup")
	v, err := ygnmi.Lookup(createContext(dev), c, q, createGetOpts[T](dev, q)...)
	if err != nil {
		t.Fatalf("Lookup(t) on %s at %v: %v", dev, q, err)
	}
	return v
}

// LookupConfig fetches the value of a ygnmi.SingletonQuery with a ONCE subscription.
// Note: This is a workaround for Go's type inference not working for this use case and may be removed in a subsequent release.
// Note: This is equivalent to calling Lookup with a ConfigQuery and providing a fully-qualified type parameter.
func LookupConfig[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.ConfigQuery[T]) *ygnmi.Value[T] {
	t.Helper()
	return Lookup[T](t, dev, q)
}

// Get fetches the value of a SingletonQuery with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// Use Lookup to also get metadata or tolerate no value present.
func Get[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.SingletonQuery[T]) T {
	t.Helper()
	c := newClient(t, dev, "Get")
	v, err := ygnmi.Get(createContext(dev), c, q, createGetOpts[T](dev, q)...)
	if err != nil {
		t.Fatalf("Get(t) on %s at %v: %v", dev, q, err)
	}
	return v
}

type watchAwaiter[T any] interface {
	Await() (*ygnmi.Value[T], error)
}

// Watcher represents an ongoing watch of telemetry values.
type Watcher[T any] struct {
	watcher  watchAwaiter[T]
	cancelFn func()
	dev      DeviceOrOpts
	query    ygnmi.AnyQuery[T]
}

func isContextErr(err error) bool {
	// https://pkg.go.dev/google.golang.org/grpc@v1.48.0/internal/status#Error
	var st interface {
		GRPCStatus() *status.Status
	}
	ok := errors.As(err, &st)
	if !ok {
		return errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled)
	}
	return st.GRPCStatus().Code() == codes.DeadlineExceeded || st.GRPCStatus().Code() == codes.Canceled
}

// Await waits for the watch to finish and returns the last received value
// and a boolean indicating whether the predicate evaluated to true.
// When Await returns the watcher is closed, and Await may not be called again.
func (w *Watcher[T]) Await(t testing.TB) (*ygnmi.Value[T], bool) {
	t.Helper()
	v, err := w.watcher.Await()
	if err != nil {
		if isContextErr(err) {
			return v, false
		}
		t.Fatalf("Await(t) on %s at %v: %v", w.dev, w.query, err)
	}
	return v, true
}

// Cancel stops the watch immediately.
func (w *Watcher[T]) Cancel() {
	w.cancelFn()
}

// Watch starts an asynchronous STREAM subscription, evaluating each observed value with the
// specified predicate. The subscription completes when either the predicate is true
// or the timeout is reached. Calling Await on the returned Watcher waits for the subscription
// to complete. It returns the last observed value and a boolean that indicates whether
// that value satisfies the predicate.
func Watch[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.SingletonQuery[T], timeout time.Duration, pred func(*ygnmi.Value[T]) bool) *Watcher[T] {
	t.Helper()
	c := newClient(t, dev, "Watch")
	ctx, cancel := context.WithTimeout(createContext(dev), timeout)
	w := ygnmi.Watch(ctx, c, q, func(v *ygnmi.Value[T]) error {
		if ok := pred(v); ok {
			return nil
		}
		return ygnmi.Continue
	}, dev.GNMIOpts().opts...)
	return &Watcher[T]{
		watcher:  w,
		cancelFn: cancel,
		dev:      dev,
		query:    q,
	}
}

// Await observes values at Query with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or the timeout is reached. To wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func Await[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.SingletonQuery[T], timeout time.Duration, val T) *ygnmi.Value[T] {
	t.Helper()
	c := newClient(t, dev, "Await")
	ctx, cancel := context.WithTimeout(createContext(dev), timeout)
	defer cancel()

	v, err := ygnmi.Await(ctx, c, q, val, dev.GNMIOpts().opts...)
	if err != nil {
		t.Fatalf("Await(t) on %s at %v: %v", dev, q, err)
	}
	return v
}

type collectAwaiter[T any] interface {
	Await() ([]*ygnmi.Value[T], error)
}

// Collector represents an ongoing collection of telemetry values.
type Collector[T any] struct {
	collector collectAwaiter[T]
	cancelFn  func()
	dev       DeviceOrOpts
	query     ygnmi.AnyQuery[T]
}

// Await waits for the collection to finish and returns all received values.
// When Await returns the watcher is closed, and Await may not be called again.
// Note: the func blocks until the timeout is reached.
func (c *Collector[T]) Await(t testing.TB) []*ygnmi.Value[T] {
	t.Helper()
	vals, err := c.collector.Await()
	if err != nil {
		if isContextErr(err) {
			return vals
		}
		t.Fatalf("Await(t) on %s at %v: %v", c.dev, c.query, err)
	}
	return vals
}

// Collect starts an asynchronous collection of the values at the query with a STREAM subscription.
// Calling Await on the return Collection waits until the timeout is reached and returns the collected values.
func Collect[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.SingletonQuery[T], timeout time.Duration) *Collector[T] {
	t.Helper()
	c := newClient(t, dev, "Collect")
	ctx, cancel := context.WithTimeout(createContext(dev), timeout)
	collect := &Collector[T]{
		collector: ygnmi.Collect(ctx, c, q, dev.GNMIOpts().opts...),
		cancelFn:  cancel,
		dev:       dev,
		query:     q,
	}

	return collect
}

// LookupAll fetches the values of a ygnmi.WildcardQuery with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func LookupAll[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.WildcardQuery[T]) []*ygnmi.Value[T] {
	t.Helper()
	c := newClient(t, dev, "LookupAll")
	v, err := ygnmi.LookupAll(createContext(dev), c, q, createGetOpts[T](dev, q)...)
	if err != nil {
		t.Fatalf("LookupAll(t) on %s at %v: %v", dev, q, err)
	}
	return v
}

// GetAll fetches the value of a WildcardQuery with a ONCE subscription skipping any non-present paths.
// It fails the test fatally if no value is present at the path
// Use LookupAll to also get metadata or tolerate no values present.
func GetAll[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.WildcardQuery[T]) []T {
	t.Helper()
	c := newClient(t, dev, "GetAll")
	v, err := ygnmi.GetAll(createContext(dev), c, q, createGetOpts[T](dev, q)...)
	if err != nil {
		t.Fatalf("GetAll(t) on %s at %v: %v", dev, q, err)
	}
	return v
}

// WatchAll starts an asynchronous STREAM subscription, evaluating each observed value with the
// specified predicate.  The subscription completes when either the predicate is true
// or the timeout is reached. Calling Await on the returned Watcher waits for the subscription
// to complete. It returns the last observed value and a boolean that indicates whether
// that value satisfies the predicate.
func WatchAll[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.WildcardQuery[T], timeout time.Duration, pred func(*ygnmi.Value[T]) bool) *Watcher[T] {
	t.Helper()
	c := newClient(t, dev, "WatchAll")
	ctx, cancel := context.WithTimeout(createContext(dev), timeout)
	w := ygnmi.WatchAll(ctx, c, q, func(v *ygnmi.Value[T]) error {
		if ok := pred(v); ok {
			return nil
		}
		return ygnmi.Continue
	}, dev.GNMIOpts().opts...)
	return &Watcher[T]{
		watcher:  w,
		cancelFn: cancel,
		dev:      dev,
		query:    q,
	}
}

// CollectAll starts an asynchronous collection of the values at the query with a STREAM subscription.
// Calling Await on the return Collection waits until the timeout is reached and returns the collected values.
func CollectAll[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.WildcardQuery[T], timeout time.Duration) *Collector[T] {
	t.Helper()
	c := newClient(t, dev, "CollectAll")
	ctx, cancel := context.WithTimeout(createContext(dev), timeout)
	collect := &Collector[T]{
		collector: ygnmi.CollectAll(ctx, c, q, dev.GNMIOpts().opts...),
		cancelFn:  cancel,
		dev:       dev,
		query:     q,
	}

	return collect
}

// Update updates the configuration at the given query path with the val.
func Update[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.ConfigQuery[T], val T) *ygnmi.Result {
	t.Helper()
	c := newClient(t, dev, "Update")
	res, err := ygnmi.Update(createContext(dev), c, q, val)
	if err != nil {
		t.Fatalf("Update(t) on %s %v: %v", dev, q, err)
	}
	return res
}

// Replace replaces the configuration at the given query path with the val.
func Replace[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.ConfigQuery[T], val T) *ygnmi.Result {
	t.Helper()
	c := newClient(t, dev, "Replace")
	res, err := ygnmi.Replace(createContext(dev), c, q, val)
	if err != nil {
		t.Fatalf("Replace(t) on %s at %v: %v", dev, q, err)
	}
	return res
}

// Delete deletes the configuration at the given query path.
func Delete[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.ConfigQuery[T]) *ygnmi.Result {
	t.Helper()
	c := newClient(t, dev, "Delete")
	res, err := ygnmi.Delete(createContext(dev), c, q)
	if err != nil {
		t.Fatalf("Delete(t) on %s at %v: %v", dev, q, err)
	}
	return res
}

// SetBatch allows multiple Set operations (Replace, Update, Delete) to be applied as part of a single Set transaction.
// Use BatchUpdate, BatchReplace, BatchDelete to add operations, and then call the Set method to send the SetRequest.
type SetBatch struct {
	sb ygnmi.SetBatch
}

// Set performs the gnmi.Set request with all queued operations.
func (sb *SetBatch) Set(t testing.TB, dev DeviceOrOpts) *ygnmi.Result {
	t.Helper()
	c := newClient(t, dev, "Set")
	res, err := sb.sb.Set(createContext(dev), c)
	if err != nil {
		t.Fatalf("Set(t) on %s: %v", dev, err)
	}
	return res
}

// BatchUpdate stores an update operation in the SetBatch.
func BatchUpdate[T any](sb *SetBatch, q ygnmi.ConfigQuery[T], val T) {
	ygnmi.BatchUpdate(&sb.sb, q, val)
}

// BatchReplace stores an replace operation in the SetBatch.
func BatchReplace[T any](sb *SetBatch, q ygnmi.ConfigQuery[T], val T) {
	ygnmi.BatchReplace(&sb.sb, q, val)
}

// BatchUnionReplace stores a union_replace operation in the SetBatch.
//
// https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-union_replace.md
func BatchUnionReplace[T any](sb *SetBatch, q ygnmi.ConfigQuery[T], val T) {
	ygnmi.BatchUnionReplace(&sb.sb, q, val)
}

// BatchUnionReplaceCLI stores a CLI union_replace operation in the SetBatch.
//
//   - nos is the name of the Network operating system.
//     "_cli" is appended to it to form the origin, see
//     https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-union_replace.md#24-native-cli-configuration-cli
//   - ascii is the full CLI text.
//
// https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-union_replace.md
func BatchUnionReplaceCLI(sb *SetBatch, nos, ascii string) {
	ygnmi.BatchUnionReplaceCLI(&sb.sb, nos, ascii)
}

// BatchDelete stores a delete operation in the SetBatch.
func BatchDelete[T any](sb *SetBatch, q ygnmi.ConfigQuery[T]) {
	ygnmi.BatchDelete(&sb.sb, q)
}

func createContext(d DeviceOrOpts) context.Context {
	ctx := context.Background()
	if d.GNMIOpts().md == nil {
		return ctx
	}
	return metadata.NewOutgoingContext(ctx, d.GNMIOpts().md)
}

func createGetOpts[T any](d DeviceOrOpts, q ygnmi.AnyQuery[T]) []ygnmi.Option {
	opts := d.GNMIOpts().opts
	if d.GNMIOpts().useGetForConfig && !q.IsState() {
		opts = append(opts, ygnmi.WithUseGet())
	}
	return opts
}

func newClient(t testing.TB, dev DeviceOrOpts, method string) *ygnmi.Client {
	t.Helper()
	gnmiC := dev.GNMIOpts().client
	if gnmiC == nil {
		var err error
		gnmiC, err = dev.GNMIOpts().initClientFn(context.Background())
		if err != nil {
			t.Fatalf("Failed to get client: %v", err)
		}
	}

	c, err := ygnmi.NewClient(gnmiC, ygnmi.WithTarget(dev.GNMIOpts().id))
	if err != nil {
		t.Fatalf("%s(t) on %s: %v", method, dev, err)
	}
	return c
}
