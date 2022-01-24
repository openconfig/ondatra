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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, false, false)
		return (&oc.QualifiedQos_QueueManagementProfile_Wred_Uniform{
			Metadata: md,
		}).SetVal(gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_UniformPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_QueueManagementProfile_Wred_Uniform) bool) *oc.Qos_QueueManagementProfile_Wred_UniformWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_UniformPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return convertQos_QueueManagementProfile_Wred_Uniform_DropPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/drop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_DropPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_DropPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return convertQos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/enable-ecn with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_EnableEcnPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return convertQos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-drop-probability-percent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_MaxDropProbabilityPercentPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return convertQos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/max-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_MaxThresholdPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return convertQos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/min-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_MinThresholdPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_QueueManagementProfile_Wred_Uniform", gs, queryPath, true, false)
		return convertQos_QueueManagementProfile_Wred_Uniform_WeightPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queue-management-profiles/queue-management-profile/wred/uniform/state/weight with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueueManagementProfile_Wred_Uniform_WeightPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_QueueManagementProfile_Wred_Uniform_WeightPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Queue", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Queue{
			Metadata: md,
		}).SetVal(gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_QueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Queue) bool) *oc.Qos_QueueWatcher {
	t.Helper()
	return watch_Qos_QueuePath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Queue", gs, queryPath, true, false)
		return convertQos_Queue_NamePath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Queue_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Queue_NamePath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy", gs, queryPath, false, false)
		return (&oc.QualifiedQos_SchedulerPolicy{
			Metadata: md,
		}).SetVal(gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_SchedulerPolicy) bool) *oc.Qos_SchedulerPolicyWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicyPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_SchedulerPolicy", gs, queryPath, true, false)
		return convertQos_SchedulerPolicy_NamePath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/scheduler-policies/scheduler-policy/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_SchedulerPolicy_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_SchedulerPolicy_NamePath(t, n, timeout, predicate)
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
