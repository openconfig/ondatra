package system

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/wait/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Wait_InstantPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Wait{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Wait", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Wait_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/wait/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Wait_InstantPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/wait/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Wait_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Wait{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Wait", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Wait_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/wait/instant with a ONCE subscription.
func (n *System_Cpu_Wait_InstantPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Wait", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Wait_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/wait/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Wait_InstantPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/wait/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/instant to the batch object.
func (n *System_Cpu_Wait_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Wait{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Wait", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Wait_InstantPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/instant to the batch object.
func (n *System_Cpu_Wait_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Wait_InstantPath extracts the value of the leaf Instant from its parent oc.System_Cpu_Wait
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Wait_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Wait) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Instant
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/wait/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Wait_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Wait{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Wait", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Wait_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/wait/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Wait_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/wait/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Wait_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Wait{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Wait", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Wait_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/wait/interval with a ONCE subscription.
func (n *System_Cpu_Wait_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Wait", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Wait_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/wait/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Wait_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/wait/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/interval to the batch object.
func (n *System_Cpu_Wait_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Wait{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Wait", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Wait_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/interval to the batch object.
func (n *System_Cpu_Wait_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Wait_IntervalPath extracts the value of the leaf Interval from its parent oc.System_Cpu_Wait
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Wait_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Wait) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/wait/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Wait_MaxPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Wait{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Wait", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Wait_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/wait/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Wait_MaxPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/wait/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Wait_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Wait{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Wait", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Wait_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/wait/max with a ONCE subscription.
func (n *System_Cpu_Wait_MaxPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Wait", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Wait_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/wait/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Wait_MaxPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/wait/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/max to the batch object.
func (n *System_Cpu_Wait_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Wait{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Wait", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Wait_MaxPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/max to the batch object.
func (n *System_Cpu_Wait_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Wait_MaxPath extracts the value of the leaf Max from its parent oc.System_Cpu_Wait
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Wait_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Wait) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Max
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/wait/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Wait_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Wait{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Wait", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Wait_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/wait/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Wait_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/wait/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Wait_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Wait{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Wait", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Wait_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/wait/max-time with a ONCE subscription.
func (n *System_Cpu_Wait_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Wait", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Wait_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/wait/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Wait_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/wait/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/max-time to the batch object.
func (n *System_Cpu_Wait_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Wait{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Wait", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Wait_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/max-time to the batch object.
func (n *System_Cpu_Wait_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Wait_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.System_Cpu_Wait
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Wait_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Wait) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/wait/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Wait_MinPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Wait{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Wait", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Wait_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/wait/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Wait_MinPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/wait/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Wait_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Wait{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Wait", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Wait_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/wait/min with a ONCE subscription.
func (n *System_Cpu_Wait_MinPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Wait", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Wait_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/wait/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Wait_MinPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/wait/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/min to the batch object.
func (n *System_Cpu_Wait_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Wait{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Wait", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Wait_MinPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/min to the batch object.
func (n *System_Cpu_Wait_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Wait_MinPath extracts the value of the leaf Min from its parent oc.System_Cpu_Wait
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Wait_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Wait) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Min
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/wait/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Wait_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Wait{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Wait", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Wait_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/wait/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Wait_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/wait/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Wait_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Wait{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Wait", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Wait_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/wait/min-time with a ONCE subscription.
func (n *System_Cpu_Wait_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Wait", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Wait_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/wait/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Wait_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/wait/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/min-time to the batch object.
func (n *System_Cpu_Wait_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Wait{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Wait", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Wait_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/min-time to the batch object.
func (n *System_Cpu_Wait_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Wait_MinTimePath extracts the value of the leaf MinTime from its parent oc.System_Cpu_Wait
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Wait_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Wait) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MinTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/state/current-datetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_CurrentDatetimePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, true, false)
	if ok {
		return convertSystem_CurrentDatetimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/state/current-datetime with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_CurrentDatetimePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/state/current-datetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_CurrentDatetimePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_CurrentDatetimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/state/current-datetime with a ONCE subscription.
func (n *System_CurrentDatetimePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/current-datetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_CurrentDatetimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_CurrentDatetimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_CurrentDatetimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/current-datetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_CurrentDatetimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_CurrentDatetimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/state/current-datetime with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_CurrentDatetimePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/state/current-datetime failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/state/current-datetime to the batch object.
func (n *System_CurrentDatetimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/current-datetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_CurrentDatetimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_CurrentDatetimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System", structs[pre], queryPath, true, false)
			qv := convertSystem_CurrentDatetimePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/current-datetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_CurrentDatetimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_CurrentDatetimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/state/current-datetime to the batch object.
func (n *System_CurrentDatetimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_CurrentDatetimePath extracts the value of the leaf CurrentDatetime from its parent oc.System
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_CurrentDatetimePath(t testing.TB, md *genutil.Metadata, parent *oc.System) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CurrentDatetime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_DnsPath) Lookup(t testing.TB) *oc.QualifiedSystem_Dns {
	t.Helper()
	goStruct := &oc.System_Dns{}
	md, ok := oc.Lookup(t, n, "System_Dns", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Dns{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_DnsPath) Get(t testing.TB) *oc.System_Dns {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_DnsPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Dns {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Dns
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Dns{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns with a ONCE subscription.
func (n *System_DnsPathAny) Get(t testing.TB) []*oc.System_Dns {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Dns
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_DnsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Dns {
	t.Helper()
	c := &oc.CollectionSystem_Dns{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Dns) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Dns{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Dns)))
		return false
	})
	return c
}

func watch_System_DnsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Dns) bool) *oc.System_DnsWatcher {
	t.Helper()
	w := &oc.System_DnsWatcher{}
	gs := &oc.System_Dns{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Dns{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Dns)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_DnsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Dns) bool) *oc.System_DnsWatcher {
	t.Helper()
	return watch_System_DnsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_DnsPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Dns) *oc.QualifiedSystem_Dns {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Dns) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns to the batch object.
func (n *System_DnsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_DnsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Dns {
	t.Helper()
	c := &oc.CollectionSystem_Dns{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Dns) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_DnsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Dns) bool) *oc.System_DnsWatcher {
	t.Helper()
	w := &oc.System_DnsWatcher{}
	structs := map[string]*oc.System_Dns{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Dns{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Dns)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_DnsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Dns) bool) *oc.System_DnsWatcher {
	t.Helper()
	return watch_System_DnsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns to the batch object.
func (n *System_DnsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntryPath) Lookup(t testing.TB) *oc.QualifiedSystem_Dns_HostEntry {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Dns_HostEntry{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntryPath) Get(t testing.TB) *oc.System_Dns_HostEntry {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntryPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Dns_HostEntry {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Dns_HostEntry
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Dns_HostEntry{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry with a ONCE subscription.
func (n *System_Dns_HostEntryPathAny) Get(t testing.TB) []*oc.System_Dns_HostEntry {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Dns_HostEntry
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntryPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Dns_HostEntry {
	t.Helper()
	c := &oc.CollectionSystem_Dns_HostEntry{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Dns_HostEntry) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Dns_HostEntry{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Dns_HostEntry)))
		return false
	})
	return c
}

func watch_System_Dns_HostEntryPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Dns_HostEntry) bool) *oc.System_Dns_HostEntryWatcher {
	t.Helper()
	w := &oc.System_Dns_HostEntryWatcher{}
	gs := &oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_HostEntry", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Dns_HostEntry{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Dns_HostEntry)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntryPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Dns_HostEntry) bool) *oc.System_Dns_HostEntryWatcher {
	t.Helper()
	return watch_System_Dns_HostEntryPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/host-entries/host-entry with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_HostEntryPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Dns_HostEntry) *oc.QualifiedSystem_Dns_HostEntry {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Dns_HostEntry) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/host-entries/host-entry failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry to the batch object.
func (n *System_Dns_HostEntryPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntryPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Dns_HostEntry {
	t.Helper()
	c := &oc.CollectionSystem_Dns_HostEntry{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Dns_HostEntry) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntryPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Dns_HostEntry) bool) *oc.System_Dns_HostEntryWatcher {
	t.Helper()
	w := &oc.System_Dns_HostEntryWatcher{}
	structs := map[string]*oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_HostEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_HostEntry", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Dns_HostEntry{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Dns_HostEntry)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntryPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Dns_HostEntry) bool) *oc.System_Dns_HostEntryWatcher {
	t.Helper()
	return watch_System_Dns_HostEntryPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry to the batch object.
func (n *System_Dns_HostEntryPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/alias with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_AliasPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, false)
	if ok {
		return convertSystem_Dns_HostEntry_AliasPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/alias with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_AliasPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/alias with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_AliasPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_HostEntry_AliasPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/alias with a ONCE subscription.
func (n *System_Dns_HostEntry_AliasPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/alias with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_AliasPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_AliasPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_HostEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_HostEntry_AliasPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/alias with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_AliasPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_AliasPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/host-entries/host-entry/state/alias with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_HostEntry_AliasPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/host-entries/host-entry/state/alias failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/alias to the batch object.
func (n *System_Dns_HostEntry_AliasPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/alias with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_AliasPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_AliasPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_HostEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_HostEntry", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_HostEntry_AliasPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/alias with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_AliasPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_AliasPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/alias to the batch object.
func (n *System_Dns_HostEntry_AliasPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_HostEntry_AliasPath extracts the value of the leaf Alias from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_HostEntry_AliasPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Alias
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_HostnamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, false)
	if ok {
		return convertSystem_Dns_HostEntry_HostnamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_HostnamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_HostnamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_HostEntry_HostnamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a ONCE subscription.
func (n *System_Dns_HostEntry_HostnamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_HostnamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_HostnamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_HostEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_HostEntry_HostnamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_HostnamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_HostnamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_HostEntry_HostnamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/host-entries/host-entry/state/hostname failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/hostname to the batch object.
func (n *System_Dns_HostEntry_HostnamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_HostnamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_HostnamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_HostEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_HostEntry", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_HostEntry_HostnamePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_HostnamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_HostnamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/hostname to the batch object.
func (n *System_Dns_HostEntry_HostnamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_HostEntry_HostnamePath extracts the value of the leaf Hostname from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Dns_HostEntry_HostnamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Hostname
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, false)
	if ok {
		return convertSystem_Dns_HostEntry_Ipv4AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_HostEntry_Ipv4AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a ONCE subscription.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_Ipv4AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_HostEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_HostEntry_Ipv4AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_Ipv4AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address to the batch object.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_Ipv4AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_HostEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_HostEntry", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_HostEntry_Ipv4AddressPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_Ipv4AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address to the batch object.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_HostEntry_Ipv4AddressPath extracts the value of the leaf Ipv4Address from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_HostEntry_Ipv4AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Ipv4Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, false)
	if ok {
		return convertSystem_Dns_HostEntry_Ipv6AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_HostEntry_Ipv6AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a ONCE subscription.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_Ipv6AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_HostEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_HostEntry_Ipv6AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_Ipv6AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address to the batch object.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_Ipv6AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_HostEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_HostEntry", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_HostEntry_Ipv6AddressPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_Ipv6AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address to the batch object.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_HostEntry_Ipv6AddressPath extracts the value of the leaf Ipv6Address from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_HostEntry_Ipv6AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Ipv6Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/state/search with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_SearchPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns{}
	md, ok := oc.Lookup(t, n, "System_Dns", goStruct, true, false)
	if ok {
		return convertSystem_Dns_SearchPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/state/search with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_SearchPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/state/search with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_SearchPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_SearchPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/state/search with a ONCE subscription.
func (n *System_Dns_SearchPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/state/search with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_SearchPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_SearchPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.System_Dns{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_SearchPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/state/search with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_SearchPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_SearchPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/state/search with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_SearchPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/state/search failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/state/search to the batch object.
func (n *System_Dns_SearchPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/state/search with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_SearchPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_SearchPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.System_Dns{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_SearchPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/state/search with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_SearchPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_SearchPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/state/search to the batch object.
func (n *System_Dns_SearchPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_SearchPath extracts the value of the leaf Search from its parent oc.System_Dns
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_SearchPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Search
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/servers/server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_ServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_Dns_Server {
	t.Helper()
	goStruct := &oc.System_Dns_Server{}
	md, ok := oc.Lookup(t, n, "System_Dns_Server", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Dns_Server{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/servers/server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_ServerPath) Get(t testing.TB) *oc.System_Dns_Server {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/servers/server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_ServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Dns_Server {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Dns_Server
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_Server", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Dns_Server{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/servers/server with a ONCE subscription.
func (n *System_Dns_ServerPathAny) Get(t testing.TB) []*oc.System_Dns_Server {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Dns_Server
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_ServerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Dns_Server {
	t.Helper()
	c := &oc.CollectionSystem_Dns_Server{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Dns_Server) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Dns_Server{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Dns_Server)))
		return false
	})
	return c
}

func watch_System_Dns_ServerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Dns_Server) bool) *oc.System_Dns_ServerWatcher {
	t.Helper()
	w := &oc.System_Dns_ServerWatcher{}
	gs := &oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_Server", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Dns_Server{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Dns_Server)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_ServerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Dns_Server) bool) *oc.System_Dns_ServerWatcher {
	t.Helper()
	return watch_System_Dns_ServerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/servers/server with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_ServerPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Dns_Server) *oc.QualifiedSystem_Dns_Server {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Dns_Server) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/servers/server failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/servers/server to the batch object.
func (n *System_Dns_ServerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_ServerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Dns_Server {
	t.Helper()
	c := &oc.CollectionSystem_Dns_Server{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Dns_Server) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_ServerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Dns_Server) bool) *oc.System_Dns_ServerWatcher {
	t.Helper()
	w := &oc.System_Dns_ServerWatcher{}
	structs := map[string]*oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_Server{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_Server", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Dns_Server{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Dns_Server)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_ServerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Dns_Server) bool) *oc.System_Dns_ServerWatcher {
	t.Helper()
	return watch_System_Dns_ServerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/servers/server to the batch object.
func (n *System_Dns_ServerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/dns/servers/server/state/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_Server_AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Dns_Server{}
	md, ok := oc.Lookup(t, n, "System_Dns_Server", goStruct, true, false)
	if ok {
		return convertSystem_Dns_Server_AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/servers/server/state/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_Server_AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/servers/server/state/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_Server_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_Server_AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/servers/server/state/address with a ONCE subscription.
func (n *System_Dns_Server_AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_Server_AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_Server_AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_Server_AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_Server_AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Dns_Server_AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/servers/server/state/address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_Server_AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/servers/server/state/address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/servers/server/state/address to the batch object.
func (n *System_Dns_Server_AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_Server_AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_Server_AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_Server{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_Server", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_Server_AddressPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_Server_AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Dns_Server_AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/servers/server/state/address to the batch object.
func (n *System_Dns_Server_AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_Server_AddressPath extracts the value of the leaf Address from its parent oc.System_Dns_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Dns_Server_AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_Server) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/servers/server/state/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_Server_PortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Dns_Server{}
	md, ok := oc.Lookup(t, n, "System_Dns_Server", goStruct, true, false)
	if ok {
		return convertSystem_Dns_Server_PortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPort())
}

// Get fetches the value at /openconfig-system/system/dns/servers/server/state/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_Server_PortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/servers/server/state/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_Server_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_Server_PortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/servers/server/state/port with a ONCE subscription.
func (n *System_Dns_Server_PortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server/state/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_Server_PortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_Server_PortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_Server_PortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server/state/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_Server_PortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Dns_Server_PortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/servers/server/state/port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_Server_PortPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/servers/server/state/port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/servers/server/state/port to the batch object.
func (n *System_Dns_Server_PortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server/state/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_Server_PortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_Server_PortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_Server{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_Server", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_Server_PortPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server/state/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_Server_PortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Dns_Server_PortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/servers/server/state/port to the batch object.
func (n *System_Dns_Server_PortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_Server_PortPath extracts the value of the leaf Port from its parent oc.System_Dns_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Dns_Server_PortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_Server) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Port
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/state/domain-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_DomainNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, true, false)
	if ok {
		return convertSystem_DomainNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/state/domain-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_DomainNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/state/domain-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_DomainNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_DomainNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/state/domain-name with a ONCE subscription.
func (n *System_DomainNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/domain-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_DomainNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_DomainNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_DomainNamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/domain-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_DomainNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_DomainNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/state/domain-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_DomainNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/state/domain-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/state/domain-name to the batch object.
func (n *System_DomainNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/domain-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_DomainNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_DomainNamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System", structs[pre], queryPath, true, false)
			qv := convertSystem_DomainNamePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/domain-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_DomainNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_DomainNamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/state/domain-name to the batch object.
func (n *System_DomainNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_DomainNamePath extracts the value of the leaf DomainName from its parent oc.System
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_DomainNamePath(t testing.TB, md *genutil.Metadata, parent *oc.System) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DomainName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_GrpcServer {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_GrpcServer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServerPath) Get(t testing.TB) *oc.System_GrpcServer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_GrpcServer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_GrpcServer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_GrpcServer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription.
func (n *System_GrpcServerPathAny) Get(t testing.TB) []*oc.System_GrpcServer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_GrpcServer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GrpcServer {
	t.Helper()
	c := &oc.CollectionSystem_GrpcServer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GrpcServer) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_GrpcServer{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_GrpcServer)))
		return false
	})
	return c
}

func watch_System_GrpcServerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer) bool) *oc.System_GrpcServerWatcher {
	t.Helper()
	w := &oc.System_GrpcServerWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_GrpcServer{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GrpcServer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer) bool) *oc.System_GrpcServerWatcher {
	t.Helper()
	return watch_System_GrpcServerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServerPath) Await(t testing.TB, timeout time.Duration, val *oc.System_GrpcServer) *oc.QualifiedSystem_GrpcServer {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_GrpcServer) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server to the batch object.
func (n *System_GrpcServerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GrpcServer {
	t.Helper()
	c := &oc.CollectionSystem_GrpcServer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GrpcServer) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer) bool) *oc.System_GrpcServerWatcher {
	t.Helper()
	w := &oc.System_GrpcServerWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_GrpcServer{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GrpcServer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer) bool) *oc.System_GrpcServerWatcher {
	t.Helper()
	return watch_System_GrpcServerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server to the batch object.
func (n *System_GrpcServerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_CertificateIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_CertificateIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_CertificateIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_CertificateIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_CertificateIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a ONCE subscription.
func (n *System_GrpcServer_CertificateIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CertificateIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CertificateIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_CertificateIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CertificateIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_CertificateIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_CertificateIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id to the batch object.
func (n *System_GrpcServer_CertificateIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CertificateIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CertificateIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_CertificateIdPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CertificateIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_CertificateIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id to the batch object.
func (n *System_GrpcServer_CertificateIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_CertificateIdPath extracts the value of the leaf CertificateId from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcServer_CertificateIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CertificateId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/enable with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_EnablePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_EnablePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/enable with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_EnablePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/enable with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_EnablePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_EnablePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/enable with a ONCE subscription.
func (n *System_GrpcServer_EnablePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/enable with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_EnablePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_EnablePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_EnablePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/enable with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_EnablePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_GrpcServer_EnablePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/enable with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_EnablePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/enable failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/enable to the batch object.
func (n *System_GrpcServer_EnablePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/enable with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_EnablePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_EnablePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_EnablePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/enable with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_EnablePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_GrpcServer_EnablePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/enable to the batch object.
func (n *System_GrpcServer_EnablePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_EnablePath extracts the value of the leaf Enable from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_GrpcServer_EnablePath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enable
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_ListenAddressesPath) Lookup(t testing.TB) *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_ListenAddressesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_ListenAddressesPath) Get(t testing.TB) []oc.System_GrpcServer_ListenAddresses_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_ListenAddressesPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_ListenAddressesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses with a ONCE subscription.
func (n *System_GrpcServer_ListenAddressesPathAny) Get(t testing.TB) [][]oc.System_GrpcServer_ListenAddresses_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.System_GrpcServer_ListenAddresses_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_ListenAddressesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GrpcServer_ListenAddresses_UnionSlice {
	t.Helper()
	c := &oc.CollectionSystem_GrpcServer_ListenAddresses_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_ListenAddressesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) bool) *oc.System_GrpcServer_ListenAddresses_UnionSliceWatcher {
	t.Helper()
	w := &oc.System_GrpcServer_ListenAddresses_UnionSliceWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_ListenAddressesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_ListenAddressesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) bool) *oc.System_GrpcServer_ListenAddresses_UnionSliceWatcher {
	t.Helper()
	return watch_System_GrpcServer_ListenAddressesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_ListenAddressesPath) Await(t testing.TB, timeout time.Duration, val []oc.System_GrpcServer_ListenAddresses_Union) *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses to the batch object.
func (n *System_GrpcServer_ListenAddressesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_ListenAddressesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GrpcServer_ListenAddresses_UnionSlice {
	t.Helper()
	c := &oc.CollectionSystem_GrpcServer_ListenAddresses_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_ListenAddressesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) bool) *oc.System_GrpcServer_ListenAddresses_UnionSliceWatcher {
	t.Helper()
	w := &oc.System_GrpcServer_ListenAddresses_UnionSliceWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_ListenAddressesPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_ListenAddressesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) bool) *oc.System_GrpcServer_ListenAddresses_UnionSliceWatcher {
	t.Helper()
	return watch_System_GrpcServer_ListenAddressesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/listen-addresses to the batch object.
func (n *System_GrpcServer_ListenAddressesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_ListenAddressesPath extracts the value of the leaf ListenAddresses from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice.
func convertSystem_GrpcServer_ListenAddressesPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice{
		Metadata: md,
	}
	val := parent.ListenAddresses
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_MetadataAuthenticationPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_MetadataAuthenticationPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_MetadataAuthenticationPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_MetadataAuthenticationPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_MetadataAuthenticationPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication with a ONCE subscription.
func (n *System_GrpcServer_MetadataAuthenticationPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_MetadataAuthenticationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_MetadataAuthenticationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_MetadataAuthenticationPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_MetadataAuthenticationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_GrpcServer_MetadataAuthenticationPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_MetadataAuthenticationPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication to the batch object.
func (n *System_GrpcServer_MetadataAuthenticationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_MetadataAuthenticationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_MetadataAuthenticationPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_MetadataAuthenticationPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_MetadataAuthenticationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_GrpcServer_MetadataAuthenticationPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/metadata-authentication to the batch object.
func (n *System_GrpcServer_MetadataAuthenticationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_MetadataAuthenticationPath extracts the value of the leaf MetadataAuthentication from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_GrpcServer_MetadataAuthenticationPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.MetadataAuthentication
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_NamePath(t, md, goStruct)
	}
	return (&oc.QualifiedString{
		Metadata: md,
	}).SetVal(goStruct.GetName())
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/name with a ONCE subscription.
func (n *System_GrpcServer_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/name to the batch object.
func (n *System_GrpcServer_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_NamePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/name to the batch object.
func (n *System_GrpcServer_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_NamePath extracts the value of the leaf Name from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcServer_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Name
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/network-instance with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_NetworkInstancePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_NetworkInstancePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/network-instance with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_NetworkInstancePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/network-instance with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_NetworkInstancePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_NetworkInstancePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/network-instance with a ONCE subscription.
func (n *System_GrpcServer_NetworkInstancePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/network-instance with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_NetworkInstancePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_NetworkInstancePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_NetworkInstancePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/network-instance with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_NetworkInstancePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_NetworkInstancePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/network-instance with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_NetworkInstancePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/network-instance failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/network-instance to the batch object.
func (n *System_GrpcServer_NetworkInstancePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/network-instance with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_NetworkInstancePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_NetworkInstancePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_NetworkInstancePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/network-instance with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_NetworkInstancePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_NetworkInstancePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/network-instance to the batch object.
func (n *System_GrpcServer_NetworkInstancePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_NetworkInstancePath extracts the value of the leaf NetworkInstance from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcServer_NetworkInstancePath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.NetworkInstance
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
