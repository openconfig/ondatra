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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_SchedulerPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_SchedulerPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_SchedulerPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a ONCE subscription.
func (n *Qos_SchedulerPolicy_SchedulerPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_SchedulerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_SchedulerPolicy_Scheduler{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_SchedulerPolicy_Scheduler)))
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_SchedulerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler) bool) *oc.Qos_SchedulerPolicy_SchedulerWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_SchedulerWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", gs, queryPath, false, false)
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_SchedulerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler) bool) *oc.Qos_SchedulerPolicy_SchedulerWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_SchedulerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_SchedulerPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_SchedulerPolicy_Scheduler) *oc.QualifiedQos_SchedulerPolicy_Scheduler {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_SchedulerPolicy_Scheduler) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler to the batch object.
func (n *Qos_SchedulerPolicy_SchedulerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_SchedulerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_SchedulerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler) bool) *oc.Qos_SchedulerPolicy_SchedulerWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_SchedulerPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler to the batch object.
func (n *Qos_SchedulerPolicy_SchedulerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Input", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Input{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_InputPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_Input
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Input{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_InputPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_Input
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_Input{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Input{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_SchedulerPolicy_Scheduler_Input)))
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_InputPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input) bool) *oc.Qos_SchedulerPolicy_Scheduler_InputWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_InputWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", gs, queryPath, false, false)
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Input{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_Input)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input) bool) *oc.Qos_SchedulerPolicy_Scheduler_InputWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_InputPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_SchedulerPolicy_Scheduler_Input) *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_InputPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_InputPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_Input{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_InputPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input) bool) *oc.Qos_SchedulerPolicy_Scheduler_InputWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_InputPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_InputPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Input", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Input_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Input_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_Input_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_Input_IdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_IdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_Input_IdPath extracts the value of the leaf Id from its parent oc.Qos_SchedulerPolicy_Scheduler_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_SchedulerPolicy_Scheduler_Input_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Input) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Id
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Lookup(t testing.TB) *oc.QualifiedE_Input_InputType {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Input", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Get(t testing.TB) oc.E_Input_InputType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Input_InputType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Input_InputType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePathAny) Get(t testing.TB) []oc.E_Input_InputType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Input_InputType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Input_InputType {
	t.Helper()
	c := &oc.CollectionE_Input_InputType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Input_InputType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_Input_InputTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Input_InputType) bool) *oc.E_Input_InputTypeWatcher {
	t.Helper()
	w := &oc.E_Input_InputTypeWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Input_InputType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Input_InputType) bool) *oc.E_Input_InputTypeWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_InputTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Input_InputType) *oc.QualifiedE_Input_InputType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Input_InputType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Input_InputType {
	t.Helper()
	c := &oc.CollectionE_Input_InputType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Input_InputType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Input_InputType) bool) *oc.E_Input_InputTypeWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_InputTypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/input-type to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath extracts the value of the leaf InputType from its parent oc.Qos_SchedulerPolicy_Scheduler_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Input_InputType.
func convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Input) *oc.QualifiedE_Input_InputType {
	t.Helper()
	qv := &oc.QualifiedE_Input_InputType{
		Metadata: md,
	}
	val := parent.InputType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Input", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Input_QueuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Input_QueuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_Input_QueuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_Input_QueuePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_QueuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_QueuePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_Input_QueuePath extracts the value of the leaf Queue from its parent oc.Qos_SchedulerPolicy_Scheduler_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_SchedulerPolicy_Scheduler_Input_QueuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Input) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Queue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_Input", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_Input_WeightPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_Input_WeightPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_Input_WeightPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_Input_WeightPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_WeightPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_WeightPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_Input_WeightPath extracts the value of the leaf Weight from its parent oc.Qos_SchedulerPolicy_Scheduler_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_SchedulerPolicy_Scheduler_Input_WeightPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_Input) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Weight
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor)))
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, false, false)
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
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
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath extracts the value of the leaf Bc from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath extracts the value of the leaf Cir from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
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
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath extracts the value of the leaf CirPct from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
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
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath extracts the value of the leaf CirPctRemaining from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *oc.QualifiedUint8 {
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
