package opentrafficgeneratorport

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

// Lookup fetches the value at /open-traffic-generator-port/ports/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *PortPath) Lookup(t testing.TB) *oc.QualifiedPort {
	t.Helper()
	goStruct := &oc.Port{}
	md, ok := oc.Lookup(t, n, "Port", goStruct, false, false)
	if ok {
		return (&oc.QualifiedPort{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-port/ports/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *PortPath) Get(t testing.TB) *oc.Port {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-port/ports/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *PortPathAny) Lookup(t testing.TB) []*oc.QualifiedPort {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedPort
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Port{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Port", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedPort{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-port/ports/port with a ONCE subscription.
func (n *PortPathAny) Get(t testing.TB) []*oc.Port {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Port
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *PortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionPort {
	t.Helper()
	c := &oc.CollectionPort{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedPort) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedPort{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Port)))
		return false
	})
	return c
}

func watch_PortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedPort) bool) *oc.PortWatcher {
	t.Helper()
	w := &oc.PortWatcher{}
	gs := &oc.Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Port", gs, queryPath, false, false)
		qv := (&oc.QualifiedPort{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedPort)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *PortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedPort) bool) *oc.PortWatcher {
	t.Helper()
	return watch_PortPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-port/ports/port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *PortPath) Await(t testing.TB, timeout time.Duration, val *oc.Port) *oc.QualifiedPort {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedPort) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-port/ports/port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-port/ports/port to the batch object.
func (n *PortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *PortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionPort {
	t.Helper()
	c := &oc.CollectionPort{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedPort) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_PortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedPort) bool) *oc.PortWatcher {
	t.Helper()
	w := &oc.PortWatcher{}
	structs := map[string]*oc.Port{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Port", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedPort{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedPort)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *PortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedPort) bool) *oc.PortWatcher {
	t.Helper()
	return watch_PortPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-port/ports/port to the batch object.
func (n *PortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-port/ports/port/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Port_CountersPath) Lookup(t testing.TB) *oc.QualifiedPort_Counters {
	t.Helper()
	goStruct := &oc.Port_Counters{}
	md, ok := oc.Lookup(t, n, "Port_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedPort_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-port/ports/port/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Port_CountersPath) Get(t testing.TB) *oc.Port_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-port/ports/port/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Port_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedPort_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedPort_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Port_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Port_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedPort_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-port/ports/port/state/counters with a ONCE subscription.
func (n *Port_CountersPathAny) Get(t testing.TB) []*oc.Port_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Port_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionPort_Counters {
	t.Helper()
	c := &oc.CollectionPort_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedPort_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedPort_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Port_Counters)))
		return false
	})
	return c
}

func watch_Port_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedPort_Counters) bool) *oc.Port_CountersWatcher {
	t.Helper()
	w := &oc.Port_CountersWatcher{}
	gs := &oc.Port_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Port_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedPort_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedPort_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedPort_Counters) bool) *oc.Port_CountersWatcher {
	t.Helper()
	return watch_Port_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-port/ports/port/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Port_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Port_Counters) *oc.QualifiedPort_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedPort_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-port/ports/port/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-port/ports/port/state/counters to the batch object.
func (n *Port_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionPort_Counters {
	t.Helper()
	c := &oc.CollectionPort_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedPort_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedPort_Counters) bool) *oc.Port_CountersWatcher {
	t.Helper()
	w := &oc.Port_CountersWatcher{}
	structs := map[string]*oc.Port_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Port_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedPort_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedPort_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedPort_Counters) bool) *oc.Port_CountersWatcher {
	t.Helper()
	return watch_Port_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-port/ports/port/state/counters to the batch object.
func (n *Port_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-port/ports/port/state/counters/in-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Port_Counters_InFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Port_Counters{}
	md, ok := oc.Lookup(t, n, "Port_Counters", goStruct, true, false)
	if ok {
		return convertPort_Counters_InFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-port/ports/port/state/counters/in-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Port_Counters_InFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-port/ports/port/state/counters/in-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Port_Counters_InFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertPort_Counters_InFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-port/ports/port/state/counters/in-frames with a ONCE subscription.
func (n *Port_Counters_InFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/counters/in-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_Counters_InFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_Counters_InFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Port_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Port_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertPort_Counters_InFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/counters/in-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_Counters_InFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Port_Counters_InFramesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-port/ports/port/state/counters/in-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Port_Counters_InFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-port/ports/port/state/counters/in-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-port/ports/port/state/counters/in-frames to the batch object.
func (n *Port_Counters_InFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/counters/in-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_Counters_InFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_Counters_InFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertPort_Counters_InFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/counters/in-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_Counters_InFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Port_Counters_InFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-port/ports/port/state/counters/in-frames to the batch object.
func (n *Port_Counters_InFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertPort_Counters_InFramesPath extracts the value of the leaf InFrames from its parent oc.Port_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertPort_Counters_InFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.Port_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-port/ports/port/state/counters/in-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Port_Counters_InOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Port_Counters{}
	md, ok := oc.Lookup(t, n, "Port_Counters", goStruct, true, false)
	if ok {
		return convertPort_Counters_InOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-port/ports/port/state/counters/in-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Port_Counters_InOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-port/ports/port/state/counters/in-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Port_Counters_InOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertPort_Counters_InOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-port/ports/port/state/counters/in-octets with a ONCE subscription.
func (n *Port_Counters_InOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_Counters_InOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_Counters_InOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Port_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Port_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertPort_Counters_InOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_Counters_InOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Port_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-port/ports/port/state/counters/in-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Port_Counters_InOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-port/ports/port/state/counters/in-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-port/ports/port/state/counters/in-octets to the batch object.
func (n *Port_Counters_InOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-port/ports/port/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Port_Counters_InOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Port_Counters_InOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertPort_Counters_InOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-port/ports/port/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Port_Counters_InOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Port_Counters_InOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-port/ports/port/state/counters/in-octets to the batch object.
func (n *Port_Counters_InOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertPort_Counters_InOctetsPath extracts the value of the leaf InOctets from its parent oc.Port_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertPort_Counters_InOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Port_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
