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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_SoftwareInterrupt_AvgPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_SoftwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_SoftwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_SoftwareInterrupt_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_SoftwareInterrupt_AvgPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_SoftwareInterrupt_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_SoftwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_SoftwareInterrupt_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/avg with a ONCE subscription.
func (n *System_Cpu_SoftwareInterrupt_AvgPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_SoftwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_SoftwareInterrupt_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/software-interrupt/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_SoftwareInterrupt_AvgPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/software-interrupt/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/avg to the batch object.
func (n *System_Cpu_SoftwareInterrupt_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_SoftwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_SoftwareInterrupt_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/avg to the batch object.
func (n *System_Cpu_SoftwareInterrupt_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_SoftwareInterrupt_AvgPath extracts the value of the leaf Avg from its parent oc.System_Cpu_SoftwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_SoftwareInterrupt_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_SoftwareInterrupt) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_SoftwareInterrupt_InstantPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_SoftwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_SoftwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_SoftwareInterrupt_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_SoftwareInterrupt_InstantPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_SoftwareInterrupt_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_SoftwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_SoftwareInterrupt_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/instant with a ONCE subscription.
func (n *System_Cpu_SoftwareInterrupt_InstantPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_SoftwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_SoftwareInterrupt_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/software-interrupt/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_SoftwareInterrupt_InstantPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/software-interrupt/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/instant to the batch object.
func (n *System_Cpu_SoftwareInterrupt_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_SoftwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_SoftwareInterrupt_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/instant to the batch object.
func (n *System_Cpu_SoftwareInterrupt_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_SoftwareInterrupt_InstantPath extracts the value of the leaf Instant from its parent oc.System_Cpu_SoftwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_SoftwareInterrupt_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_SoftwareInterrupt) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_SoftwareInterrupt_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_SoftwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_SoftwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_SoftwareInterrupt_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_SoftwareInterrupt_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_SoftwareInterrupt_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_SoftwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_SoftwareInterrupt_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/interval with a ONCE subscription.
func (n *System_Cpu_SoftwareInterrupt_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_SoftwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_SoftwareInterrupt_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/software-interrupt/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_SoftwareInterrupt_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/software-interrupt/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/interval to the batch object.
func (n *System_Cpu_SoftwareInterrupt_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_SoftwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_SoftwareInterrupt_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/interval to the batch object.
func (n *System_Cpu_SoftwareInterrupt_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_SoftwareInterrupt_IntervalPath extracts the value of the leaf Interval from its parent oc.System_Cpu_SoftwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_SoftwareInterrupt_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_SoftwareInterrupt) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_SoftwareInterrupt_MaxPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_SoftwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_SoftwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_SoftwareInterrupt_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_SoftwareInterrupt_MaxPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_SoftwareInterrupt_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_SoftwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_SoftwareInterrupt_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max with a ONCE subscription.
func (n *System_Cpu_SoftwareInterrupt_MaxPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_SoftwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_SoftwareInterrupt_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_SoftwareInterrupt_MaxPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/software-interrupt/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/max to the batch object.
func (n *System_Cpu_SoftwareInterrupt_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_SoftwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_SoftwareInterrupt_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/max to the batch object.
func (n *System_Cpu_SoftwareInterrupt_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_SoftwareInterrupt_MaxPath extracts the value of the leaf Max from its parent oc.System_Cpu_SoftwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_SoftwareInterrupt_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_SoftwareInterrupt) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_SoftwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_SoftwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_SoftwareInterrupt_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_SoftwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_SoftwareInterrupt_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time with a ONCE subscription.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_SoftwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_SoftwareInterrupt_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time to the batch object.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_SoftwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_SoftwareInterrupt_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/max-time to the batch object.
func (n *System_Cpu_SoftwareInterrupt_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_SoftwareInterrupt_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.System_Cpu_SoftwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_SoftwareInterrupt_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_SoftwareInterrupt) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_SoftwareInterrupt_MinPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_SoftwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_SoftwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_SoftwareInterrupt_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_SoftwareInterrupt_MinPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_SoftwareInterrupt_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_SoftwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_SoftwareInterrupt_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min with a ONCE subscription.
func (n *System_Cpu_SoftwareInterrupt_MinPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_SoftwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_SoftwareInterrupt_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_SoftwareInterrupt_MinPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/software-interrupt/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/min to the batch object.
func (n *System_Cpu_SoftwareInterrupt_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_SoftwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_SoftwareInterrupt_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/min to the batch object.
func (n *System_Cpu_SoftwareInterrupt_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_SoftwareInterrupt_MinPath extracts the value of the leaf Min from its parent oc.System_Cpu_SoftwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_SoftwareInterrupt_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_SoftwareInterrupt) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_SoftwareInterrupt_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_SoftwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_SoftwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_SoftwareInterrupt_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_SoftwareInterrupt_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_SoftwareInterrupt_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_SoftwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_SoftwareInterrupt_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time with a ONCE subscription.
func (n *System_Cpu_SoftwareInterrupt_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_SoftwareInterrupt", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_SoftwareInterrupt_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_SoftwareInterrupt_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time to the batch object.
func (n *System_Cpu_SoftwareInterrupt_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_SoftwareInterrupt_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_SoftwareInterrupt_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_SoftwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_SoftwareInterrupt{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_SoftwareInterrupt", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_SoftwareInterrupt_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_SoftwareInterrupt_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_SoftwareInterrupt_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/software-interrupt/min-time to the batch object.
func (n *System_Cpu_SoftwareInterrupt_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_SoftwareInterrupt_MinTimePath extracts the value of the leaf MinTime from its parent oc.System_Cpu_SoftwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_SoftwareInterrupt_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_SoftwareInterrupt) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/total with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_TotalPath) Lookup(t testing.TB) *oc.QualifiedSystem_Cpu_Total {
	t.Helper()
	goStruct := &oc.System_Cpu_Total{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Total", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Cpu_Total{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/total with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_TotalPath) Get(t testing.TB) *oc.System_Cpu_Total {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/total with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_TotalPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Cpu_Total {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Cpu_Total
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Total{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Total", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Cpu_Total{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/total with a ONCE subscription.
func (n *System_Cpu_TotalPathAny) Get(t testing.TB) []*oc.System_Cpu_Total {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Cpu_Total
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_TotalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Total {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Total{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Total) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Cpu_Total{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Cpu_Total)))
		return false
	})
	return c
}

func watch_System_Cpu_TotalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Total) bool) *oc.System_Cpu_TotalWatcher {
	t.Helper()
	w := &oc.System_Cpu_TotalWatcher{}
	gs := &oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Total", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Cpu_Total{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_Total)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_TotalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Total) bool) *oc.System_Cpu_TotalWatcher {
	t.Helper()
	return watch_System_Cpu_TotalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/total with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_TotalPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Cpu_Total) *oc.QualifiedSystem_Cpu_Total {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Cpu_Total) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/total failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total to the batch object.
func (n *System_Cpu_TotalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_TotalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Total {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Total{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Total) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_TotalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Total) bool) *oc.System_Cpu_TotalWatcher {
	t.Helper()
	w := &oc.System_Cpu_TotalWatcher{}
	structs := map[string]*oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Total{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Total", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Cpu_Total{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_Total)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_TotalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Total) bool) *oc.System_Cpu_TotalWatcher {
	t.Helper()
	return watch_System_Cpu_TotalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total to the batch object.
func (n *System_Cpu_TotalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/total/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Total_AvgPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Total{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Total", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Total_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/total/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Total_AvgPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/total/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Total_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Total{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Total", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Total_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/total/avg with a ONCE subscription.
func (n *System_Cpu_Total_AvgPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Total", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Total_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Total_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/total/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Total_AvgPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/total/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/avg to the batch object.
func (n *System_Cpu_Total_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Total{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Total", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Total_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Total_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/avg to the batch object.
func (n *System_Cpu_Total_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Total_AvgPath extracts the value of the leaf Avg from its parent oc.System_Cpu_Total
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Total_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Total) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/total/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Total_InstantPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Total{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Total", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Total_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/total/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Total_InstantPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/total/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Total_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Total{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Total", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Total_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/total/instant with a ONCE subscription.
func (n *System_Cpu_Total_InstantPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Total", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Total_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Total_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/total/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Total_InstantPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/total/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/instant to the batch object.
func (n *System_Cpu_Total_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Total{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Total", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Total_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Total_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/instant to the batch object.
func (n *System_Cpu_Total_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Total_InstantPath extracts the value of the leaf Instant from its parent oc.System_Cpu_Total
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Total_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Total) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/total/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Total_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Total{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Total", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Total_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/total/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Total_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/total/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Total_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Total{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Total", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Total_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/total/interval with a ONCE subscription.
func (n *System_Cpu_Total_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Total", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Total_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Total_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/total/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Total_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/total/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/interval to the batch object.
func (n *System_Cpu_Total_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Total{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Total", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Total_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Total_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/interval to the batch object.
func (n *System_Cpu_Total_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Total_IntervalPath extracts the value of the leaf Interval from its parent oc.System_Cpu_Total
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Total_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Total) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/total/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Total_MaxPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Total{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Total", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Total_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/total/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Total_MaxPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/total/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Total_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Total{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Total", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Total_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/total/max with a ONCE subscription.
func (n *System_Cpu_Total_MaxPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Total", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Total_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Total_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/total/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Total_MaxPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/total/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/max to the batch object.
func (n *System_Cpu_Total_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Total{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Total", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Total_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Total_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/max to the batch object.
func (n *System_Cpu_Total_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Total_MaxPath extracts the value of the leaf Max from its parent oc.System_Cpu_Total
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Total_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Total) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/total/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Total_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Total{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Total", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Total_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/total/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Total_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/total/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Total_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Total{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Total", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Total_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/total/max-time with a ONCE subscription.
func (n *System_Cpu_Total_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Total", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Total_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Total_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/total/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Total_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/total/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/max-time to the batch object.
func (n *System_Cpu_Total_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Total{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Total", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Total_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Total_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/max-time to the batch object.
func (n *System_Cpu_Total_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Total_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.System_Cpu_Total
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Total_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Total) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/total/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Total_MinPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Total{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Total", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Total_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/total/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Total_MinPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/total/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Total_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Total{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Total", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Total_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/total/min with a ONCE subscription.
func (n *System_Cpu_Total_MinPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Total", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Total_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Total_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/total/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Total_MinPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/total/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/min to the batch object.
func (n *System_Cpu_Total_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Total{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Total", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Total_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Total_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/min to the batch object.
func (n *System_Cpu_Total_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Total_MinPath extracts the value of the leaf Min from its parent oc.System_Cpu_Total
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Total_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Total) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/total/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Total_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Total{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Total", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Total_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/total/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Total_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/total/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Total_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Total{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Total", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Total_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/total/min-time with a ONCE subscription.
func (n *System_Cpu_Total_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Total", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Total_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Total_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/total/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Total_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/total/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/min-time to the batch object.
func (n *System_Cpu_Total_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/total/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Total_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Total_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_Total{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_Total{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Total", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_Total_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/total/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Total_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Total_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/total/min-time to the batch object.
func (n *System_Cpu_Total_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Total_MinTimePath extracts the value of the leaf MinTime from its parent oc.System_Cpu_Total
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Total_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Total) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/user with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_UserPath) Lookup(t testing.TB) *oc.QualifiedSystem_Cpu_User {
	t.Helper()
	goStruct := &oc.System_Cpu_User{}
	md, ok := oc.Lookup(t, n, "System_Cpu_User", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Cpu_User{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/user with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_UserPath) Get(t testing.TB) *oc.System_Cpu_User {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/user with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_UserPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Cpu_User {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Cpu_User
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_User", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Cpu_User{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/user with a ONCE subscription.
func (n *System_Cpu_UserPathAny) Get(t testing.TB) []*oc.System_Cpu_User {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Cpu_User
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_UserPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_User {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_User{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_User) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Cpu_User{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Cpu_User)))
		return false
	})
	return c
}

func watch_System_Cpu_UserPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_User) bool) *oc.System_Cpu_UserWatcher {
	t.Helper()
	w := &oc.System_Cpu_UserWatcher{}
	gs := &oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_User", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Cpu_User{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_User)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_UserPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_User) bool) *oc.System_Cpu_UserWatcher {
	t.Helper()
	return watch_System_Cpu_UserPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/user with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_UserPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Cpu_User) *oc.QualifiedSystem_Cpu_User {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Cpu_User) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/user failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user to the batch object.
func (n *System_Cpu_UserPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_UserPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_User {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_User{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_User) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_UserPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_User) bool) *oc.System_Cpu_UserWatcher {
	t.Helper()
	w := &oc.System_Cpu_UserWatcher{}
	structs := map[string]*oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_User", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Cpu_User{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_User)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_UserPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_User) bool) *oc.System_Cpu_UserWatcher {
	t.Helper()
	return watch_System_Cpu_UserPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user to the batch object.
func (n *System_Cpu_UserPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/user/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_User_AvgPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_User{}
	md, ok := oc.Lookup(t, n, "System_Cpu_User", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_User_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/user/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_User_AvgPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/user/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_User_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_User_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/user/avg with a ONCE subscription.
func (n *System_Cpu_User_AvgPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_User_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_User_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/user/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_User_AvgPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/user/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/avg to the batch object.
func (n *System_Cpu_User_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_User_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_User_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/avg to the batch object.
func (n *System_Cpu_User_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_User_AvgPath extracts the value of the leaf Avg from its parent oc.System_Cpu_User
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_User_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_User) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/user/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_User_InstantPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_User{}
	md, ok := oc.Lookup(t, n, "System_Cpu_User", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_User_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/user/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_User_InstantPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/user/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_User_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_User_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/user/instant with a ONCE subscription.
func (n *System_Cpu_User_InstantPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_User_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_User_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/user/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_User_InstantPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/user/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/instant to the batch object.
func (n *System_Cpu_User_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_User_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_User_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/instant to the batch object.
func (n *System_Cpu_User_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_User_InstantPath extracts the value of the leaf Instant from its parent oc.System_Cpu_User
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_User_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_User) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/user/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_User_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_User{}
	md, ok := oc.Lookup(t, n, "System_Cpu_User", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_User_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/user/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_User_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/user/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_User_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_User_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/user/interval with a ONCE subscription.
func (n *System_Cpu_User_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_User_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_User_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/user/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_User_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/user/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/interval to the batch object.
func (n *System_Cpu_User_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_User_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_User_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/interval to the batch object.
func (n *System_Cpu_User_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_User_IntervalPath extracts the value of the leaf Interval from its parent oc.System_Cpu_User
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_User_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_User) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/user/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_User_MaxPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_User{}
	md, ok := oc.Lookup(t, n, "System_Cpu_User", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_User_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/user/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_User_MaxPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/user/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_User_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_User_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/user/max with a ONCE subscription.
func (n *System_Cpu_User_MaxPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_User_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_User_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/user/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_User_MaxPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/user/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/max to the batch object.
func (n *System_Cpu_User_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_User_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_User_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/max to the batch object.
func (n *System_Cpu_User_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_User_MaxPath extracts the value of the leaf Max from its parent oc.System_Cpu_User
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_User_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_User) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/user/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_User_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_User{}
	md, ok := oc.Lookup(t, n, "System_Cpu_User", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_User_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/user/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_User_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/user/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_User_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_User_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/user/max-time with a ONCE subscription.
func (n *System_Cpu_User_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_User_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_User_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/user/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_User_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/user/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/max-time to the batch object.
func (n *System_Cpu_User_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_User_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_User_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/max-time to the batch object.
func (n *System_Cpu_User_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_User_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.System_Cpu_User
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_User_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_User) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/user/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_User_MinPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_User{}
	md, ok := oc.Lookup(t, n, "System_Cpu_User", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_User_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/user/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_User_MinPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/user/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_User_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_User_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/user/min with a ONCE subscription.
func (n *System_Cpu_User_MinPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_User_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_User_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/user/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_User_MinPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/user/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/min to the batch object.
func (n *System_Cpu_User_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_User_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_User_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/min to the batch object.
func (n *System_Cpu_User_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_User_MinPath extracts the value of the leaf Min from its parent oc.System_Cpu_User
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_User_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_User) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/user/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_User_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_User{}
	md, ok := oc.Lookup(t, n, "System_Cpu_User", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_User_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/user/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_User_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/user/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_User_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_User_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/user/min-time with a ONCE subscription.
func (n *System_Cpu_User_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_User_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_User_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/user/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_User_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/user/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/min-time to the batch object.
func (n *System_Cpu_User_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/user/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_User_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_User_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Cpu_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Cpu_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Cpu_User_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/user/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_User_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_User_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/user/min-time to the batch object.
func (n *System_Cpu_User_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_User_MinTimePath extracts the value of the leaf MinTime from its parent oc.System_Cpu_User
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_User_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_User) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/wait with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_WaitPath) Lookup(t testing.TB) *oc.QualifiedSystem_Cpu_Wait {
	t.Helper()
	goStruct := &oc.System_Cpu_Wait{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Wait", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Cpu_Wait{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/wait with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_WaitPath) Get(t testing.TB) *oc.System_Cpu_Wait {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/wait with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_WaitPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Cpu_Wait {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Cpu_Wait
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Wait{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Wait", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Cpu_Wait{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/wait with a ONCE subscription.
func (n *System_Cpu_WaitPathAny) Get(t testing.TB) []*oc.System_Cpu_Wait {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Cpu_Wait
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_WaitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Wait {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Wait{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Wait) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Cpu_Wait{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Cpu_Wait)))
		return false
	})
	return c
}

func watch_System_Cpu_WaitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Wait) bool) *oc.System_Cpu_WaitWatcher {
	t.Helper()
	w := &oc.System_Cpu_WaitWatcher{}
	gs := &oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Wait", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Cpu_Wait{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_Wait)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_WaitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Wait) bool) *oc.System_Cpu_WaitWatcher {
	t.Helper()
	return watch_System_Cpu_WaitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/wait with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_WaitPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Cpu_Wait) *oc.QualifiedSystem_Cpu_Wait {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Cpu_Wait) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/wait failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait to the batch object.
func (n *System_Cpu_WaitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_WaitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Wait {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Wait{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Wait) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_WaitPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Wait) bool) *oc.System_Cpu_WaitWatcher {
	t.Helper()
	w := &oc.System_Cpu_WaitWatcher{}
	structs := map[string]*oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Cpu_Wait", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Cpu_Wait{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_Wait)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_WaitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Wait) bool) *oc.System_Cpu_WaitWatcher {
	t.Helper()
	return watch_System_Cpu_WaitPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait to the batch object.
func (n *System_Cpu_WaitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/wait/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Wait_AvgPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Wait{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Wait", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Wait_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/wait/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Wait_AvgPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/wait/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Wait_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
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
		qv := convertSystem_Cpu_Wait_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/wait/avg with a ONCE subscription.
func (n *System_Cpu_Wait_AvgPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Wait{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Wait", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Cpu_Wait_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/wait/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Wait_AvgPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/wait/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/avg to the batch object.
func (n *System_Cpu_Wait_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/wait/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Wait_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Wait_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
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
			qv := convertSystem_Cpu_Wait_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/wait/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Wait_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Wait_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/wait/avg to the batch object.
func (n *System_Cpu_Wait_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Wait_AvgPath extracts the value of the leaf Avg from its parent oc.System_Cpu_Wait
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Wait_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Wait) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
