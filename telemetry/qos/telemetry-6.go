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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a ONCE subscription.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts to the batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/exceeding-pkts to the batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath extracts the value of the leaf ExceedingPkts from its parent oc.Qos_Interface_Output_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Output_SchedulerPolicy_Scheduler_ExceedingPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_SchedulerPolicy_Scheduler) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence with a ONCE subscription.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence to the batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/sequence to the batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_SequencePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath extracts the value of the leaf Sequence from its parent oc.Qos_Interface_Output_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_Interface_Output_SchedulerPolicy_Scheduler_SequencePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_SchedulerPolicy_Scheduler) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets with a ONCE subscription.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets to the batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-octets to the batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath extracts the value of the leaf ViolatingOctets from its parent oc.Qos_Interface_Output_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_SchedulerPolicy_Scheduler) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output_SchedulerPolicy_Scheduler", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts with a ONCE subscription.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts to the batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Output_SchedulerPolicy_Scheduler{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Output_SchedulerPolicy_Scheduler", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/scheduler-policy/schedulers/scheduler/state/violating-pkts to the batch object.
func (n *Qos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath extracts the value of the leaf ViolatingPkts from its parent oc.Qos_Interface_Output_SchedulerPolicy_Scheduler
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_Interface_Output_SchedulerPolicy_Scheduler_ViolatingPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output_SchedulerPolicy_Scheduler) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Output{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Output", goStruct, true, false)
	if ok {
		return convertQos_Interface_Output_UnicastBufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertQos_Interface_Output_UnicastBufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_UnicastBufferAllocationProfilePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Output{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Output", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Output_UnicastBufferAllocationProfilePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_UnicastBufferAllocationProfilePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Output_UnicastBufferAllocationProfilePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertQos_Interface_Output_UnicastBufferAllocationProfilePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Output_UnicastBufferAllocationProfilePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/output/state/unicast-buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Output_UnicastBufferAllocationProfilePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Output_UnicastBufferAllocationProfilePath extracts the value of the leaf UnicastBufferAllocationProfile from its parent oc.Qos_Interface_Output
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Output_UnicastBufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Output) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfilePath) Lookup(t testing.TB) *oc.QualifiedQos_QueueManagementProfile {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_QueueManagementProfile{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfilePath) Get(t testing.TB) *oc.Qos_QueueManagementProfile {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_QueueManagementProfile {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_QueueManagementProfile
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_QueueManagementProfile{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a ONCE subscription.
func (n *Qos_QueueManagementProfilePathAny) Get(t testing.TB) []*oc.Qos_QueueManagementProfile {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_QueueManagementProfile
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfilePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_QueueManagementProfile {
	t.Helper()
	c := &oc.CollectionQos_QueueManagementProfile{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_QueueManagementProfile) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_QueueManagementProfile{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_QueueManagementProfile)))
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfilePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile) bool) *oc.Qos_QueueManagementProfileWatcher {
	t.Helper()
	w := &oc.Qos_QueueManagementProfileWatcher{}
	gs := &oc.Qos_QueueManagementProfile{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_QueueManagementProfile{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_QueueManagementProfile)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfilePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile) bool) *oc.Qos_QueueManagementProfileWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfilePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfilePath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_QueueManagementProfile) *oc.QualifiedQos_QueueManagementProfile {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_QueueManagementProfile) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile to the batch object.
func (n *Qos_QueueManagementProfilePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfilePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_QueueManagementProfile {
	t.Helper()
	c := &oc.CollectionQos_QueueManagementProfile{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_QueueManagementProfile) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfilePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile) bool) *oc.Qos_QueueManagementProfileWatcher {
	t.Helper()
	w := &oc.Qos_QueueManagementProfileWatcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_QueueManagementProfile{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_QueueManagementProfile)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfilePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile) bool) *oc.Qos_QueueManagementProfileWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfilePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile to the batch object.
func (n *Qos_QueueManagementProfilePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name with a ONCE subscription.
func (n *Qos_QueueManagementProfile_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_QueueManagementProfile{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name to the batch object.
func (n *Qos_QueueManagementProfile_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/state/name to the batch object.
func (n *Qos_QueueManagementProfile_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_NamePath extracts the value of the leaf Name from its parent oc.Qos_QueueManagementProfile
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_QueueManagementProfile_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_RedPath) Lookup(t testing.TB) *oc.QualifiedQos_QueueManagementProfile_Red {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_QueueManagementProfile_Red{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_RedPath) Get(t testing.TB) *oc.Qos_QueueManagementProfile_Red {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_RedPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_QueueManagementProfile_Red {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_QueueManagementProfile_Red
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_QueueManagementProfile_Red{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a ONCE subscription.
func (n *Qos_QueueManagementProfile_RedPathAny) Get(t testing.TB) []*oc.Qos_QueueManagementProfile_Red {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_QueueManagementProfile_Red
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_RedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_QueueManagementProfile_Red {
	t.Helper()
	c := &oc.CollectionQos_QueueManagementProfile_Red{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_QueueManagementProfile_Red) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_QueueManagementProfile_Red{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_QueueManagementProfile_Red)))
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_RedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Red) bool) *oc.Qos_QueueManagementProfile_RedWatcher {
	t.Helper()
	w := &oc.Qos_QueueManagementProfile_RedWatcher{}
	gs := &oc.Qos_QueueManagementProfile_Red{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Red", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_QueueManagementProfile_Red{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_QueueManagementProfile_Red)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_RedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Red) bool) *oc.Qos_QueueManagementProfile_RedWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_RedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_RedPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_QueueManagementProfile_Red) *oc.QualifiedQos_QueueManagementProfile_Red {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_QueueManagementProfile_Red) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red to the batch object.
func (n *Qos_QueueManagementProfile_RedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_RedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_QueueManagementProfile_Red {
	t.Helper()
	c := &oc.CollectionQos_QueueManagementProfile_Red{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_QueueManagementProfile_Red) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_RedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Red) bool) *oc.Qos_QueueManagementProfile_RedWatcher {
	t.Helper()
	w := &oc.Qos_QueueManagementProfile_RedWatcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Red{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Red{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Red", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_QueueManagementProfile_Red{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_QueueManagementProfile_Red)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_RedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Red) bool) *oc.Qos_QueueManagementProfile_RedWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_RedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red to the batch object.
func (n *Qos_QueueManagementProfile_RedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Lookup(t testing.TB) *oc.QualifiedQos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red_Uniform", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_QueueManagementProfile_Red_Uniform{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Get(t testing.TB) *oc.Qos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Red_UniformPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_QueueManagementProfile_Red_Uniform
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_QueueManagementProfile_Red_Uniform{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Red_UniformPathAny) Get(t testing.TB) []*oc.Qos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_QueueManagementProfile_Red_Uniform
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	c := &oc.CollectionQos_QueueManagementProfile_Red_Uniform{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_QueueManagementProfile_Red_Uniform) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_QueueManagementProfile_Red_Uniform{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_QueueManagementProfile_Red_Uniform)))
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Red_UniformPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Red_Uniform) bool) *oc.Qos_QueueManagementProfile_Red_UniformWatcher {
	t.Helper()
	w := &oc.Qos_QueueManagementProfile_Red_UniformWatcher{}
	gs := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_QueueManagementProfile_Red_Uniform{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_QueueManagementProfile_Red_Uniform)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Red_Uniform) bool) *oc.Qos_QueueManagementProfile_Red_UniformWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Red_UniformPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_QueueManagementProfile_Red_Uniform) *oc.QualifiedQos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_QueueManagementProfile_Red_Uniform) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform to the batch object.
func (n *Qos_QueueManagementProfile_Red_UniformPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Red_UniformPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	c := &oc.CollectionQos_QueueManagementProfile_Red_Uniform{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_QueueManagementProfile_Red_Uniform) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Red_UniformPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Red_Uniform) bool) *oc.Qos_QueueManagementProfile_Red_UniformWatcher {
	t.Helper()
	w := &oc.Qos_QueueManagementProfile_Red_UniformWatcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Red_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Red_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_QueueManagementProfile_Red_Uniform{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_QueueManagementProfile_Red_Uniform)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Red_UniformPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Red_Uniform) bool) *oc.Qos_QueueManagementProfile_Red_UniformWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Red_UniformPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform to the batch object.
func (n *Qos_QueueManagementProfile_Red_UniformPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red_Uniform", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_Red_Uniform_DropPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDrop())
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Red_Uniform_DropPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Red_Uniform_DropPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_Red_Uniform_DropPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Red_Uniform_DropPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop to the batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Red_Uniform_DropPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Red_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Red_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_Red_Uniform_DropPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Red_Uniform_DropPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/drop to the batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_DropPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_Red_Uniform_DropPath extracts the value of the leaf Drop from its parent oc.Qos_QueueManagementProfile_Red_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_QueueManagementProfile_Red_Uniform_DropPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Red_Uniform) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red_Uniform", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_Red_Uniform_EnableEcnPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnableEcn())
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Red_Uniform_EnableEcnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_Red_Uniform_EnableEcnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn to the batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Red_Uniform_EnableEcnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Red_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Red_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_Red_Uniform_EnableEcnPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Red_Uniform_EnableEcnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/enable-ecn to the batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_EnableEcnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_Red_Uniform_EnableEcnPath extracts the value of the leaf EnableEcn from its parent oc.Qos_QueueManagementProfile_Red_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_QueueManagementProfile_Red_Uniform_EnableEcnPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Red_Uniform) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.EnableEcn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red_Uniform", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_Red_Uniform_MaxThresholdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Red_Uniform_MaxThresholdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_Red_Uniform_MaxThresholdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold to the batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Red_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Red_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_Red_Uniform_MaxThresholdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/max-threshold to the batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_MaxThresholdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_Red_Uniform_MaxThresholdPath extracts the value of the leaf MaxThreshold from its parent oc.Qos_QueueManagementProfile_Red_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_QueueManagementProfile_Red_Uniform_MaxThresholdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Red_Uniform) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxThreshold
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Red_Uniform", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_Red_Uniform_MinThresholdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Red_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Red_Uniform_MinThresholdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_QueueManagementProfile_Red_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_Red_Uniform_MinThresholdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold to the batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Red_Uniform_MinThresholdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Red_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Red_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Red_Uniform", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_Red_Uniform_MinThresholdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Red_Uniform_MinThresholdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/red/uniform/state/min-threshold to the batch object.
func (n *Qos_QueueManagementProfile_Red_Uniform_MinThresholdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_Red_Uniform_MinThresholdPath extracts the value of the leaf MinThreshold from its parent oc.Qos_QueueManagementProfile_Red_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_QueueManagementProfile_Red_Uniform_MinThresholdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Red_Uniform) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MinThreshold
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_WredPath) Lookup(t testing.TB) *oc.QualifiedQos_QueueManagementProfile_Wred {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_QueueManagementProfile_Wred{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_WredPath) Get(t testing.TB) *oc.Qos_QueueManagementProfile_Wred {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_WredPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_QueueManagementProfile_Wred {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_QueueManagementProfile_Wred
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_QueueManagementProfile_Wred{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a ONCE subscription.
func (n *Qos_QueueManagementProfile_WredPathAny) Get(t testing.TB) []*oc.Qos_QueueManagementProfile_Wred {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_QueueManagementProfile_Wred
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_WredPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_QueueManagementProfile_Wred {
	t.Helper()
	c := &oc.CollectionQos_QueueManagementProfile_Wred{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_QueueManagementProfile_Wred) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_QueueManagementProfile_Wred{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_QueueManagementProfile_Wred)))
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_WredPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Wred) bool) *oc.Qos_QueueManagementProfile_WredWatcher {
	t.Helper()
	w := &oc.Qos_QueueManagementProfile_WredWatcher{}
	gs := &oc.Qos_QueueManagementProfile_Wred{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_QueueManagementProfile_Wred{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_QueueManagementProfile_Wred)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_WredPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Wred) bool) *oc.Qos_QueueManagementProfile_WredWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_WredPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_WredPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_QueueManagementProfile_Wred) *oc.QualifiedQos_QueueManagementProfile_Wred {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_QueueManagementProfile_Wred) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred to the batch object.
func (n *Qos_QueueManagementProfile_WredPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_WredPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_QueueManagementProfile_Wred {
	t.Helper()
	c := &oc.CollectionQos_QueueManagementProfile_Wred{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_QueueManagementProfile_Wred) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_WredPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Wred) bool) *oc.Qos_QueueManagementProfile_WredWatcher {
	t.Helper()
	w := &oc.Qos_QueueManagementProfile_WredWatcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Wred{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Wred{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Wred", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_QueueManagementProfile_Wred{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_QueueManagementProfile_Wred)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_WredPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Wred) bool) *oc.Qos_QueueManagementProfile_WredWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_WredPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred to the batch object.
func (n *Qos_QueueManagementProfile_WredPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Lookup(t testing.TB) *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_QueueManagementProfile_Wred_Uniform{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Get(t testing.TB) *oc.Qos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_UniformPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_QueueManagementProfile_Wred_Uniform
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_QueueManagementProfile_Wred_Uniform{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_UniformPathAny) Get(t testing.TB) []*oc.Qos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_QueueManagementProfile_Wred_Uniform
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	c := &oc.CollectionQos_QueueManagementProfile_Wred_Uniform{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_QueueManagementProfile_Wred_Uniform{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_QueueManagementProfile_Wred_Uniform)))
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_UniformPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform) bool) *oc.Qos_QueueManagementProfile_Wred_UniformWatcher {
	t.Helper()
	w := &oc.Qos_QueueManagementProfile_Wred_UniformWatcher{}
	gs := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_QueueManagementProfile_Wred_Uniform{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_QueueManagementProfile_Wred_Uniform)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform) bool) *oc.Qos_QueueManagementProfile_Wred_UniformWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_UniformPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform to the batch object.
func (n *Qos_QueueManagementProfile_Wred_UniformPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_UniformPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	c := &oc.CollectionQos_QueueManagementProfile_Wred_Uniform{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_UniformPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform) bool) *oc.Qos_QueueManagementProfile_Wred_UniformWatcher {
	t.Helper()
	w := &oc.Qos_QueueManagementProfile_Wred_UniformWatcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Wred_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_QueueManagementProfile_Wred_Uniform{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_QueueManagementProfile_Wred_Uniform)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_UniformPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform) bool) *oc.Qos_QueueManagementProfile_Wred_UniformWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_UniformPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform to the batch object.
func (n *Qos_QueueManagementProfile_Wred_UniformPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_DropPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDrop())
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_DropPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_DropPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_Wred_Uniform_DropPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_DropPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_DropPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Wred_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_Wred_Uniform_DropPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_DropPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_Wred_Uniform_DropPath extracts the value of the leaf Drop from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_QueueManagementProfile_Wred_Uniform_DropPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnableEcn())
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Wred_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath extracts the value of the leaf EnableEcn from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.EnableEcn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Wred_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath extracts the value of the leaf MaxDropProbabilityPercent from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.MaxDropProbabilityPercent
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Wred_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath extracts the value of the leaf MaxThreshold from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxThreshold
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Wred_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath extracts the value of the leaf MinThreshold from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MinThreshold
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	md, ok := oc.Lookup(t, n, "Qos_QueueManagementProfile_Wred_Uniform", goStruct, true, false)
	if ok {
		return convertQos_QueueManagementProfile_Wred_Uniform_WeightPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_QueueManagementProfile_Wred_Uniform_WeightPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight with a ONCE subscription.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_WeightPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_QueueManagementProfile_Wred_Uniform_WeightPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_WeightPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueueManagementProfile_Wred_Uniform_WeightPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Qos_QueueManagementProfile_Wred_Uniform{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_QueueManagementProfile_Wred_Uniform{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", structs[pre], queryPath, true, false)
			qv := convertQos_QueueManagementProfile_Wred_Uniform_WeightPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_WeightPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight to the batch object.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_QueueManagementProfile_Wred_Uniform_WeightPath extracts the value of the leaf Weight from its parent oc.Qos_QueueManagementProfile_Wred_Uniform
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_QueueManagementProfile_Wred_Uniform_WeightPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_QueueManagementProfile_Wred_Uniform) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Weight
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/queues/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_QueuePath) Lookup(t testing.TB) *oc.QualifiedQos_Queue {
	t.Helper()
	goStruct := &oc.Qos_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Queue", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Queue{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queues/queue with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_QueuePath) Get(t testing.TB) *oc.Qos_Queue {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queues/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Queue {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Queue
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Queue", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Queue{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queues/queue with a ONCE subscription.
func (n *Qos_QueuePathAny) Get(t testing.TB) []*oc.Qos_Queue {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Queue
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queues/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Queue {
	t.Helper()
	c := &oc.CollectionQos_Queue{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Queue) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Queue{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Queue)))
		return false
	})
	return c
}

func watch_Qos_QueuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Queue) bool) *oc.Qos_QueueWatcher {
	t.Helper()
	w := &oc.Qos_QueueWatcher{}
	gs := &oc.Qos_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Queue", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Queue{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Queue)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Queue) bool) *oc.Qos_QueueWatcher {
	t.Helper()
	return watch_Qos_QueuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queues/queue with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_QueuePath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Queue) *oc.QualifiedQos_Queue {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Queue) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queues/queue failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queues/queue to the batch object.
func (n *Qos_QueuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queues/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_QueuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Queue {
	t.Helper()
	c := &oc.CollectionQos_Queue{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Queue) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_QueuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Queue) bool) *oc.Qos_QueueWatcher {
	t.Helper()
	w := &oc.Qos_QueueWatcher{}
	structs := map[string]*oc.Qos_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Queue{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Queue", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Queue{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Queue)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Queue) bool) *oc.Qos_QueueWatcher {
	t.Helper()
	return watch_Qos_QueuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queues/queue to the batch object.
func (n *Qos_QueuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
