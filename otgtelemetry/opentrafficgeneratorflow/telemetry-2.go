package opentrafficgeneratorflow

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/otgtelemetry"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/state/out-frame-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_OutFrameRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Flow{}
	md, ok := oc.Lookup(t, n, "Flow", goStruct, true, false)
	if ok {
		return convertFlow_OutFrameRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/state/out-frame-rate with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_OutFrameRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/state/out-frame-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_OutFrameRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_OutFrameRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/state/out-frame-rate with a ONCE subscription.
func (n *Flow_OutFrameRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/out-frame-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_OutFrameRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_OutFrameRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Flow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow", gs, queryPath, true, false)
		return convertFlow_OutFrameRatePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/out-frame-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_OutFrameRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_OutFrameRatePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/state/out-frame-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_OutFrameRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/state/out-frame-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/out-frame-rate to the batch object.
func (n *Flow_OutFrameRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/out-frame-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_OutFrameRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/out-frame-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_OutFrameRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_OutFrameRatePath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/out-frame-rate to the batch object.
func (n *Flow_OutFrameRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_OutFrameRatePath extracts the value of the leaf OutFrameRate from its parent oc.Flow
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertFlow_OutFrameRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Flow) *oc.QualifiedFloat32 {
	t.Helper()
	qv := &oc.QualifiedFloat32{
		Metadata: md,
	}
	val := parent.OutFrameRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/state/out-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_OutRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Flow{}
	md, ok := oc.Lookup(t, n, "Flow", goStruct, true, false)
	if ok {
		return convertFlow_OutRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/state/out-rate with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_OutRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/state/out-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_OutRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_OutRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/state/out-rate with a ONCE subscription.
func (n *Flow_OutRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/out-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_OutRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_OutRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Flow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow", gs, queryPath, true, false)
		return convertFlow_OutRatePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/out-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_OutRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_OutRatePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/state/out-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_OutRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/state/out-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/out-rate to the batch object.
func (n *Flow_OutRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/out-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_OutRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/out-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_OutRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_OutRatePath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/out-rate to the batch object.
func (n *Flow_OutRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_OutRatePath extracts the value of the leaf OutRate from its parent oc.Flow
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertFlow_OutRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Flow) *oc.QualifiedFloat32 {
	t.Helper()
	qv := &oc.QualifiedFloat32{
		Metadata: md,
	}
	val := parent.OutRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/state/transmit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_TransmitPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Flow{}
	md, ok := oc.Lookup(t, n, "Flow", goStruct, true, false)
	if ok {
		return convertFlow_TransmitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/state/transmit with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_TransmitPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/state/transmit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_TransmitPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_TransmitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/state/transmit with a ONCE subscription.
func (n *Flow_TransmitPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/transmit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_TransmitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_TransmitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Flow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow", gs, queryPath, true, false)
		return convertFlow_TransmitPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/transmit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_TransmitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Flow_TransmitPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/state/transmit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_TransmitPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/state/transmit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/transmit to the batch object.
func (n *Flow_TransmitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/transmit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_TransmitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/transmit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_TransmitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Flow_TransmitPath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/transmit to the batch object.
func (n *Flow_TransmitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_TransmitPath extracts the value of the leaf Transmit from its parent oc.Flow
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertFlow_TransmitPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Transmit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
