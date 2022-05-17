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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/exceeding-pkts to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath extracts the value of the leaf ExceedingPkts from its parent oc.Qos_Interface_Input_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_SchedulerPolicy_Scheduler_ExceedingPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_SchedulerPolicy_Scheduler) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ExceedingPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/sequence to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_SequencePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath extracts the value of the leaf Sequence from its parent oc.Qos_Interface_Input_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_Interface_Input_SchedulerPolicy_Scheduler_SequencePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_SchedulerPolicy_Scheduler) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-octets to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath extracts the value of the leaf ViolatingOctets from its parent oc.Qos_Interface_Input_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_SchedulerPolicy_Scheduler) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ViolatingOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts with a ONCE subscription.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_SchedulerPolicy_Scheduler{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_SchedulerPolicy_Scheduler", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/scheduler-policy/schedulers/scheduler/state/violating-pkts to the batch object.
func (n *Qos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath extracts the value of the leaf ViolatingPkts from its parent oc.Qos_Interface_Input_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_SchedulerPolicy_Scheduler_ViolatingPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_SchedulerPolicy_Scheduler) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ViolatingPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_UnicastBufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_UnicastBufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_UnicastBufferAllocationProfilePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_UnicastBufferAllocationProfilePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_UnicastBufferAllocationProfilePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_UnicastBufferAllocationProfilePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_Interface_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_UnicastBufferAllocationProfilePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_UnicastBufferAllocationProfilePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/state/unicast-buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Input_UnicastBufferAllocationProfilePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_UnicastBufferAllocationProfilePath extracts the value of the leaf UnicastBufferAllocationProfile from its parent oc.Qos_Interface_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_UnicastBufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.UnicastBufferAllocationProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterfacePath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_VoqInterface {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_VoqInterface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterfacePath) Get(t testing.TB) *oc.Qos_Interface_Input_VoqInterface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_VoqInterface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_VoqInterface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_VoqInterface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterfacePathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_VoqInterface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_VoqInterface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_VoqInterface {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_VoqInterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_VoqInterface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Input_VoqInterface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Input_VoqInterface)))
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_VoqInterface) bool) *oc.Qos_Interface_Input_VoqInterfaceWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Input_VoqInterfaceWatcher{}
	gs := &oc.Qos_Interface_Input_VoqInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_VoqInterface", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Interface_Input_VoqInterface{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input_VoqInterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_VoqInterface) bool) *oc.Qos_Interface_Input_VoqInterfaceWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_VoqInterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Input_VoqInterface) *oc.QualifiedQos_Interface_Input_VoqInterface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Input_VoqInterface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface to the batch object.
func (n *Qos_Interface_Input_VoqInterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_VoqInterface {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_VoqInterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_VoqInterface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_VoqInterface) bool) *oc.Qos_Interface_Input_VoqInterfaceWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Input_VoqInterfaceWatcher{}
	structs := map[string]*oc.Qos_Interface_Input_VoqInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_VoqInterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_VoqInterface", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Interface_Input_VoqInterface{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input_VoqInterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_VoqInterface) bool) *oc.Qos_Interface_Input_VoqInterfaceWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface to the batch object.
func (n *Qos_Interface_Input_VoqInterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_VoqInterface_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_VoqInterface_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Input_VoqInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_VoqInterface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_VoqInterface_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name to the batch object.
func (n *Qos_Interface_Input_VoqInterface_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_Interface_Input_VoqInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_VoqInterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_VoqInterface", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_VoqInterface_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/state/name to the batch object.
func (n *Qos_Interface_Input_VoqInterface_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_VoqInterface_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Input_VoqInterface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_VoqInterface_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_VoqInterface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface_Queue", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_VoqInterface_Queue{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Get(t testing.TB) *oc.Qos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_VoqInterface_Queue
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_VoqInterface_Queue{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_QueuePathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_VoqInterface_Queue
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_VoqInterface_Queue{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_VoqInterface_Queue) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Input_VoqInterface_Queue{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Input_VoqInterface_Queue)))
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_QueuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_VoqInterface_Queue) bool) *oc.Qos_Interface_Input_VoqInterface_QueueWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Input_VoqInterface_QueueWatcher{}
	gs := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Interface_Input_VoqInterface_Queue{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input_VoqInterface_Queue)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_VoqInterface_Queue) bool) *oc.Qos_Interface_Input_VoqInterface_QueueWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_QueuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Input_VoqInterface_Queue) *oc.QualifiedQos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Input_VoqInterface_Queue) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue to the batch object.
func (n *Qos_Interface_Input_VoqInterface_QueuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_QueuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_VoqInterface_Queue{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_VoqInterface_Queue) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_QueuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_VoqInterface_Queue) bool) *oc.Qos_Interface_Input_VoqInterface_QueueWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Input_VoqInterface_QueueWatcher{}
	structs := map[string]*oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_VoqInterface_Queue{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Interface_Input_VoqInterface_Queue{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input_VoqInterface_Queue)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_QueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_VoqInterface_Queue) bool) *oc.Qos_Interface_Input_VoqInterface_QueueWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_QueuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue to the batch object.
func (n *Qos_Interface_Input_VoqInterface_QueuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_VoqInterface_Queue{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/avg-queue-len to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_AvgQueueLenPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath extracts the value of the leaf AvgQueueLen from its parent oc.Qos_Interface_Input_VoqInterface_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_VoqInterface_Queue_AvgQueueLenPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_VoqInterface_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AvgQueueLen
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_VoqInterface_Queue_DroppedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_VoqInterface_Queue_DroppedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_VoqInterface_Queue_DroppedPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_VoqInterface_Queue{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_VoqInterface_Queue_DroppedPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/dropped-pkts to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_DroppedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_VoqInterface_Queue_DroppedPktsPath extracts the value of the leaf DroppedPkts from its parent oc.Qos_Interface_Input_VoqInterface_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_VoqInterface_Queue_DroppedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_VoqInterface_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.DroppedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_VoqInterface_Queue{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/max-queue-len to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_MaxQueueLenPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath extracts the value of the leaf MaxQueueLen from its parent oc.Qos_Interface_Input_VoqInterface_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_VoqInterface_Queue_MaxQueueLenPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_VoqInterface_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxQueueLen
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_VoqInterface_Queue_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_VoqInterface_Queue_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_VoqInterface_Queue_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_VoqInterface_Queue{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_VoqInterface_Queue_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/name to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_VoqInterface_Queue_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Input_VoqInterface_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_VoqInterface_Queue_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_VoqInterface_Queue) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_VoqInterface_Queue{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-octets to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath extracts the value of the leaf TransmitOctets from its parent oc.Qos_Interface_Input_VoqInterface_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_VoqInterface_Queue_TransmitOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_VoqInterface_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TransmitOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_VoqInterface_Queue", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_VoqInterface_Queue_TransmitPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_VoqInterface_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_VoqInterface_Queue_TransmitPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts with a ONCE subscription.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_VoqInterface_Queue_TransmitPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Input_VoqInterface_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_VoqInterface_Queue{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_VoqInterface_Queue", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_VoqInterface_Queue_TransmitPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/virtual-output-queues/voq-interface/queues/queue/state/transmit-pkts to the batch object.
func (n *Qos_Interface_Input_VoqInterface_Queue_TransmitPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_VoqInterface_Queue_TransmitPktsPath extracts the value of the leaf TransmitPkts from its parent oc.Qos_Interface_Input_VoqInterface_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Input_VoqInterface_Queue_TransmitPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_VoqInterface_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TransmitPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/state/interface-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_InterfaceIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface{}
	md, ok := oc.Lookup(t, n, "Qos_Interface", goStruct, true, false)
	if ok {
		return convertQos_Interface_InterfaceIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/state/interface-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_InterfaceIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/state/interface-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_InterfaceIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_InterfaceIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/state/interface-id with a ONCE subscription.
func (n *Qos_Interface_InterfaceIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/state/interface-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_InterfaceIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_InterfaceIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_InterfaceIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/state/interface-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_InterfaceIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_InterfaceIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/state/interface-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_InterfaceIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/state/interface-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/state/interface-id to the batch object.
func (n *Qos_Interface_InterfaceIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/state/interface-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_InterfaceIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_InterfaceIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_InterfaceIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/state/interface-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_InterfaceIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_InterfaceIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/state/interface-id to the batch object.
func (n *Qos_Interface_InterfaceIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_InterfaceIdPath extracts the value of the leaf InterfaceId from its parent oc.Qos_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_InterfaceIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.InterfaceId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_InterfaceRef {
	t.Helper()
	goStruct := &oc.Qos_Interface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_InterfaceRef", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_InterfaceRefPath) Get(t testing.TB) *oc.Qos_Interface_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_InterfaceRef", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref with a ONCE subscription.
func (n *Qos_Interface_InterfaceRefPathAny) Get(t testing.TB) []*oc.Qos_Interface_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_InterfaceRefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_InterfaceRef {
	t.Helper()
	c := &oc.CollectionQos_Interface_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_InterfaceRef) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_InterfaceRef{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_InterfaceRef)))
		return false
	})
	return c
}

func watch_Qos_Interface_InterfaceRefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_InterfaceRef) bool) *oc.Qos_Interface_InterfaceRefWatcher {
	t.Helper()
	w := &oc.Qos_Interface_InterfaceRefWatcher{}
	gs := &oc.Qos_Interface_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_InterfaceRef", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Interface_InterfaceRef{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_InterfaceRefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_InterfaceRef) bool) *oc.Qos_Interface_InterfaceRefWatcher {
	t.Helper()
	return watch_Qos_Interface_InterfaceRefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/interface-ref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_InterfaceRefPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_InterfaceRef) *oc.QualifiedQos_Interface_InterfaceRef {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_InterfaceRef) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/interface-ref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/interface-ref to the batch object.
func (n *Qos_Interface_InterfaceRefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_InterfaceRefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_InterfaceRef {
	t.Helper()
	c := &oc.CollectionQos_Interface_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_InterfaceRef) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_InterfaceRefPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_InterfaceRef) bool) *oc.Qos_Interface_InterfaceRefWatcher {
	t.Helper()
	w := &oc.Qos_Interface_InterfaceRefWatcher{}
	structs := map[string]*oc.Qos_Interface_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_InterfaceRef", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Interface_InterfaceRef{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_InterfaceRefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_InterfaceRef) bool) *oc.Qos_Interface_InterfaceRefWatcher {
	t.Helper()
	return watch_Qos_Interface_InterfaceRefPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/interface-ref to the batch object.
func (n *Qos_Interface_InterfaceRefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_InterfaceRef", goStruct, true, false)
	if ok {
		return convertQos_Interface_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface with a ONCE subscription.
func (n *Qos_Interface_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_InterfaceRef_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_InterfaceRef", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_InterfaceRef_InterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface to the batch object.
func (n *Qos_Interface_InterfaceRef_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_InterfaceRef_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_InterfaceRef_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_Interface_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_InterfaceRef", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_InterfaceRef_InterfacePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_InterfaceRef_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_InterfaceRef_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/interface-ref/state/interface to the batch object.
func (n *Qos_Interface_InterfaceRef_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Qos_Interface_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_InterfaceRef) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Interface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_Interface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_InterfaceRef", goStruct, true, false)
	if ok {
		return convertQos_Interface_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface with a ONCE subscription.
func (n *Qos_Interface_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_InterfaceRef_SubinterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_Interface_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_InterfaceRef", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_InterfaceRef_SubinterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_Interface_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface to the batch object.
func (n *Qos_Interface_InterfaceRef_SubinterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_InterfaceRef_SubinterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_InterfaceRef_SubinterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Qos_Interface_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_InterfaceRef", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_InterfaceRef_SubinterfacePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_InterfaceRef_SubinterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_Interface_InterfaceRef_SubinterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/interface-ref/state/subinterface to the batch object.
func (n *Qos_Interface_InterfaceRef_SubinterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Qos_Interface_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_Interface_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_InterfaceRef) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Subinterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_OutputPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Output {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Output{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_OutputPath) Get(t testing.TB) *oc.Qos_Interface_Output {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_OutputPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Output {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Output
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Output{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output with a ONCE subscription.
func (n *Qos_Interface_OutputPathAny) Get(t testing.TB) []*oc.Qos_Interface_Output {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Output
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_OutputPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Output {
	t.Helper()
	c := &oc.CollectionQos_Interface_Output{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Output) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Output{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Output)))
		return false
	})
	return c
}

func watch_Qos_Interface_OutputPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output) bool) *oc.Qos_Interface_OutputWatcher {
	t.Helper()
	w := &oc.Qos_Interface_OutputWatcher{}
	gs := &oc.Qos_Interface_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Interface_Output{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Output)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_OutputPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output) bool) *oc.Qos_Interface_OutputWatcher {
	t.Helper()
	return watch_Qos_Interface_OutputPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_OutputPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Output) *oc.QualifiedQos_Interface_Output {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Output) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output to the batch object.
func (n *Qos_Interface_OutputPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_OutputPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Output {
	t.Helper()
	c := &oc.CollectionQos_Interface_Output{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Output) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_OutputPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output) bool) *oc.Qos_Interface_OutputWatcher {
	t.Helper()
	w := &oc.Qos_Interface_OutputWatcher{}
	structs := map[string]*oc.Qos_Interface_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Output{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Output", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Interface_Output{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Output)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_OutputPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output) bool) *oc.Qos_Interface_OutputWatcher {
	t.Helper()
	return watch_Qos_Interface_OutputPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output to the batch object.
func (n *Qos_Interface_OutputPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_BufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_BufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_BufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Output_BufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_BufferAllocationProfilePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Output_BufferAllocationProfilePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_BufferAllocationProfilePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Output_BufferAllocationProfilePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_BufferAllocationProfilePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_BufferAllocationProfilePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_Interface_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Output{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Output", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Output_BufferAllocationProfilePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_BufferAllocationProfilePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_BufferAllocationProfilePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/state/buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Output_BufferAllocationProfilePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_BufferAllocationProfilePath extracts the value of the leaf BufferAllocationProfile from its parent oc.Qos_Interface_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_BufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.BufferAllocationProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_ClassifierPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Output_Classifier {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_Classifier", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Output_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_ClassifierPath) Get(t testing.TB) *oc.Qos_Interface_Output_Classifier {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_ClassifierPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Output_Classifier {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Output_Classifier
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_Classifier", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Output_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a ONCE subscription.
func (n *Qos_Interface_Output_ClassifierPathAny) Get(t testing.TB) []*oc.Qos_Interface_Output_Classifier {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Output_Classifier
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_ClassifierPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Output_Classifier {
	t.Helper()
	c := &oc.CollectionQos_Interface_Output_Classifier{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Output_Classifier) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Output_Classifier{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Output_Classifier)))
		return false
	})
	return c
}

func watch_Qos_Interface_Output_ClassifierPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output_Classifier) bool) *oc.Qos_Interface_Output_ClassifierWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Output_ClassifierWatcher{}
	gs := &oc.Qos_Interface_Output_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_Classifier", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Interface_Output_Classifier{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Output_Classifier)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_ClassifierPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output_Classifier) bool) *oc.Qos_Interface_Output_ClassifierWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_ClassifierPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_ClassifierPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Output_Classifier) *oc.QualifiedQos_Interface_Output_Classifier {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Output_Classifier) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier to the batch object.
func (n *Qos_Interface_Output_ClassifierPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_ClassifierPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Output_Classifier {
	t.Helper()
	c := &oc.CollectionQos_Interface_Output_Classifier{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Output_Classifier) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_ClassifierPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output_Classifier) bool) *oc.Qos_Interface_Output_ClassifierWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Output_ClassifierWatcher{}
	structs := map[string]*oc.Qos_Interface_Output_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Output_Classifier{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Output_Classifier", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Interface_Output_Classifier{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Output_Classifier)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_ClassifierPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Output_Classifier) bool) *oc.Qos_Interface_Output_ClassifierWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_ClassifierPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/classifiers/classifier to the batch object.
func (n *Qos_Interface_Output_ClassifierPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
