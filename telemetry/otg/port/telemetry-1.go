package port

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry/otg"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /open-traffic-generator-port/ports/port/state/counters/out-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Port_Counters_OutFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Port_Counters{}
	md, ok := oc.Lookup(t, n, "Port_Counters", goStruct, true, false)
	if ok {
		return convertPort_Counters_OutFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-port/ports/port/state/counters/out-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Port_Counters_OutFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-port/ports/port/state/counters/out-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Port_Counters_OutFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Port_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Port_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertPort_Counters_OutFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-port/ports/port/state/counters/out-frames with a ONCE subscription.
func (n *Port_Counters_OutFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/counters/out-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_Counters_OutFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_Counters_OutFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Port_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Port_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertPort_Counters_OutFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/counters/out-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_Counters_OutFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Port_Counters_OutFramesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-port/ports/port/state/counters/out-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Port_Counters_OutFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-port/ports/port/state/counters/out-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-port/ports/port/state/counters/out-frames to the batch object.
func (n *Port_Counters_OutFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/counters/out-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_Counters_OutFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_Counters_OutFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Port_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Port_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Port_Counters", structs[pre], queryPath, true, false)
			qv := convertPort_Counters_OutFramesPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/counters/out-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_Counters_OutFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Port_Counters_OutFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-port/ports/port/state/counters/out-frames to the batch object.
func (n *Port_Counters_OutFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertPort_Counters_OutFramesPath extracts the value of the leaf OutFrames from its parent oc.Port_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertPort_Counters_OutFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Port_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-port/ports/port/state/counters/out-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Port_Counters_OutOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Port_Counters{}
	md, ok := oc.Lookup(t, n, "Port_Counters", goStruct, true, false)
	if ok {
		return convertPort_Counters_OutOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-port/ports/port/state/counters/out-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Port_Counters_OutOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-port/ports/port/state/counters/out-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Port_Counters_OutOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Port_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Port_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertPort_Counters_OutOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-port/ports/port/state/counters/out-octets with a ONCE subscription.
func (n *Port_Counters_OutOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_Counters_OutOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_Counters_OutOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Port_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Port_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertPort_Counters_OutOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_Counters_OutOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Port_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-port/ports/port/state/counters/out-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Port_Counters_OutOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-port/ports/port/state/counters/out-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-port/ports/port/state/counters/out-octets to the batch object.
func (n *Port_Counters_OutOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_Counters_OutOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_Counters_OutOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Port_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Port_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Port_Counters", structs[pre], queryPath, true, false)
			qv := convertPort_Counters_OutOctetsPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_Counters_OutOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Port_Counters_OutOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-port/ports/port/state/counters/out-octets to the batch object.
func (n *Port_Counters_OutOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertPort_Counters_OutOctetsPath extracts the value of the leaf OutOctets from its parent oc.Port_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertPort_Counters_OutOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Port_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-port/ports/port/state/in-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Port_InRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Port{}
	md, ok := oc.Lookup(t, n, "Port", goStruct, true, false)
	if ok {
		return convertPort_InRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-port/ports/port/state/in-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Port_InRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-port/ports/port/state/in-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Port_InRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Port{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Port", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertPort_InRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-port/ports/port/state/in-rate with a ONCE subscription.
func (n *Port_InRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/in-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_InRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_InRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Port", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertPort_InRatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/in-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_InRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Port_InRatePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-port/ports/port/state/in-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Port_InRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-port/ports/port/state/in-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-port/ports/port/state/in-rate to the batch object.
func (n *Port_InRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/in-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_InRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_InRatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	structs := map[string]*oc.Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Port{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Port", structs[pre], queryPath, true, false)
			qv := convertPort_InRatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/in-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_InRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Port_InRatePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-port/ports/port/state/in-rate to the batch object.
func (n *Port_InRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertPort_InRatePath extracts the value of the leaf InRate from its parent oc.Port
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertPort_InRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Port) *oc.QualifiedFloat32 {
	t.Helper()
	qv := &oc.QualifiedFloat32{
		Metadata: md,
	}
	val := parent.InRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-port/ports/port/state/link with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Port_LinkPath) Lookup(t testing.TB) *oc.QualifiedE_Port_Link {
	t.Helper()
	goStruct := &oc.Port{}
	md, ok := oc.Lookup(t, n, "Port", goStruct, true, false)
	if ok {
		return convertPort_LinkPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-port/ports/port/state/link with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Port_LinkPath) Get(t testing.TB) oc.E_Port_Link {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-port/ports/port/state/link with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Port_LinkPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Port_Link {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Port_Link
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Port{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Port", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertPort_LinkPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-port/ports/port/state/link with a ONCE subscription.
func (n *Port_LinkPathAny) Get(t testing.TB) []oc.E_Port_Link {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Port_Link
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/link with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_LinkPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Port_Link {
	t.Helper()
	c := &oc.CollectionE_Port_Link{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Port_Link) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_LinkPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Port_Link) bool) *oc.E_Port_LinkWatcher {
	t.Helper()
	w := &oc.E_Port_LinkWatcher{}
	gs := &oc.Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Port", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertPort_LinkPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Port_Link)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/link with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_LinkPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Port_Link) bool) *oc.E_Port_LinkWatcher {
	t.Helper()
	return watch_Port_LinkPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-port/ports/port/state/link with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Port_LinkPath) Await(t testing.TB, timeout time.Duration, val oc.E_Port_Link) *oc.QualifiedE_Port_Link {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Port_Link) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-port/ports/port/state/link failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-port/ports/port/state/link to the batch object.
func (n *Port_LinkPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/link with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_LinkPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Port_Link {
	t.Helper()
	c := &oc.CollectionE_Port_Link{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Port_Link) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_LinkPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Port_Link) bool) *oc.E_Port_LinkWatcher {
	t.Helper()
	w := &oc.E_Port_LinkWatcher{}
	structs := map[string]*oc.Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Port{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Port", structs[pre], queryPath, true, false)
			qv := convertPort_LinkPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Port_Link)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/link with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_LinkPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Port_Link) bool) *oc.E_Port_LinkWatcher {
	t.Helper()
	return watch_Port_LinkPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-port/ports/port/state/link to the batch object.
func (n *Port_LinkPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertPort_LinkPath extracts the value of the leaf Link from its parent oc.Port
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Port_Link.
func convertPort_LinkPath(t testing.TB, md *genutil.Metadata, parent *oc.Port) *oc.QualifiedE_Port_Link {
	t.Helper()
	qv := &oc.QualifiedE_Port_Link{
		Metadata: md,
	}
	val := parent.Link
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
