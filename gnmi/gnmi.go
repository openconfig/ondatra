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

// Package gnmi contains test-friendly gNMI funcs.
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

// GetConfig fetches the value of a SingletonQuery with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// Use Lookup to also get metadata or tolerate no value present.
// Note: This is a workaround for Go's type inference not working for this use case and may be removed in a subsequent release.
// Note: This is equivalent to calling Get with a ConfigQuery and providing a fully-qualified type parameter.
func GetConfig[T any](t testing.TB, dev DeviceOrOpts, q ygnmi.ConfigQuery[T]) T {
	t.Helper()
	return Get[T](t, dev, q)
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
	return ok && (st.GRPCStatus().Code() == codes.DeadlineExceeded || st.GRPCStatus().Code() == codes.Canceled)
}

// statusErr is an interface implemented by errors returned by gRPC.
// https://pkg.go.dev/google.golang.org/grpc@v1.48.0/internal/status#Error
type statusErr interface {
	GRPCStatus() *status.Status
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

// BatchDelete stores an delete operation in the SetBatch.
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
