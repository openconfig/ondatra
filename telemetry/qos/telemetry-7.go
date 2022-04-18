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

// Lookup fetches the value at /openconfig-qos/qos/queues/queue/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Queue_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Queue", goStruct, true, false)
	if ok {
		return convertQos_Queue_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/queues/queue/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Queue_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/queues/queue/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Queue_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Queue_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/queues/queue/state/name with a ONCE subscription.
func (n *Qos_Queue_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queues/queue/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Queue_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Queue_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Queue", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Queue_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Queue_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Queue_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/queues/queue/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Queue_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/queues/queue/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/queues/queue/state/name to the batch object.
func (n *Qos_Queue_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/queues/queue/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Queue_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Queue_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Queue", structs[pre], queryPath, true, false)
			qv := convertQos_Queue_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Queue_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Queue_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/queues/queue/state/name to the batch object.
func (n *Qos_Queue_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Queue_NamePath extracts the value of the leaf Name from its parent oc.Qos_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Queue_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Queue) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicyPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicyPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicyPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a ONCE subscription.
func (n *Qos_SchedulerPolicyPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_SchedulerPolicy{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_SchedulerPolicy)))
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy) bool) *oc.Qos_SchedulerPolicyWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicyWatcher{}
	gs := &oc.Qos_SchedulerPolicy{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_SchedulerPolicy{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy) bool) *oc.Qos_SchedulerPolicyWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicyPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_SchedulerPolicy) *oc.QualifiedQos_SchedulerPolicy {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_SchedulerPolicy) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy to the batch object.
func (n *Qos_SchedulerPolicyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicyPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy) bool) *oc.Qos_SchedulerPolicyWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicyWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_SchedulerPolicy{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy) bool) *oc.Qos_SchedulerPolicyWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicyPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy to the batch object.
func (n *Qos_SchedulerPolicyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_SchedulerPolicy_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name with a ONCE subscription.
func (n *Qos_SchedulerPolicy_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_SchedulerPolicy{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name to the batch object.
func (n *Qos_SchedulerPolicy_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name to the batch object.
func (n *Qos_SchedulerPolicy_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_NamePath extracts the value of the leaf Name from its parent oc.Qos_SchedulerPolicy
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_SchedulerPolicy_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy) *oc.QualifiedString {
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_Qos_SchedulerPolicy_SchedulerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler) bool) *oc.Qos_SchedulerPolicy_SchedulerWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_SchedulerWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *Qos_SchedulerPolicy_SchedulerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler) bool) *oc.Qos_SchedulerPolicy_SchedulerWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_SchedulerPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Input{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_InputPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input) bool) *oc.Qos_SchedulerPolicy_Scheduler_InputWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_InputWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_Input{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_Input{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *Qos_SchedulerPolicy_Scheduler_InputPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_Input) bool) *oc.Qos_SchedulerPolicy_Scheduler_InputWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_InputPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_Input_IdPath(t, md, gs)}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_Input_IdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_Input{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_Input_IdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_IdPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath(t, md, gs)}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_Input_InputTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Input_InputType) bool) *oc.E_Input_InputTypeWatcher {
	t.Helper()
	w := &oc.E_Input_InputTypeWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_Input{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_Input_InputTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *Qos_SchedulerPolicy_Scheduler_Input_InputTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Input_InputType) bool) *oc.E_Input_InputTypeWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_InputTypePathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_Input_QueuePath(t, md, gs)}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_Input_QueuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_Input{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_Input_QueuePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_QueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_QueuePathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_Input_WeightPath(t, md, gs)}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_Input_WeightPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_SchedulerPolicy_Scheduler_Input{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_Input", structs[pre], queryPath, true, false)
			qv := convertQos_SchedulerPolicy_Scheduler_Input_WeightPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/inputs/input/state/weight with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_Input_WeightPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_Input_WeightPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColorPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t, md, gs)}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
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
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/bc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_BcPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t, md, gs)}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t, md, gs)}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
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
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t, md, gs)}, nil
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

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
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
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/state/cir-pct-remaining with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_CirPctRemainingPathAny(t, n, timeout, predicate)
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Lookup(t testing.TB) *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Get(t testing.TB) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPathAny) Get(t testing.TB) []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction)))
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionWatcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	c := &oc.CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionWatcher {
	t.Helper()
	w := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionWatcher{}
	structs := map[string]*oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) bool) *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
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
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
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
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dot1p to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	md, ok := oc.Lookup(t, n, "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", goStruct, true, false)
	if ok {
		return convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
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
		qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp with a ONCE subscription.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
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
			qv := convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/scheduler-policies/scheduler-policy/schedulers/scheduler/one-rate-two-color/conform-action/state/set-dscp to the batch object.
func (n *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) *oc.QualifiedUint8 {
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
