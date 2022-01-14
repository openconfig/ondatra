package qos

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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath extracts the value of the leaf Cir from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Cir
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath extracts the value of the leaf CirPct from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.CirPct
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct-remaining to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath extracts the value of the leaf CirPctRemaining from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctRemainingPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.CirPctRemaining
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction)))
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", gs, queryPath, false, false)
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dot1p to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetDot1P
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-dscp to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetDscp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/conform-action/state/set-mpls-tc to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetMplsTc
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction)))
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", gs, queryPath, false, false)
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/drop to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath extracts the value of the leaf Drop from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_DropPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Drop
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dot1p to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetDot1P
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/exceed-action/state/set-dscp to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.SetDscp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
