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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-mpls-tc to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction)))
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/drop to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath extracts the value of the leaf Drop from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_DropPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dot1p to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-dscp to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/exceed-action/state/set-mpls-tc to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-bytes to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath extracts the value of the leaf MaxQueueDepthBytes from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthBytesPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.MaxQueueDepthBytes
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-packets to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath extracts the value of the leaf MaxQueueDepthPackets from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPacketsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.MaxQueueDepthPackets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/max-queue-depth-percent to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath extracts the value of the leaf MaxQueueDepthPercent from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_MaxQueueDepthPercentPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.MaxQueueDepthPercent
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Lookup(t testing.TB) *oc.QualifiedE_QosTypes_QueueBehavior {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Get(t testing.TB) oc.E_QosTypes_QueueBehavior {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPathAny) Lookup(t testing.TB) []*oc.QualifiedE_QosTypes_QueueBehavior {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_QosTypes_QueueBehavior
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPathAny) Get(t testing.TB) []oc.E_QosTypes_QueueBehavior {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_QosTypes_QueueBehavior
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_QosTypes_QueueBehavior {
	t.Helper()
	c := &oc.CollectionE_QosTypes_QueueBehavior{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_QosTypes_QueueBehavior) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_QosTypes_QueueBehavior) bool) *oc.E_QosTypes_QueueBehaviorWatcher {
	t.Helper()
	w := &oc.E_QosTypes_QueueBehaviorWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_QosTypes_QueueBehavior)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_QosTypes_QueueBehavior) bool) *oc.E_QosTypes_QueueBehaviorWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Await(t testing.TB, timeout time.Duration, val oc.E_QosTypes_QueueBehavior) *oc.QualifiedE_QosTypes_QueueBehavior {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_QosTypes_QueueBehavior) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_QosTypes_QueueBehavior {
	t.Helper()
	c := &oc.CollectionE_QosTypes_QueueBehavior{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_QosTypes_QueueBehavior) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_QosTypes_QueueBehavior) bool) *oc.E_QosTypes_QueueBehaviorWatcher {
	t.Helper()
	w := &oc.E_QosTypes_QueueBehaviorWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_QosTypes_QueueBehavior)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_QosTypes_QueueBehavior) bool) *oc.E_QosTypes_QueueBehaviorWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/queuing-behavior to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath extracts the value of the leaf QueuingBehavior from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_QosTypes_QueueBehavior.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_QueuingBehaviorPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedE_QosTypes_QueueBehavior {
	t.Helper()
	qv := &oc.QualifiedE_QosTypes_QueueBehavior{
		Metadata: md,
	}
	val := parent.QueuingBehavior
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Output", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Output{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_Output
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Output{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_Output
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_Output{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_Output) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Output{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_SchedulerPolicy_Scheduler_Output)))
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OutputPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_Output) bool) *oc.Qos_SchedulerPolicy_Scheduler_OutputWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_OutputWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Output{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_Output)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_Output) bool) *oc.Qos_SchedulerPolicy_Scheduler_OutputWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OutputPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_SchedulerPolicy_Scheduler_Output) *oc.QualifiedQos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_SchedulerPolicy_Scheduler_Output) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_Output{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_Output) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OutputPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_Output) bool) *oc.Qos_SchedulerPolicy_Scheduler_OutputWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_OutputWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_Output{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Output{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_Output)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_Output) bool) *oc.Qos_SchedulerPolicy_Scheduler_OutputWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OutputPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OutputPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Output", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_Output{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/child-scheduler to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath extracts the value of the leaf ChildScheduler from its parent oc.Qos_SchedulerPolicy_Scheduler_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_SchedulerPolicy_Scheduler_Output_ChildSchedulerPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Output) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ChildScheduler
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Output", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_Output{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-fwd-group to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath extracts the value of the leaf OutputFwdGroup from its parent oc.Qos_SchedulerPolicy_Scheduler_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_SchedulerPolicy_Scheduler_Output_OutputFwdGroupPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Output) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.OutputFwdGroup
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Lookup(t testing.TB) *oc.QualifiedE_Output_OutputType {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Output", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Output_OutputTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Get(t testing.TB) oc.E_Output_OutputType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Output_OutputType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Output_OutputType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Output_OutputTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePathAny) Get(t testing.TB) []oc.E_Output_OutputType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Output_OutputType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Output_OutputType {
	t.Helper()
	c := &oc.CollectionE_Output_OutputType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Output_OutputType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Output_OutputType) bool) *oc.E_Output_OutputTypeWatcher {
	t.Helper()
	w := &oc.E_Output_OutputTypeWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_Output_OutputTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Output_OutputType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Output_OutputType) bool) *oc.E_Output_OutputTypeWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Output_OutputType) *oc.QualifiedE_Output_OutputType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Output_OutputType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Output_OutputType {
	t.Helper()
	c := &oc.CollectionE_Output_OutputType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Output_OutputType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_Output_OutputTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Output_OutputType) bool) *oc.E_Output_OutputTypeWatcher {
	t.Helper()
	w := &oc.E_Output_OutputTypeWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_Output{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Output", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_Output_OutputTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Output_OutputType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Output_OutputType) bool) *oc.E_Output_OutputTypeWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Output_OutputTypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/output/state/output-type to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Output_OutputTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_Output_OutputTypePath extracts the value of the leaf OutputType from its parent oc.Qos_SchedulerPolicy_Scheduler_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Output_OutputType.
func convertQos_SchedulerPolicy_Scheduler_Output_OutputTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Output) *oc.QualifiedE_Output_OutputType {
	t.Helper()
	qv := &oc.QualifiedE_Output_OutputType{
		Metadata: md,
	}
	val := parent.OutputType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Lookup(t testing.TB) *oc.QualifiedE_Scheduler_Priority {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_PriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Get(t testing.TB) oc.E_Scheduler_Priority {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Scheduler_Priority {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Scheduler_Priority
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_PriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPathAny) Get(t testing.TB) []oc.E_Scheduler_Priority {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Scheduler_Priority
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Scheduler_Priority {
	t.Helper()
	c := &oc.CollectionE_Scheduler_Priority{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Scheduler_Priority) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_PriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Scheduler_Priority) bool) *oc.E_Scheduler_PriorityWatcher {
	t.Helper()
	w := &oc.E_Scheduler_PriorityWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_PriorityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Scheduler_Priority)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Scheduler_Priority) bool) *oc.E_Scheduler_PriorityWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_PriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Await(t testing.TB, timeout time.Duration, val oc.E_Scheduler_Priority) *oc.QualifiedE_Scheduler_Priority {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Scheduler_Priority) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Scheduler_Priority {
	t.Helper()
	c := &oc.CollectionE_Scheduler_Priority{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Scheduler_Priority) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_PriorityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Scheduler_Priority) bool) *oc.E_Scheduler_PriorityWatcher {
	t.Helper()
	w := &oc.E_Scheduler_PriorityWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_PriorityPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Scheduler_Priority)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Scheduler_Priority) bool) *oc.E_Scheduler_PriorityWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_PriorityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/priority to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_PriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_PriorityPath extracts the value of the leaf Priority from its parent oc.Qos_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Scheduler_Priority.
func convertQos_SchedulerPolicy_Scheduler_PriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler) *oc.QualifiedE_Scheduler_Priority {
	t.Helper()
	qv := &oc.QualifiedE_Scheduler_Priority{
		Metadata: md,
	}
	val := parent.Priority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_SequencePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_SequencePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_SequencePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_SequencePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_SequencePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_SequencePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_SequencePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_SequencePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/state/sequence to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_SequencePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_SequencePath extracts the value of the leaf Sequence from its parent oc.Qos_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_SequencePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Sequence
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor)))
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) bool) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) bool) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) bool) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) bool) *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/bc to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath extracts the value of the leaf Bc from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Bc
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/be to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath extracts the value of the leaf Be from its parent oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_BePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Be
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t, md, gs)}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t, md, gs)}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/two-rate-three-color/state/cir-pct with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_CirPctPathAny(t, n, timeout, predicate)
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
