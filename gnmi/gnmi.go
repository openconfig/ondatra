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

//go:build go1.18

// Package gnmi contains test-friendly gNMI funcs.
package gnmi

import (
	"errors"
	"io"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra"
	"github.com/openconfig/ygnmi/ygnmi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Lookup fetches the value of a ygnmi.SingletonQuery with a ONCE subscription.
func Lookup[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.SingletonQuery[T]) *ygnmi.Value[T] {
	c := newClient(t, dut, "Lookup")
	v, err := ygnmi.Lookup(context.Background(), c, q)
	if err != nil {
		t.Fatalf("Lookup(t) on %s at %v: %v", dut, q, err)
	}
	return v
}

// Get fetches the value of a SingletonQuery with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// Use Lookup to also get metadata or tolerate no value present.
func Get[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.SingletonQuery[T]) T {
	c := newClient(t, dut, "Get")
	v, err := ygnmi.Get(context.Background(), c, q)
	if err != nil {
		t.Fatalf("Get(t) on %s at %v: %v", dut, q, err)
	}
	return v
}

// Watcher represents an ongoing watch of telemetry values.
type Watcher[T any] struct {
	watcher  *ygnmi.Watcher[T]
	cancelFn func()
	dut      *ondatra.DUTDevice
	query    ygnmi.AnyQuery[T]
}

// Await waits for the watch to finish and returns the last received value
// and a boolean indicating whether the predicate evaluated to true.
// When Await returns the watcher is closed, and Await may not be called again.
func (w *Watcher[T]) Await(t testing.TB) (*ygnmi.Value[T], bool) {
	v, err := w.watcher.Await()
	if err != nil {
		if st, ok := status.FromError(err); (ok && st.Code() == codes.DeadlineExceeded) || errors.Is(err, io.EOF) {
			return v, false
		}
		t.Fatalf("Await(t) on %s at %v: %v", w.dut, w.query, err)
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
func Watch[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.SingletonQuery[T], timeout time.Duration, pred func(*ygnmi.Value[T]) bool) *Watcher[T] {
	c := newClient(t, dut, "Watch")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	w := ygnmi.Watch(ctx, c, q, func(v *ygnmi.Value[T]) error {
		if ok := pred(v); ok {
			return nil
		}
		return ygnmi.Continue
	})
	return &Watcher[T]{
		watcher:  w,
		cancelFn: cancel,
		dut:      dut,
		query:    q,
	}
}

// Await observes values at Query with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or the timeout is reached. To wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func Await[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.SingletonQuery[T], timeout time.Duration, val T) *ygnmi.Value[T] {
	c := newClient(t, dut, "Await")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	v, err := ygnmi.Await(ctx, c, q, val)
	if err != nil {
		t.Fatalf("Await(t) on %s at %v: %v", dut, q, err)
	}
	return v
}

// Collector represents an ongoing collection of telemetry values.
type Collector[T any] struct {
	collector *ygnmi.Collector[T]
	cancelFn  func()
	dut       *ondatra.DUTDevice
	query     ygnmi.AnyQuery[T]
}

// Await waits for the collection to finish and returns all received values.
// When Await returns the watcher is closed, and Await may not be called again.
// Note: the func blocks until the timeout is reached.
func (c *Collector[T]) Await(t testing.TB) []*ygnmi.Value[T] {
	vals, err := c.collector.Await()
	if err != nil {
		t.Fatalf("Await(t) on %s at %v: %v", c.dut, c.query, err)
	}
	return vals
}

// Collect starts an asynchronous collection of the values at the query with a STREAM subscription.
// Calling Await on the return Collection waits until the timeout is reached and returns the collected values.
func Collect[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.SingletonQuery[T], timeout time.Duration) *Collector[T] {
	c := newClient(t, dut, "Collect")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	collect := &Collector[T]{
		collector: ygnmi.Collect(ctx, c, q),
		cancelFn:  cancel,
		dut:       dut,
		query:     q,
	}

	return collect
}

// LookupAll fetches the values of a ygnmi.WildcardQuery with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func LookupAll[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.WildcardQuery[T]) []*ygnmi.Value[T] {
	c := newClient(t, dut, "LookupAll")
	v, err := ygnmi.LookupAll(context.Background(), c, q)
	if err != nil {
		t.Fatalf("LookupAll(t) on %s at %v: %v", dut, q, err)
	}
	return v
}

// GetAll fetches the value of a WildcardQuery with a ONCE subscription skipping any non-present paths.
// It fails the test fatally if no value is present at the path
// Use LookupAll to also get metadata or tolerate no values present.
func GetAll[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.WildcardQuery[T]) []T {
	c := newClient(t, dut, "GetAll")
	v, err := ygnmi.GetAll(context.Background(), c, q)
	if err != nil {
		t.Fatalf("GetAll(t) on %s at %v: %v", dut, q, err)
	}
	return v
}

// WatchAll starts an asynchronous STREAM subscription, evaluating each observed value with the
// specified predicate.  The subscription completes when either the predicate is true
// or the timeout is reached. Calling Await on the returned Watcher waits for the subscription
// to complete. It returns the last observed value and a boolean that indicates whether
// that value satisfies the predicate.
func WatchAll[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.WildcardQuery[T], timeout time.Duration, pred func(*ygnmi.Value[T]) bool) *Watcher[T] {
	c := newClient(t, dut, "WatchAll")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	w := ygnmi.WatchAll(ctx, c, q, func(v *ygnmi.Value[T]) error {
		if ok := pred(v); ok {
			return nil
		}
		return ygnmi.Continue
	})
	return &Watcher[T]{
		watcher:  w,
		cancelFn: cancel,
		dut:      dut,
		query:    q,
	}
}

// CollectAll starts an asynchronous collection of the values at the query with a STREAM subscription.
// Calling Await on the return Collection waits until the timeout is reached and returns the collected values.
func CollectAll[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.WildcardQuery[T], timeout time.Duration) *Collector[T] {
	c := newClient(t, dut, "CollectAll")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	collect := &Collector[T]{
		collector: ygnmi.CollectAll(ctx, c, q),
		cancelFn:  cancel,
		dut:       dut,
		query:     q,
	}

	return collect
}

// Update updates the configuration at the given query path with the val.
func Update[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.ConfigQuery[T], val T) *ygnmi.Result {
	c := newClient(t, dut, "Update")
	res, err := ygnmi.Update(context.Background(), c, q, val)
	if err != nil {
		t.Fatalf("Update(t) on %s %v: %v", dut, q, err)
	}
	return res
}

// Replace replaces the configuration at the given query path with the val.
func Replace[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.ConfigQuery[T], val T) *ygnmi.Result {
	c := newClient(t, dut, "Replace")
	res, err := ygnmi.Replace(context.Background(), c, q, val)
	if err != nil {
		t.Fatalf("Replace(t) on %s at %v: %v", dut, q, err)
	}
	return res
}

// Delete deletes the configuration at the given query path.
func Delete[T any](t testing.TB, dut *ondatra.DUTDevice, q ygnmi.ConfigQuery[T]) *ygnmi.Result {
	c := newClient(t, dut, "Delete")
	res, err := ygnmi.Delete(context.Background(), c, q)
	if err != nil {
		t.Fatalf("Delete(t) on %s at %v: %v", dut, q, err)
	}
	return res
}

// SetBatch allows multiple Set operations (Replace, Update, Delete) to be applied as part of a single Set transaction.
// Use BatchUpdate, BatchReplace, BatchDelete to add operations, and then call the Set method to send the SetRequest.
type SetBatch struct {
	sb *ygnmi.SetBatch
}

// Set performs the gnmi.Set request with all queued operations.
func (sb *SetBatch) Set(t testing.TB, dut *ondatra.DUTDevice) *ygnmi.Result {
	c := newClient(t, dut, "Set")
	res, err := sb.sb.Set(context.Background(), c)
	if err != nil {
		t.Fatalf("Set(t) on %s: %v", dut, err)
	}
	return res
}

// BatchUpdate stores an update operation in the SetBatch.
func BatchUpdate[T any](sb *SetBatch, q ygnmi.ConfigQuery[T], val T) {
	ygnmi.BatchUpdate(sb.sb, q, val)
}

// BatchReplace stores an replace operation in the SetBatch.
func BatchReplace[T any](sb *SetBatch, q ygnmi.ConfigQuery[T], val T) {
	ygnmi.BatchReplace(sb.sb, q, val)
}

// BatchDelete stores an delete operation in the SetBatch.
func BatchDelete[T any](sb *SetBatch, q ygnmi.ConfigQuery[T]) {
	ygnmi.BatchDelete(sb.sb, q)
}

func newClient(t testing.TB, dut *ondatra.DUTDevice, method string) *ygnmi.Client {
	c, err := ygnmi.NewClient(dut.RawAPIs().GNMI().Default(t), ygnmi.WithTarget(dut.ID()))
	if err != nil {
		t.Fatalf("%s(t) on %s: %v", method, dut, err)
	}
	return c
}
