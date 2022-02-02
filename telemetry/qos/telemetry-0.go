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

// Lookup fetches the value at /openconfig-qos/qos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *QosPath) Lookup(t testing.TB) *oc.QualifiedQos {
	t.Helper()
	goStruct := &oc.Qos{}
	md, ok := oc.Lookup(t, n, "Qos", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *QosPath) Get(t testing.TB) *oc.Qos {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *QosPathAny) Lookup(t testing.TB) []*oc.QualifiedQos {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos with a ONCE subscription.
func (n *QosPathAny) Get(t testing.TB) []*oc.Qos {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *QosPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos {
	t.Helper()
	c := &oc.CollectionQos{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos)))
		return false
	})
	return c
}

func watch_QosPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos) bool) *oc.QosWatcher {
	t.Helper()
	w := &oc.QosWatcher{}
	gs := &oc.Qos{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos", gs, queryPath, false, false)
		return (&oc.QualifiedQos{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *QosPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos) bool) *oc.QosWatcher {
	t.Helper()
	return watch_QosPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *QosPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos) *oc.QualifiedQos {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos to the batch object.
func (n *QosPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *QosPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos {
	t.Helper()
	c := &oc.CollectionQos{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *QosPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos) bool) *oc.QosWatcher {
	t.Helper()
	return watch_QosPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos to the batch object.
func (n *QosPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedQos_BufferAllocationProfile {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_BufferAllocationProfile{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfilePath) Get(t testing.TB) *oc.Qos_BufferAllocationProfile {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_BufferAllocationProfile {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_BufferAllocationProfile
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_BufferAllocationProfile{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a ONCE subscription.
func (n *Qos_BufferAllocationProfilePathAny) Get(t testing.TB) []*oc.Qos_BufferAllocationProfile {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_BufferAllocationProfile
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfilePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_BufferAllocationProfile {
	t.Helper()
	c := &oc.CollectionQos_BufferAllocationProfile{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_BufferAllocationProfile) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_BufferAllocationProfile{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_BufferAllocationProfile)))
		return false
	})
	return c
}

func watch_Qos_BufferAllocationProfilePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_BufferAllocationProfile) bool) *oc.Qos_BufferAllocationProfileWatcher {
	t.Helper()
	w := &oc.Qos_BufferAllocationProfileWatcher{}
	gs := &oc.Qos_BufferAllocationProfile{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_BufferAllocationProfile", gs, queryPath, false, false)
		return (&oc.QualifiedQos_BufferAllocationProfile{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_BufferAllocationProfile)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfilePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_BufferAllocationProfile) bool) *oc.Qos_BufferAllocationProfileWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfilePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_BufferAllocationProfilePath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_BufferAllocationProfile) *oc.QualifiedQos_BufferAllocationProfile {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_BufferAllocationProfile) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile to the batch object.
func (n *Qos_BufferAllocationProfilePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfilePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_BufferAllocationProfile {
	t.Helper()
	c := &oc.CollectionQos_BufferAllocationProfile{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_BufferAllocationProfile) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfilePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_BufferAllocationProfile) bool) *oc.Qos_BufferAllocationProfileWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfilePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile to the batch object.
func (n *Qos_BufferAllocationProfilePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile", goStruct, true, false)
	if ok {
		return convertQos_BufferAllocationProfile_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_BufferAllocationProfile_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_BufferAllocationProfile{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_BufferAllocationProfile", gs, queryPath, true, false)
		return convertQos_BufferAllocationProfile_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_BufferAllocationProfile_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name to the batch object.
func (n *Qos_BufferAllocationProfile_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/state/name to the batch object.
func (n *Qos_BufferAllocationProfile_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_BufferAllocationProfile_NamePath extracts the value of the leaf Name from its parent oc.Qos_BufferAllocationProfile
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_BufferAllocationProfile_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_QueuePath) Lookup(t testing.TB) *oc.QualifiedQos_BufferAllocationProfile_Queue {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_BufferAllocationProfile_Queue{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_QueuePath) Get(t testing.TB) *oc.Qos_BufferAllocationProfile_Queue {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_BufferAllocationProfile_Queue {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_BufferAllocationProfile_Queue
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_BufferAllocationProfile_Queue{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_QueuePathAny) Get(t testing.TB) []*oc.Qos_BufferAllocationProfile_Queue {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_BufferAllocationProfile_Queue
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_QueuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_BufferAllocationProfile_Queue {
	t.Helper()
	c := &oc.CollectionQos_BufferAllocationProfile_Queue{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_BufferAllocationProfile_Queue) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_BufferAllocationProfile_Queue{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_BufferAllocationProfile_Queue)))
		return false
	})
	return c
}

func watch_Qos_BufferAllocationProfile_QueuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_BufferAllocationProfile_Queue) bool) *oc.Qos_BufferAllocationProfile_QueueWatcher {
	t.Helper()
	w := &oc.Qos_BufferAllocationProfile_QueueWatcher{}
	gs := &oc.Qos_BufferAllocationProfile_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", gs, queryPath, false, false)
		return (&oc.QualifiedQos_BufferAllocationProfile_Queue{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_BufferAllocationProfile_Queue)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_QueuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_BufferAllocationProfile_Queue) bool) *oc.Qos_BufferAllocationProfile_QueueWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_QueuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_BufferAllocationProfile_QueuePath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedQos_BufferAllocationProfile_Queue {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_BufferAllocationProfile_Queue) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue to the batch object.
func (n *Qos_BufferAllocationProfile_QueuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_QueuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_BufferAllocationProfile_Queue {
	t.Helper()
	c := &oc.CollectionQos_BufferAllocationProfile_Queue{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_BufferAllocationProfile_Queue) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_QueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_BufferAllocationProfile_Queue) bool) *oc.Qos_BufferAllocationProfile_QueueWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_QueuePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue to the batch object.
func (n *Qos_BufferAllocationProfile_QueuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, false)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_DedicatedBufferPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_DedicatedBufferPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_BufferAllocationProfile_Queue_DedicatedBufferPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Qos_BufferAllocationProfile_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", gs, queryPath, true, false)
		return convertQos_BufferAllocationProfile_Queue_DedicatedBufferPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_DedicatedBufferPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_DedicatedBufferPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dedicated-buffer to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_DedicatedBufferPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_BufferAllocationProfile_Queue_DedicatedBufferPath extracts the value of the leaf DedicatedBuffer from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertQos_BufferAllocationProfile_Queue_DedicatedBufferPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.DedicatedBuffer
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Lookup(t testing.TB) *oc.QualifiedInt32 {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, false)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Get(t testing.TB) int32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPathAny) Lookup(t testing.TB) []*oc.QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPathAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt32 {
	t.Helper()
	c := &oc.CollectionInt32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	w := &oc.Int32Watcher{}
	gs := &oc.Qos_BufferAllocationProfile_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", gs, queryPath, true, false)
		return convertQos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Await(t testing.TB, timeout time.Duration, val int32) *oc.QualifiedInt32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt32 {
	t.Helper()
	c := &oc.CollectionInt32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/dynamic-limit-scaling-factor to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath extracts the value of the leaf DynamicLimitScalingFactor from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedInt32.
func convertQos_BufferAllocationProfile_Queue_DynamicLimitScalingFactorPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedInt32 {
	t.Helper()
	qv := &oc.QualifiedInt32{
		Metadata: md,
	}
	val := parent.DynamicLimitScalingFactor
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, false)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_BufferAllocationProfile_Queue_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_BufferAllocationProfile_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", gs, queryPath, true, false)
		return convertQos_BufferAllocationProfile_Queue_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/name to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_BufferAllocationProfile_Queue_NamePath extracts the value of the leaf Name from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_BufferAllocationProfile_Queue_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Lookup(t testing.TB) *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, false)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Get(t testing.TB) oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePathAny) Get(t testing.TB) []oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	c := &oc.CollectionE_Qos_SHARED_BUFFER_LIMIT_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE) bool) *oc.E_Qos_SHARED_BUFFER_LIMIT_TYPEWatcher {
	t.Helper()
	w := &oc.E_Qos_SHARED_BUFFER_LIMIT_TYPEWatcher{}
	gs := &oc.Qos_BufferAllocationProfile_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", gs, queryPath, true, false)
		return convertQos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE) bool) *oc.E_Qos_SHARED_BUFFER_LIMIT_TYPEWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Qos_SHARED_BUFFER_LIMIT_TYPE) *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	c := &oc.CollectionE_Qos_SHARED_BUFFER_LIMIT_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE) bool) *oc.E_Qos_SHARED_BUFFER_LIMIT_TYPEWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/shared-buffer-limit-type to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_SharedBufferLimitTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath extracts the value of the leaf SharedBufferLimitType from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE.
func convertQos_BufferAllocationProfile_Queue_SharedBufferLimitTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE{
		Metadata: md,
	}
	val := parent.SharedBufferLimitType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, false)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_BufferAllocationProfile_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", gs, queryPath, true, false)
		return convertQos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/static-shared-buffer-limit to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath extracts the value of the leaf StaticSharedBufferLimit from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_BufferAllocationProfile_Queue_StaticSharedBufferLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.StaticSharedBufferLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_BufferAllocationProfile_Queue", goStruct, true, false)
	if ok {
		return convertQos_BufferAllocationProfile_Queue_UseSharedBufferPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_BufferAllocationProfile_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_BufferAllocationProfile_Queue_UseSharedBufferPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer with a ONCE subscription.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_BufferAllocationProfile_Queue_UseSharedBufferPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Qos_BufferAllocationProfile_Queue{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_BufferAllocationProfile_Queue", gs, queryPath, true, false)
		return convertQos_BufferAllocationProfile_Queue_UseSharedBufferPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_UseSharedBufferPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Qos_BufferAllocationProfile_Queue_UseSharedBufferPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/buffer-allocation-profiles/buffer-allocation-profile/queues/queue/state/use-shared-buffer to the batch object.
func (n *Qos_BufferAllocationProfile_Queue_UseSharedBufferPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_BufferAllocationProfile_Queue_UseSharedBufferPath extracts the value of the leaf UseSharedBuffer from its parent oc.Qos_BufferAllocationProfile_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertQos_BufferAllocationProfile_Queue_UseSharedBufferPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_BufferAllocationProfile_Queue) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.UseSharedBuffer
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ClassifierPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier {
	t.Helper()
	goStruct := &oc.Qos_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ClassifierPath) Get(t testing.TB) *oc.Qos_Classifier {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ClassifierPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier with a ONCE subscription.
func (n *Qos_ClassifierPathAny) Get(t testing.TB) []*oc.Qos_Classifier {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ClassifierPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier {
	t.Helper()
	c := &oc.CollectionQos_Classifier{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier)))
		return false
	})
	return c
}

func watch_Qos_ClassifierPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier) bool) *oc.Qos_ClassifierWatcher {
	t.Helper()
	w := &oc.Qos_ClassifierWatcher{}
	gs := &oc.Qos_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ClassifierPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier) bool) *oc.Qos_ClassifierWatcher {
	t.Helper()
	return watch_Qos_ClassifierPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_ClassifierPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier) *oc.QualifiedQos_Classifier {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier to the batch object.
func (n *Qos_ClassifierPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ClassifierPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier {
	t.Helper()
	c := &oc.CollectionQos_Classifier{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ClassifierPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier) bool) *oc.Qos_ClassifierWatcher {
	t.Helper()
	return watch_Qos_ClassifierPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier to the batch object.
func (n *Qos_ClassifierPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier", goStruct, true, false)
	if ok {
		return convertQos_Classifier_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/state/name with a ONCE subscription.
func (n *Qos_Classifier_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier", gs, queryPath, true, false)
		return convertQos_Classifier_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/state/name to the batch object.
func (n *Qos_Classifier_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/state/name to the batch object.
func (n *Qos_Classifier_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_NamePath extracts the value of the leaf Name from its parent oc.Qos_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_TermPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_TermPath) Get(t testing.TB) *oc.Qos_Classifier_Term {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_TermPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a ONCE subscription.
func (n *Qos_Classifier_TermPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_TermPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term)))
		return false
	})
	return c
}

func watch_Qos_Classifier_TermPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term) bool) *oc.Qos_Classifier_TermWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_TermWatcher{}
	gs := &oc.Qos_Classifier_Term{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_TermPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term) bool) *oc.Qos_Classifier_TermWatcher {
	t.Helper()
	return watch_Qos_Classifier_TermPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_TermPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term) *oc.QualifiedQos_Classifier_Term {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term to the batch object.
func (n *Qos_Classifier_TermPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_TermPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_TermPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term) bool) *oc.Qos_Classifier_TermWatcher {
	t.Helper()
	return watch_Qos_Classifier_TermPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term to the batch object.
func (n *Qos_Classifier_TermPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_ActionsPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Actions {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Actions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_ActionsPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Actions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_ActionsPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Actions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Actions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Actions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a ONCE subscription.
func (n *Qos_Classifier_Term_ActionsPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Actions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Actions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_ActionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Actions {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Actions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Actions) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Actions{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Actions)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_ActionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions) bool) *oc.Qos_Classifier_Term_ActionsWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_ActionsWatcher{}
	gs := &oc.Qos_Classifier_Term_Actions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term_Actions{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Actions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_ActionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions) bool) *oc.Qos_Classifier_Term_ActionsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_ActionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_ActionsPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Actions) *oc.QualifiedQos_Classifier_Term_Actions {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Actions) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions to the batch object.
func (n *Qos_Classifier_Term_ActionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_ActionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Actions {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Actions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Actions) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_ActionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions) bool) *oc.Qos_Classifier_Term_ActionsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_ActionsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions to the batch object.
func (n *Qos_Classifier_Term_ActionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Actions_Remark {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Actions_Remark{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Actions_Remark {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Actions_Remark {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Actions_Remark
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Actions_Remark{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Actions_Remark {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Actions_Remark
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Actions_Remark {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Actions_Remark{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Actions_Remark{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Actions_Remark)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Actions_RemarkPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool) *oc.Qos_Classifier_Term_Actions_RemarkWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Actions_RemarkWatcher{}
	gs := &oc.Qos_Classifier_Term_Actions_Remark{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term_Actions_Remark{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Actions_Remark)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool) *oc.Qos_Classifier_Term_Actions_RemarkWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_RemarkPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedQos_Classifier_Term_Actions_Remark {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark to the batch object.
func (n *Qos_Classifier_Term_Actions_RemarkPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Actions_Remark {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Actions_Remark{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Actions_Remark) bool) *oc.Qos_Classifier_Term_Actions_RemarkWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_RemarkPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark to the batch object.
func (n *Qos_Classifier_Term_Actions_RemarkPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Actions_Remark_SetDot1PPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_Classifier_Term_Actions_Remark{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetDot1PPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetDot1PPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dot1p to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDot1PPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Actions_Remark_SetDot1PPath extracts the value of the leaf SetDot1P from its parent oc.Qos_Classifier_Term_Actions_Remark
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Actions_Remark_SetDot1PPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Actions_Remark_SetDscpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_Classifier_Term_Actions_Remark{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetDscpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetDscpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-dscp to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetDscpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Actions_Remark_SetDscpPath extracts the value of the leaf SetDscp from its parent oc.Qos_Classifier_Term_Actions_Remark
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Actions_Remark_SetDscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions_Remark", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions_Remark{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Actions_Remark_SetMplsTcPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_Classifier_Term_Actions_Remark{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions_Remark", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_Remark_SetMplsTcPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/remark/state/set-mpls-tc to the batch object.
func (n *Qos_Classifier_Term_Actions_Remark_SetMplsTcPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath extracts the value of the leaf SetMplsTc from its parent oc.Qos_Classifier_Term_Actions_Remark
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Actions_Remark_SetMplsTcPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions_Remark) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Actions{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Actions", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Actions_TargetGroupPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Actions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Actions_TargetGroupPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a ONCE subscription.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Actions_TargetGroupPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier_Term_Actions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Actions", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Actions_TargetGroupPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_TargetGroupPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group to the batch object.
func (n *Qos_Classifier_Term_Actions_TargetGroupPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Actions_TargetGroupPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/actions/state/target-group to the batch object.
func (n *Qos_Classifier_Term_Actions_TargetGroupPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Actions_TargetGroupPath extracts the value of the leaf TargetGroup from its parent oc.Qos_Classifier_Term_Actions
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Actions_TargetGroupPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Actions) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TargetGroup
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_ConditionsPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_ConditionsPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_ConditionsPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a ONCE subscription.
func (n *Qos_Classifier_Term_ConditionsPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_ConditionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Conditions{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Conditions)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_ConditionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions) bool) *oc.Qos_Classifier_Term_ConditionsWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_ConditionsWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term_Conditions{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_ConditionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions) bool) *oc.Qos_Classifier_Term_ConditionsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_ConditionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_ConditionsPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Conditions) *oc.QualifiedQos_Classifier_Term_Conditions {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions to the batch object.
func (n *Qos_Classifier_Term_ConditionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_ConditionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_ConditionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions) bool) *oc.Qos_Classifier_Term_ConditionsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_ConditionsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions to the batch object.
func (n *Qos_Classifier_Term_ConditionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv4", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv4
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_Ipv4
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Ipv4{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv4{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Conditions_Ipv4)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Ipv4Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool) *oc.Qos_Classifier_Term_Conditions_Ipv4Watcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Ipv4Watcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv4", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Ipv4{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Ipv4)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool) *oc.Qos_Classifier_Term_Conditions_Ipv4Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv4Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Conditions_Ipv4) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Ipv4{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Ipv4) bool) *oc.Qos_Classifier_Term_Conditions_Ipv4Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv4Path(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv4 to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv4PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
