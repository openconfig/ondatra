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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_HardwareInterrupt_MinPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_HardwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_HardwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_HardwareInterrupt_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_HardwareInterrupt_MinPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_HardwareInterrupt_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_HardwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_HardwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_HardwareInterrupt_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a ONCE subscription.
func (n *System_Cpu_HardwareInterrupt_MinPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_HardwareInterrupt", gs, queryPath, true, false)
		return convertSystem_Cpu_HardwareInterrupt_MinPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_HardwareInterrupt_MinPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min to the batch object.
func (n *System_Cpu_HardwareInterrupt_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MinPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min to the batch object.
func (n *System_Cpu_HardwareInterrupt_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_HardwareInterrupt_MinPath extracts the value of the leaf Min from its parent oc.System_Cpu_HardwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_HardwareInterrupt_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_HardwareInterrupt) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_HardwareInterrupt{}
	md, ok := oc.Lookup(t, n, "System_Cpu_HardwareInterrupt", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_HardwareInterrupt_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_HardwareInterrupt_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_HardwareInterrupt{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_HardwareInterrupt", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_HardwareInterrupt_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a ONCE subscription.
func (n *System_Cpu_HardwareInterrupt_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_HardwareInterrupt_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_HardwareInterrupt{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_HardwareInterrupt", gs, queryPath, true, false)
		return convertSystem_Cpu_HardwareInterrupt_MinTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time to the batch object.
func (n *System_Cpu_HardwareInterrupt_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_HardwareInterrupt_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_HardwareInterrupt_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_HardwareInterrupt_MinTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/hardware-interrupt/min-time to the batch object.
func (n *System_Cpu_HardwareInterrupt_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_HardwareInterrupt_MinTimePath extracts the value of the leaf MinTime from its parent oc.System_Cpu_HardwareInterrupt
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_HardwareInterrupt_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_HardwareInterrupt) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/idle with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_IdlePath) Lookup(t testing.TB) *oc.QualifiedSystem_Cpu_Idle {
	t.Helper()
	goStruct := &oc.System_Cpu_Idle{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Idle", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Cpu_Idle{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/idle with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_IdlePath) Get(t testing.TB) *oc.System_Cpu_Idle {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/idle with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_IdlePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Cpu_Idle {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Cpu_Idle
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Idle{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Idle", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Cpu_Idle{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/idle with a ONCE subscription.
func (n *System_Cpu_IdlePathAny) Get(t testing.TB) []*oc.System_Cpu_Idle {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Cpu_Idle
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_IdlePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Idle {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Idle{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Idle) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Cpu_Idle{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Cpu_Idle)))
		return false
	})
	return c
}

func watch_System_Cpu_IdlePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Idle) bool) *oc.System_Cpu_IdleWatcher {
	t.Helper()
	w := &oc.System_Cpu_IdleWatcher{}
	gs := &oc.System_Cpu_Idle{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Idle", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Cpu_Idle{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_Idle)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_IdlePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Idle) bool) *oc.System_Cpu_IdleWatcher {
	t.Helper()
	return watch_System_Cpu_IdlePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/idle with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_IdlePath) Await(t testing.TB, timeout time.Duration, val *oc.System_Cpu_Idle) *oc.QualifiedSystem_Cpu_Idle {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Cpu_Idle) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/idle failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle to the batch object.
func (n *System_Cpu_IdlePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_IdlePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Idle {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Idle{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Idle) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_IdlePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Idle) bool) *oc.System_Cpu_IdleWatcher {
	t.Helper()
	return watch_System_Cpu_IdlePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle to the batch object.
func (n *System_Cpu_IdlePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/idle/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Idle_AvgPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Idle{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Idle", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Idle_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/idle/avg with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Idle_AvgPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/idle/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Idle_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Idle{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Idle", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Idle_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/idle/avg with a ONCE subscription.
func (n *System_Cpu_Idle_AvgPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Idle_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Idle{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Idle", gs, queryPath, true, false)
		return convertSystem_Cpu_Idle_AvgPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/idle/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Idle_AvgPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/idle/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/avg to the batch object.
func (n *System_Cpu_Idle_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_AvgPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/avg to the batch object.
func (n *System_Cpu_Idle_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Idle_AvgPath extracts the value of the leaf Avg from its parent oc.System_Cpu_Idle
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Idle_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Idle) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/idle/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Idle_InstantPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Idle{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Idle", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Idle_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/idle/instant with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Idle_InstantPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/idle/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Idle_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Idle{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Idle", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Idle_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/idle/instant with a ONCE subscription.
func (n *System_Cpu_Idle_InstantPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Idle_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Idle{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Idle", gs, queryPath, true, false)
		return convertSystem_Cpu_Idle_InstantPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/idle/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Idle_InstantPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/idle/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/instant to the batch object.
func (n *System_Cpu_Idle_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_InstantPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/instant to the batch object.
func (n *System_Cpu_Idle_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Idle_InstantPath extracts the value of the leaf Instant from its parent oc.System_Cpu_Idle
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Idle_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Idle) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/idle/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Idle_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Idle{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Idle", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Idle_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/idle/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Idle_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/idle/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Idle_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Idle{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Idle", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Idle_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/idle/interval with a ONCE subscription.
func (n *System_Cpu_Idle_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Idle_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Idle{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Idle", gs, queryPath, true, false)
		return convertSystem_Cpu_Idle_IntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/idle/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Idle_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/idle/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/interval to the batch object.
func (n *System_Cpu_Idle_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_IntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/interval to the batch object.
func (n *System_Cpu_Idle_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Idle_IntervalPath extracts the value of the leaf Interval from its parent oc.System_Cpu_Idle
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Idle_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Idle) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/idle/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Idle_MaxPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Idle{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Idle", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Idle_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/idle/max with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Idle_MaxPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/idle/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Idle_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Idle{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Idle", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Idle_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/idle/max with a ONCE subscription.
func (n *System_Cpu_Idle_MaxPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Idle_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Idle{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Idle", gs, queryPath, true, false)
		return convertSystem_Cpu_Idle_MaxPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/idle/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Idle_MaxPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/idle/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/max to the batch object.
func (n *System_Cpu_Idle_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_MaxPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/max to the batch object.
func (n *System_Cpu_Idle_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Idle_MaxPath extracts the value of the leaf Max from its parent oc.System_Cpu_Idle
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Idle_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Idle) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/idle/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Idle_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Idle{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Idle", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Idle_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/idle/max-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Idle_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/idle/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Idle_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Idle{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Idle", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Idle_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/idle/max-time with a ONCE subscription.
func (n *System_Cpu_Idle_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Idle_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Idle{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Idle", gs, queryPath, true, false)
		return convertSystem_Cpu_Idle_MaxTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/idle/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Idle_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/idle/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/max-time to the batch object.
func (n *System_Cpu_Idle_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_MaxTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/max-time to the batch object.
func (n *System_Cpu_Idle_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Idle_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.System_Cpu_Idle
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Idle_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Idle) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/idle/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Idle_MinPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Idle{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Idle", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Idle_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/idle/min with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Idle_MinPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/idle/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Idle_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Idle{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Idle", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Idle_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/idle/min with a ONCE subscription.
func (n *System_Cpu_Idle_MinPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Idle_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Idle{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Idle", gs, queryPath, true, false)
		return convertSystem_Cpu_Idle_MinPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/idle/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Idle_MinPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/idle/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/min to the batch object.
func (n *System_Cpu_Idle_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_MinPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/min to the batch object.
func (n *System_Cpu_Idle_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Idle_MinPath extracts the value of the leaf Min from its parent oc.System_Cpu_Idle
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Idle_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Idle) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/idle/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Idle_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Idle{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Idle", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Idle_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/idle/min-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Idle_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/idle/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Idle_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Idle{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Idle", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Idle_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/idle/min-time with a ONCE subscription.
func (n *System_Cpu_Idle_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Idle_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Idle{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Idle", gs, queryPath, true, false)
		return convertSystem_Cpu_Idle_MinTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/idle/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Idle_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/idle/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/min-time to the batch object.
func (n *System_Cpu_Idle_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/idle/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Idle_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/idle/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Idle_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Idle_MinTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/idle/min-time to the batch object.
func (n *System_Cpu_Idle_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Idle_MinTimePath extracts the value of the leaf MinTime from its parent oc.System_Cpu_Idle
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Idle_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Idle) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_IndexPath) Lookup(t testing.TB) *oc.QualifiedSystem_Cpu_Index_Union {
	t.Helper()
	goStruct := &oc.System_Cpu{}
	md, ok := oc.Lookup(t, n, "System_Cpu", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/index with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_IndexPath) Get(t testing.TB) oc.System_Cpu_Index_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Cpu_Index_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Cpu_Index_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/index with a ONCE subscription.
func (n *System_Cpu_IndexPathAny) Get(t testing.TB) []oc.System_Cpu_Index_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.System_Cpu_Index_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_IndexPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Index_Union {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Index_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Index_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_IndexPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Index_Union) bool) *oc.System_Cpu_Index_UnionWatcher {
	t.Helper()
	w := &oc.System_Cpu_Index_UnionWatcher{}
	gs := &oc.System_Cpu{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu", gs, queryPath, true, false)
		return convertSystem_Cpu_IndexPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_Index_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_IndexPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Index_Union) bool) *oc.System_Cpu_Index_UnionWatcher {
	t.Helper()
	return watch_System_Cpu_IndexPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/index with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_IndexPath) Await(t testing.TB, timeout time.Duration, val oc.System_Cpu_Index_Union) *oc.QualifiedSystem_Cpu_Index_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Cpu_Index_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/index failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/index to the batch object.
func (n *System_Cpu_IndexPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_IndexPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Index_Union {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Index_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Index_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_IndexPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Index_Union) bool) *oc.System_Cpu_Index_UnionWatcher {
	t.Helper()
	return watch_System_Cpu_IndexPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/index to the batch object.
func (n *System_Cpu_IndexPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_IndexPath extracts the value of the leaf Index from its parent oc.System_Cpu
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_Cpu_Index_Union.
func convertSystem_Cpu_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu) *oc.QualifiedSystem_Cpu_Index_Union {
	t.Helper()
	qv := &oc.QualifiedSystem_Cpu_Index_Union{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/kernel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_KernelPath) Lookup(t testing.TB) *oc.QualifiedSystem_Cpu_Kernel {
	t.Helper()
	goStruct := &oc.System_Cpu_Kernel{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Kernel", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Cpu_Kernel{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/kernel with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_KernelPath) Get(t testing.TB) *oc.System_Cpu_Kernel {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/kernel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_KernelPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Cpu_Kernel {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Cpu_Kernel
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Kernel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Kernel", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Cpu_Kernel{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/kernel with a ONCE subscription.
func (n *System_Cpu_KernelPathAny) Get(t testing.TB) []*oc.System_Cpu_Kernel {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Cpu_Kernel
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_KernelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Kernel {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Kernel{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Kernel) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Cpu_Kernel{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Cpu_Kernel)))
		return false
	})
	return c
}

func watch_System_Cpu_KernelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Kernel) bool) *oc.System_Cpu_KernelWatcher {
	t.Helper()
	w := &oc.System_Cpu_KernelWatcher{}
	gs := &oc.System_Cpu_Kernel{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Kernel", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Cpu_Kernel{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_Kernel)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_KernelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Kernel) bool) *oc.System_Cpu_KernelWatcher {
	t.Helper()
	return watch_System_Cpu_KernelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/kernel with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_KernelPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Cpu_Kernel) *oc.QualifiedSystem_Cpu_Kernel {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Cpu_Kernel) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/kernel failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel to the batch object.
func (n *System_Cpu_KernelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_KernelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Kernel {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Kernel{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Kernel) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_KernelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Kernel) bool) *oc.System_Cpu_KernelWatcher {
	t.Helper()
	return watch_System_Cpu_KernelPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel to the batch object.
func (n *System_Cpu_KernelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Kernel_AvgPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Kernel{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Kernel", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Kernel_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/avg with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Kernel_AvgPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Kernel_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Kernel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Kernel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Kernel_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/avg with a ONCE subscription.
func (n *System_Cpu_Kernel_AvgPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Kernel_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Kernel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Kernel", gs, queryPath, true, false)
		return convertSystem_Cpu_Kernel_AvgPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/kernel/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Kernel_AvgPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/kernel/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/avg to the batch object.
func (n *System_Cpu_Kernel_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_AvgPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/avg to the batch object.
func (n *System_Cpu_Kernel_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Kernel_AvgPath extracts the value of the leaf Avg from its parent oc.System_Cpu_Kernel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Kernel_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Kernel) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Kernel_InstantPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Kernel{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Kernel", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Kernel_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/instant with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Kernel_InstantPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Kernel_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Kernel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Kernel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Kernel_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/instant with a ONCE subscription.
func (n *System_Cpu_Kernel_InstantPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Kernel_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Kernel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Kernel", gs, queryPath, true, false)
		return convertSystem_Cpu_Kernel_InstantPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/kernel/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Kernel_InstantPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/kernel/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/instant to the batch object.
func (n *System_Cpu_Kernel_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_InstantPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/instant to the batch object.
func (n *System_Cpu_Kernel_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Kernel_InstantPath extracts the value of the leaf Instant from its parent oc.System_Cpu_Kernel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Kernel_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Kernel) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Kernel_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Kernel{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Kernel", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Kernel_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Kernel_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Kernel_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Kernel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Kernel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Kernel_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/interval with a ONCE subscription.
func (n *System_Cpu_Kernel_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Kernel_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Kernel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Kernel", gs, queryPath, true, false)
		return convertSystem_Cpu_Kernel_IntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/kernel/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Kernel_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/kernel/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/interval to the batch object.
func (n *System_Cpu_Kernel_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_IntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/interval to the batch object.
func (n *System_Cpu_Kernel_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Kernel_IntervalPath extracts the value of the leaf Interval from its parent oc.System_Cpu_Kernel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Kernel_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Kernel) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Kernel_MaxPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Kernel{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Kernel", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Kernel_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/max with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Kernel_MaxPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Kernel_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Kernel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Kernel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Kernel_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/max with a ONCE subscription.
func (n *System_Cpu_Kernel_MaxPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Kernel_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Kernel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Kernel", gs, queryPath, true, false)
		return convertSystem_Cpu_Kernel_MaxPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/kernel/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Kernel_MaxPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/kernel/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/max to the batch object.
func (n *System_Cpu_Kernel_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_MaxPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/max to the batch object.
func (n *System_Cpu_Kernel_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Kernel_MaxPath extracts the value of the leaf Max from its parent oc.System_Cpu_Kernel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Kernel_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Kernel) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Kernel_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Kernel{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Kernel", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Kernel_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/max-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Kernel_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Kernel_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Kernel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Kernel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Kernel_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/max-time with a ONCE subscription.
func (n *System_Cpu_Kernel_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Kernel_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Kernel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Kernel", gs, queryPath, true, false)
		return convertSystem_Cpu_Kernel_MaxTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/kernel/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Kernel_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/kernel/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/max-time to the batch object.
func (n *System_Cpu_Kernel_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_MaxTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/max-time to the batch object.
func (n *System_Cpu_Kernel_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Kernel_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.System_Cpu_Kernel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Kernel_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Kernel) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Kernel_MinPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Kernel{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Kernel", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Kernel_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/min with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Kernel_MinPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Kernel_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Kernel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Kernel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Kernel_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/min with a ONCE subscription.
func (n *System_Cpu_Kernel_MinPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Kernel_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Kernel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Kernel", gs, queryPath, true, false)
		return convertSystem_Cpu_Kernel_MinPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/kernel/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Kernel_MinPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/kernel/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/min to the batch object.
func (n *System_Cpu_Kernel_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_MinPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/min to the batch object.
func (n *System_Cpu_Kernel_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Kernel_MinPath extracts the value of the leaf Min from its parent oc.System_Cpu_Kernel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Kernel_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Kernel) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Kernel_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Kernel{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Kernel", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Kernel_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/kernel/min-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Kernel_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Kernel_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Kernel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Kernel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Kernel_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/kernel/min-time with a ONCE subscription.
func (n *System_Cpu_Kernel_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Kernel_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Kernel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Kernel", gs, queryPath, true, false)
		return convertSystem_Cpu_Kernel_MinTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/kernel/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Kernel_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/kernel/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/min-time to the batch object.
func (n *System_Cpu_Kernel_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/kernel/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Kernel_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/kernel/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Kernel_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Kernel_MinTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/kernel/min-time to the batch object.
func (n *System_Cpu_Kernel_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Kernel_MinTimePath extracts the value of the leaf MinTime from its parent oc.System_Cpu_Kernel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Kernel_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Kernel) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/nice with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_NicePath) Lookup(t testing.TB) *oc.QualifiedSystem_Cpu_Nice {
	t.Helper()
	goStruct := &oc.System_Cpu_Nice{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Nice", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Cpu_Nice{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/nice with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_NicePath) Get(t testing.TB) *oc.System_Cpu_Nice {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/nice with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_NicePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Cpu_Nice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Cpu_Nice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Nice{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Nice", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Cpu_Nice{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/nice with a ONCE subscription.
func (n *System_Cpu_NicePathAny) Get(t testing.TB) []*oc.System_Cpu_Nice {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Cpu_Nice
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/nice with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_NicePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Nice {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Nice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Nice) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Cpu_Nice{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Cpu_Nice)))
		return false
	})
	return c
}

func watch_System_Cpu_NicePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Nice) bool) *oc.System_Cpu_NiceWatcher {
	t.Helper()
	w := &oc.System_Cpu_NiceWatcher{}
	gs := &oc.System_Cpu_Nice{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Nice", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Cpu_Nice{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Cpu_Nice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/nice with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_NicePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Nice) bool) *oc.System_Cpu_NiceWatcher {
	t.Helper()
	return watch_System_Cpu_NicePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/nice with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_NicePath) Await(t testing.TB, timeout time.Duration, val *oc.System_Cpu_Nice) *oc.QualifiedSystem_Cpu_Nice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Cpu_Nice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/nice failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/nice to the batch object.
func (n *System_Cpu_NicePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/nice with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_NicePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Cpu_Nice {
	t.Helper()
	c := &oc.CollectionSystem_Cpu_Nice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Cpu_Nice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/nice with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_NicePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Cpu_Nice) bool) *oc.System_Cpu_NiceWatcher {
	t.Helper()
	return watch_System_Cpu_NicePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/nice to the batch object.
func (n *System_Cpu_NicePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/nice/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Nice_AvgPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Nice{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Nice", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Nice_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/nice/avg with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Nice_AvgPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/nice/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Nice_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Nice{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Nice", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Nice_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/nice/avg with a ONCE subscription.
func (n *System_Cpu_Nice_AvgPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/nice/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Nice_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Nice_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Nice{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Nice", gs, queryPath, true, false)
		return convertSystem_Cpu_Nice_AvgPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/nice/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Nice_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Nice_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/nice/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Nice_AvgPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/nice/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/nice/avg to the batch object.
func (n *System_Cpu_Nice_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/nice/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Nice_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/nice/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Nice_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Nice_AvgPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/nice/avg to the batch object.
func (n *System_Cpu_Nice_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Nice_AvgPath extracts the value of the leaf Avg from its parent oc.System_Cpu_Nice
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Nice_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Nice) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/nice/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Nice_InstantPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Nice{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Nice", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Nice_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/nice/instant with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Nice_InstantPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/nice/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Nice_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Nice{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Nice", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Nice_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/nice/instant with a ONCE subscription.
func (n *System_Cpu_Nice_InstantPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/nice/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Nice_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Nice_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Nice{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Nice", gs, queryPath, true, false)
		return convertSystem_Cpu_Nice_InstantPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/nice/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Nice_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Nice_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/nice/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Nice_InstantPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/nice/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/nice/instant to the batch object.
func (n *System_Cpu_Nice_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/nice/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Nice_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/nice/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Nice_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Nice_InstantPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/nice/instant to the batch object.
func (n *System_Cpu_Nice_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Nice_InstantPath extracts the value of the leaf Instant from its parent oc.System_Cpu_Nice
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Nice_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Nice) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/nice/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Nice_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Cpu_Nice{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Nice", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Nice_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/nice/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Nice_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/nice/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Nice_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Nice{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Nice", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Nice_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/nice/interval with a ONCE subscription.
func (n *System_Cpu_Nice_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/nice/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Nice_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Nice_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Cpu_Nice{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Nice", gs, queryPath, true, false)
		return convertSystem_Cpu_Nice_IntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/nice/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Nice_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Nice_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/nice/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Nice_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/nice/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/nice/interval to the batch object.
func (n *System_Cpu_Nice_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/nice/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Nice_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/nice/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Nice_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Cpu_Nice_IntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/nice/interval to the batch object.
func (n *System_Cpu_Nice_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Nice_IntervalPath extracts the value of the leaf Interval from its parent oc.System_Cpu_Nice
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Cpu_Nice_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Nice) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-system/system/cpus/cpu/state/nice/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Cpu_Nice_MaxPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Cpu_Nice{}
	md, ok := oc.Lookup(t, n, "System_Cpu_Nice", goStruct, true, false)
	if ok {
		return convertSystem_Cpu_Nice_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/cpus/cpu/state/nice/max with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Cpu_Nice_MaxPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/cpus/cpu/state/nice/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Cpu_Nice_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Cpu_Nice{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Cpu_Nice", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Cpu_Nice_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/cpus/cpu/state/nice/max with a ONCE subscription.
func (n *System_Cpu_Nice_MaxPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/nice/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Nice_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Cpu_Nice_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Cpu_Nice{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Cpu_Nice", gs, queryPath, true, false)
		return convertSystem_Cpu_Nice_MaxPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/nice/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Nice_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Nice_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/cpus/cpu/state/nice/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Cpu_Nice_MaxPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/cpus/cpu/state/nice/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/cpus/cpu/state/nice/max to the batch object.
func (n *System_Cpu_Nice_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/cpus/cpu/state/nice/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Cpu_Nice_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/cpus/cpu/state/nice/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Cpu_Nice_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Cpu_Nice_MaxPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/cpus/cpu/state/nice/max to the batch object.
func (n *System_Cpu_Nice_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Cpu_Nice_MaxPath extracts the value of the leaf Max from its parent oc.System_Cpu_Nice
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Cpu_Nice_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Cpu_Nice) *oc.QualifiedUint8 {
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
