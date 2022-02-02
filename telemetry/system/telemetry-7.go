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

// Lookup fetches the value at /openconfig-system/system/memory with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MemoryPath) Lookup(t testing.TB) *oc.QualifiedSystem_Memory {
	t.Helper()
	goStruct := &oc.System_Memory{}
	md, ok := oc.Lookup(t, n, "System_Memory", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Memory{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/memory with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MemoryPath) Get(t testing.TB) *oc.System_Memory {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/memory with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MemoryPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Memory {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Memory
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Memory", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Memory{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/memory with a ONCE subscription.
func (n *System_MemoryPathAny) Get(t testing.TB) []*oc.System_Memory {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Memory
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MemoryPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Memory {
	t.Helper()
	c := &oc.CollectionSystem_Memory{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Memory) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Memory{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Memory)))
		return false
	})
	return c
}

func watch_System_MemoryPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Memory) bool) *oc.System_MemoryWatcher {
	t.Helper()
	w := &oc.System_MemoryWatcher{}
	gs := &oc.System_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Memory", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Memory{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Memory)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MemoryPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Memory) bool) *oc.System_MemoryWatcher {
	t.Helper()
	return watch_System_MemoryPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/memory with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_MemoryPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Memory) *oc.QualifiedSystem_Memory {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Memory) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/memory failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/memory to the batch object.
func (n *System_MemoryPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MemoryPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Memory {
	t.Helper()
	c := &oc.CollectionSystem_Memory{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Memory) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MemoryPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Memory) bool) *oc.System_MemoryWatcher {
	t.Helper()
	return watch_System_MemoryPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/memory to the batch object.
func (n *System_MemoryPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/memory/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Memory_CountersPath) Lookup(t testing.TB) *oc.QualifiedSystem_Memory_Counters {
	t.Helper()
	goStruct := &oc.System_Memory_Counters{}
	md, ok := oc.Lookup(t, n, "System_Memory_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Memory_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/memory/state/counters with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Memory_CountersPath) Get(t testing.TB) *oc.System_Memory_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/memory/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Memory_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Memory_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Memory_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Memory_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Memory_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Memory_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/memory/state/counters with a ONCE subscription.
func (n *System_Memory_CountersPathAny) Get(t testing.TB) []*oc.System_Memory_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Memory_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Memory_Counters {
	t.Helper()
	c := &oc.CollectionSystem_Memory_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Memory_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Memory_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Memory_Counters)))
		return false
	})
	return c
}

func watch_System_Memory_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Memory_Counters) bool) *oc.System_Memory_CountersWatcher {
	t.Helper()
	w := &oc.System_Memory_CountersWatcher{}
	gs := &oc.System_Memory_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Memory_Counters", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Memory_Counters{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Memory_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Memory_Counters) bool) *oc.System_Memory_CountersWatcher {
	t.Helper()
	return watch_System_Memory_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/memory/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Memory_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Memory_Counters) *oc.QualifiedSystem_Memory_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Memory_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/memory/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/memory/state/counters to the batch object.
func (n *System_Memory_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Memory_Counters {
	t.Helper()
	c := &oc.CollectionSystem_Memory_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Memory_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Memory_Counters) bool) *oc.System_Memory_CountersWatcher {
	t.Helper()
	return watch_System_Memory_CountersPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/memory/state/counters to the batch object.
func (n *System_Memory_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/memory/state/counters/correctable-ecc-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Memory_Counters_CorrectableEccErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Memory_Counters{}
	md, ok := oc.Lookup(t, n, "System_Memory_Counters", goStruct, true, false)
	if ok {
		return convertSystem_Memory_Counters_CorrectableEccErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/memory/state/counters/correctable-ecc-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Memory_Counters_CorrectableEccErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/memory/state/counters/correctable-ecc-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Memory_Counters_CorrectableEccErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Memory_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Memory_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Memory_Counters_CorrectableEccErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/memory/state/counters/correctable-ecc-errors with a ONCE subscription.
func (n *System_Memory_Counters_CorrectableEccErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/counters/correctable-ecc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_Counters_CorrectableEccErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Memory_Counters_CorrectableEccErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Memory_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Memory_Counters", gs, queryPath, true, false)
		return convertSystem_Memory_Counters_CorrectableEccErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/counters/correctable-ecc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_Counters_CorrectableEccErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_Counters_CorrectableEccErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/memory/state/counters/correctable-ecc-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Memory_Counters_CorrectableEccErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/memory/state/counters/correctable-ecc-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/memory/state/counters/correctable-ecc-errors to the batch object.
func (n *System_Memory_Counters_CorrectableEccErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/counters/correctable-ecc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_Counters_CorrectableEccErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/counters/correctable-ecc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_Counters_CorrectableEccErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_Counters_CorrectableEccErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/memory/state/counters/correctable-ecc-errors to the batch object.
func (n *System_Memory_Counters_CorrectableEccErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Memory_Counters_CorrectableEccErrorsPath extracts the value of the leaf CorrectableEccErrors from its parent oc.System_Memory_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Memory_Counters_CorrectableEccErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Memory_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CorrectableEccErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/memory/state/counters/total-ecc-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Memory_Counters_TotalEccErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Memory_Counters{}
	md, ok := oc.Lookup(t, n, "System_Memory_Counters", goStruct, true, false)
	if ok {
		return convertSystem_Memory_Counters_TotalEccErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/memory/state/counters/total-ecc-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Memory_Counters_TotalEccErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/memory/state/counters/total-ecc-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Memory_Counters_TotalEccErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Memory_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Memory_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Memory_Counters_TotalEccErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/memory/state/counters/total-ecc-errors with a ONCE subscription.
func (n *System_Memory_Counters_TotalEccErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/counters/total-ecc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_Counters_TotalEccErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Memory_Counters_TotalEccErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Memory_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Memory_Counters", gs, queryPath, true, false)
		return convertSystem_Memory_Counters_TotalEccErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/counters/total-ecc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_Counters_TotalEccErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_Counters_TotalEccErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/memory/state/counters/total-ecc-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Memory_Counters_TotalEccErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/memory/state/counters/total-ecc-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/memory/state/counters/total-ecc-errors to the batch object.
func (n *System_Memory_Counters_TotalEccErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/counters/total-ecc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_Counters_TotalEccErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/counters/total-ecc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_Counters_TotalEccErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_Counters_TotalEccErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/memory/state/counters/total-ecc-errors to the batch object.
func (n *System_Memory_Counters_TotalEccErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Memory_Counters_TotalEccErrorsPath extracts the value of the leaf TotalEccErrors from its parent oc.System_Memory_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Memory_Counters_TotalEccErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Memory_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TotalEccErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Memory_Counters_UncorrectableEccErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Memory_Counters{}
	md, ok := oc.Lookup(t, n, "System_Memory_Counters", goStruct, true, false)
	if ok {
		return convertSystem_Memory_Counters_UncorrectableEccErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Memory_Counters_UncorrectableEccErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Memory_Counters_UncorrectableEccErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Memory_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Memory_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Memory_Counters_UncorrectableEccErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors with a ONCE subscription.
func (n *System_Memory_Counters_UncorrectableEccErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_Counters_UncorrectableEccErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Memory_Counters_UncorrectableEccErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Memory_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Memory_Counters", gs, queryPath, true, false)
		return convertSystem_Memory_Counters_UncorrectableEccErrorsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_Counters_UncorrectableEccErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_Counters_UncorrectableEccErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Memory_Counters_UncorrectableEccErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors to the batch object.
func (n *System_Memory_Counters_UncorrectableEccErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_Counters_UncorrectableEccErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_Counters_UncorrectableEccErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_Counters_UncorrectableEccErrorsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/memory/state/counters/uncorrectable-ecc-errors to the batch object.
func (n *System_Memory_Counters_UncorrectableEccErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Memory_Counters_UncorrectableEccErrorsPath extracts the value of the leaf UncorrectableEccErrors from its parent oc.System_Memory_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Memory_Counters_UncorrectableEccErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Memory_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.UncorrectableEccErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/memory/state/free with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Memory_FreePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Memory{}
	md, ok := oc.Lookup(t, n, "System_Memory", goStruct, true, false)
	if ok {
		return convertSystem_Memory_FreePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/memory/state/free with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Memory_FreePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/memory/state/free with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Memory_FreePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Memory", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Memory_FreePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/memory/state/free with a ONCE subscription.
func (n *System_Memory_FreePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/free with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_FreePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Memory_FreePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Memory", gs, queryPath, true, false)
		return convertSystem_Memory_FreePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/free with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_FreePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_FreePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/memory/state/free with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Memory_FreePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/memory/state/free failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/memory/state/free to the batch object.
func (n *System_Memory_FreePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/free with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_FreePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/free with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_FreePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_FreePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/memory/state/free to the batch object.
func (n *System_Memory_FreePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Memory_FreePath extracts the value of the leaf Free from its parent oc.System_Memory
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Memory_FreePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Memory) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Free
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/memory/state/physical with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Memory_PhysicalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Memory{}
	md, ok := oc.Lookup(t, n, "System_Memory", goStruct, true, false)
	if ok {
		return convertSystem_Memory_PhysicalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/memory/state/physical with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Memory_PhysicalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/memory/state/physical with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Memory_PhysicalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Memory", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Memory_PhysicalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/memory/state/physical with a ONCE subscription.
func (n *System_Memory_PhysicalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/physical with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_PhysicalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Memory_PhysicalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Memory", gs, queryPath, true, false)
		return convertSystem_Memory_PhysicalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/physical with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_PhysicalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_PhysicalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/memory/state/physical with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Memory_PhysicalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/memory/state/physical failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/memory/state/physical to the batch object.
func (n *System_Memory_PhysicalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/physical with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_PhysicalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/physical with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_PhysicalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_PhysicalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/memory/state/physical to the batch object.
func (n *System_Memory_PhysicalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Memory_PhysicalPath extracts the value of the leaf Physical from its parent oc.System_Memory
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Memory_PhysicalPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Memory) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Physical
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/memory/state/reserved with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Memory_ReservedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Memory{}
	md, ok := oc.Lookup(t, n, "System_Memory", goStruct, true, false)
	if ok {
		return convertSystem_Memory_ReservedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/memory/state/reserved with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Memory_ReservedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/memory/state/reserved with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Memory_ReservedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Memory", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Memory_ReservedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/memory/state/reserved with a ONCE subscription.
func (n *System_Memory_ReservedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/reserved with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_ReservedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Memory_ReservedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Memory", gs, queryPath, true, false)
		return convertSystem_Memory_ReservedPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/reserved with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_ReservedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_ReservedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/memory/state/reserved with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Memory_ReservedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/memory/state/reserved failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/memory/state/reserved to the batch object.
func (n *System_Memory_ReservedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/reserved with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_ReservedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/reserved with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_ReservedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_ReservedPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/memory/state/reserved to the batch object.
func (n *System_Memory_ReservedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Memory_ReservedPath extracts the value of the leaf Reserved from its parent oc.System_Memory
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Memory_ReservedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Memory) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Reserved
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/memory/state/used with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Memory_UsedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Memory{}
	md, ok := oc.Lookup(t, n, "System_Memory", goStruct, true, false)
	if ok {
		return convertSystem_Memory_UsedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/memory/state/used with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Memory_UsedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/memory/state/used with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Memory_UsedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Memory", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Memory_UsedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/memory/state/used with a ONCE subscription.
func (n *System_Memory_UsedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/used with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_UsedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Memory_UsedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Memory", gs, queryPath, true, false)
		return convertSystem_Memory_UsedPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/used with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_UsedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_UsedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/memory/state/used with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Memory_UsedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/memory/state/used failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/memory/state/used to the batch object.
func (n *System_Memory_UsedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/memory/state/used with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Memory_UsedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/memory/state/used with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Memory_UsedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Memory_UsedPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/memory/state/used to the batch object.
func (n *System_Memory_UsedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Memory_UsedPath extracts the value of the leaf Used from its parent oc.System_Memory
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Memory_UsedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Memory) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Used
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/messages with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MessagesPath) Lookup(t testing.TB) *oc.QualifiedSystem_Messages {
	t.Helper()
	goStruct := &oc.System_Messages{}
	md, ok := oc.Lookup(t, n, "System_Messages", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Messages{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MessagesPath) Get(t testing.TB) *oc.System_Messages {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MessagesPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Messages {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Messages
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Messages{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages with a ONCE subscription.
func (n *System_MessagesPathAny) Get(t testing.TB) []*oc.System_Messages {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Messages
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MessagesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Messages {
	t.Helper()
	c := &oc.CollectionSystem_Messages{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Messages) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Messages{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Messages)))
		return false
	})
	return c
}

func watch_System_MessagesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Messages) bool) *oc.System_MessagesWatcher {
	t.Helper()
	w := &oc.System_MessagesWatcher{}
	gs := &oc.System_Messages{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Messages{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Messages)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MessagesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Messages) bool) *oc.System_MessagesWatcher {
	t.Helper()
	return watch_System_MessagesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_MessagesPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Messages) *oc.QualifiedSystem_Messages {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Messages) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages to the batch object.
func (n *System_MessagesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MessagesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Messages {
	t.Helper()
	c := &oc.CollectionSystem_Messages{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Messages) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MessagesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Messages) bool) *oc.System_MessagesWatcher {
	t.Helper()
	return watch_System_MessagesPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages to the batch object.
func (n *System_MessagesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/messages/debug-entries/debug-service with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_DebugServicePath) Lookup(t testing.TB) *oc.QualifiedSystem_Messages_DebugService {
	t.Helper()
	goStruct := &oc.System_Messages_DebugService{}
	md, ok := oc.Lookup(t, n, "System_Messages_DebugService", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Messages_DebugService{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/debug-entries/debug-service with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_DebugServicePath) Get(t testing.TB) *oc.System_Messages_DebugService {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/debug-entries/debug-service with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_DebugServicePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Messages_DebugService {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Messages_DebugService
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_DebugService{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_DebugService", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Messages_DebugService{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/debug-entries/debug-service with a ONCE subscription.
func (n *System_Messages_DebugServicePathAny) Get(t testing.TB) []*oc.System_Messages_DebugService {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Messages_DebugService
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/debug-entries/debug-service with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_DebugServicePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Messages_DebugService {
	t.Helper()
	c := &oc.CollectionSystem_Messages_DebugService{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Messages_DebugService) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Messages_DebugService{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Messages_DebugService)))
		return false
	})
	return c
}

func watch_System_Messages_DebugServicePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Messages_DebugService) bool) *oc.System_Messages_DebugServiceWatcher {
	t.Helper()
	w := &oc.System_Messages_DebugServiceWatcher{}
	gs := &oc.System_Messages_DebugService{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages_DebugService", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Messages_DebugService{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Messages_DebugService)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/debug-entries/debug-service with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_DebugServicePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Messages_DebugService) bool) *oc.System_Messages_DebugServiceWatcher {
	t.Helper()
	return watch_System_Messages_DebugServicePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages/debug-entries/debug-service with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Messages_DebugServicePath) Await(t testing.TB, timeout time.Duration, val *oc.System_Messages_DebugService) *oc.QualifiedSystem_Messages_DebugService {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Messages_DebugService) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages/debug-entries/debug-service failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages/debug-entries/debug-service to the batch object.
func (n *System_Messages_DebugServicePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/debug-entries/debug-service with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_DebugServicePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Messages_DebugService {
	t.Helper()
	c := &oc.CollectionSystem_Messages_DebugService{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Messages_DebugService) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/debug-entries/debug-service with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_DebugServicePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Messages_DebugService) bool) *oc.System_Messages_DebugServiceWatcher {
	t.Helper()
	return watch_System_Messages_DebugServicePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages/debug-entries/debug-service to the batch object.
func (n *System_Messages_DebugServicePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/messages/debug-entries/debug-service/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_DebugService_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_Messages_DebugService{}
	md, ok := oc.Lookup(t, n, "System_Messages_DebugService", goStruct, true, false)
	if ok {
		return convertSystem_Messages_DebugService_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-system/system/messages/debug-entries/debug-service/state/enabled with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_DebugService_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/debug-entries/debug-service/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_DebugService_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_DebugService{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_DebugService", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_DebugService_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/debug-entries/debug-service/state/enabled with a ONCE subscription.
func (n *System_Messages_DebugService_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/debug-entries/debug-service/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_DebugService_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Messages_DebugService_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_Messages_DebugService{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages_DebugService", gs, queryPath, true, false)
		return convertSystem_Messages_DebugService_EnabledPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/debug-entries/debug-service/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_DebugService_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_Messages_DebugService_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages/debug-entries/debug-service/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Messages_DebugService_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages/debug-entries/debug-service/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages/debug-entries/debug-service/state/enabled to the batch object.
func (n *System_Messages_DebugService_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/debug-entries/debug-service/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_DebugService_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/debug-entries/debug-service/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_DebugService_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_Messages_DebugService_EnabledPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages/debug-entries/debug-service/state/enabled to the batch object.
func (n *System_Messages_DebugService_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Messages_DebugService_EnabledPath extracts the value of the leaf Enabled from its parent oc.System_Messages_DebugService
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_Messages_DebugService_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages_DebugService) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/messages/debug-entries/debug-service/state/service with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_DebugService_ServicePath) Lookup(t testing.TB) *oc.QualifiedE_Messages_DEBUG_SERVICE {
	t.Helper()
	goStruct := &oc.System_Messages_DebugService{}
	md, ok := oc.Lookup(t, n, "System_Messages_DebugService", goStruct, true, false)
	if ok {
		return convertSystem_Messages_DebugService_ServicePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/debug-entries/debug-service/state/service with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_DebugService_ServicePath) Get(t testing.TB) oc.E_Messages_DEBUG_SERVICE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/debug-entries/debug-service/state/service with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_DebugService_ServicePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Messages_DEBUG_SERVICE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Messages_DEBUG_SERVICE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_DebugService{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_DebugService", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_DebugService_ServicePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/debug-entries/debug-service/state/service with a ONCE subscription.
func (n *System_Messages_DebugService_ServicePathAny) Get(t testing.TB) []oc.E_Messages_DEBUG_SERVICE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Messages_DEBUG_SERVICE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/debug-entries/debug-service/state/service with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_DebugService_ServicePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Messages_DEBUG_SERVICE {
	t.Helper()
	c := &oc.CollectionE_Messages_DEBUG_SERVICE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Messages_DEBUG_SERVICE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Messages_DebugService_ServicePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Messages_DEBUG_SERVICE) bool) *oc.E_Messages_DEBUG_SERVICEWatcher {
	t.Helper()
	w := &oc.E_Messages_DEBUG_SERVICEWatcher{}
	gs := &oc.System_Messages_DebugService{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages_DebugService", gs, queryPath, true, false)
		return convertSystem_Messages_DebugService_ServicePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Messages_DEBUG_SERVICE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/debug-entries/debug-service/state/service with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_DebugService_ServicePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Messages_DEBUG_SERVICE) bool) *oc.E_Messages_DEBUG_SERVICEWatcher {
	t.Helper()
	return watch_System_Messages_DebugService_ServicePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages/debug-entries/debug-service/state/service with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Messages_DebugService_ServicePath) Await(t testing.TB, timeout time.Duration, val oc.E_Messages_DEBUG_SERVICE) *oc.QualifiedE_Messages_DEBUG_SERVICE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Messages_DEBUG_SERVICE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages/debug-entries/debug-service/state/service failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages/debug-entries/debug-service/state/service to the batch object.
func (n *System_Messages_DebugService_ServicePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/debug-entries/debug-service/state/service with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_DebugService_ServicePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Messages_DEBUG_SERVICE {
	t.Helper()
	c := &oc.CollectionE_Messages_DEBUG_SERVICE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Messages_DEBUG_SERVICE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/debug-entries/debug-service/state/service with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_DebugService_ServicePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Messages_DEBUG_SERVICE) bool) *oc.E_Messages_DEBUG_SERVICEWatcher {
	t.Helper()
	return watch_System_Messages_DebugService_ServicePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages/debug-entries/debug-service/state/service to the batch object.
func (n *System_Messages_DebugService_ServicePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Messages_DebugService_ServicePath extracts the value of the leaf Service from its parent oc.System_Messages_DebugService
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Messages_DEBUG_SERVICE.
func convertSystem_Messages_DebugService_ServicePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages_DebugService) *oc.QualifiedE_Messages_DEBUG_SERVICE {
	t.Helper()
	qv := &oc.QualifiedE_Messages_DEBUG_SERVICE{
		Metadata: md,
	}
	val := parent.Service
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/messages/state/message with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_MessagePath) Lookup(t testing.TB) *oc.QualifiedSystem_Messages_Message {
	t.Helper()
	goStruct := &oc.System_Messages_Message{}
	md, ok := oc.Lookup(t, n, "System_Messages_Message", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Messages_Message{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/state/message with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_MessagePath) Get(t testing.TB) *oc.System_Messages_Message {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/state/message with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_MessagePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Messages_Message {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Messages_Message
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_Message{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_Message", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Messages_Message{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/state/message with a ONCE subscription.
func (n *System_Messages_MessagePathAny) Get(t testing.TB) []*oc.System_Messages_Message {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Messages_Message
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_MessagePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Messages_Message {
	t.Helper()
	c := &oc.CollectionSystem_Messages_Message{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Messages_Message) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Messages_Message{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Messages_Message)))
		return false
	})
	return c
}

func watch_System_Messages_MessagePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Messages_Message) bool) *oc.System_Messages_MessageWatcher {
	t.Helper()
	w := &oc.System_Messages_MessageWatcher{}
	gs := &oc.System_Messages_Message{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages_Message", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Messages_Message{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Messages_Message)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_MessagePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Messages_Message) bool) *oc.System_Messages_MessageWatcher {
	t.Helper()
	return watch_System_Messages_MessagePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages/state/message with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Messages_MessagePath) Await(t testing.TB, timeout time.Duration, val *oc.System_Messages_Message) *oc.QualifiedSystem_Messages_Message {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Messages_Message) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages/state/message failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages/state/message to the batch object.
func (n *System_Messages_MessagePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_MessagePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Messages_Message {
	t.Helper()
	c := &oc.CollectionSystem_Messages_Message{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Messages_Message) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_MessagePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Messages_Message) bool) *oc.System_Messages_MessageWatcher {
	t.Helper()
	return watch_System_Messages_MessagePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages/state/message to the batch object.
func (n *System_Messages_MessagePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/messages/state/message/app-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_Message_AppNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Messages_Message{}
	md, ok := oc.Lookup(t, n, "System_Messages_Message", goStruct, true, false)
	if ok {
		return convertSystem_Messages_Message_AppNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/state/message/app-name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_Message_AppNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/state/message/app-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_Message_AppNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_Message{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_Message", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_Message_AppNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/state/message/app-name with a ONCE subscription.
func (n *System_Messages_Message_AppNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message/app-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_Message_AppNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Messages_Message_AppNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Messages_Message{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages_Message", gs, queryPath, true, false)
		return convertSystem_Messages_Message_AppNamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message/app-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_Message_AppNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Messages_Message_AppNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages/state/message/app-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Messages_Message_AppNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages/state/message/app-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages/state/message/app-name to the batch object.
func (n *System_Messages_Message_AppNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message/app-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_Message_AppNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message/app-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_Message_AppNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Messages_Message_AppNamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages/state/message/app-name to the batch object.
func (n *System_Messages_Message_AppNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Messages_Message_AppNamePath extracts the value of the leaf AppName from its parent oc.System_Messages_Message
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Messages_Message_AppNamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages_Message) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AppName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/messages/state/message/msg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_Message_MsgPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Messages_Message{}
	md, ok := oc.Lookup(t, n, "System_Messages_Message", goStruct, true, false)
	if ok {
		return convertSystem_Messages_Message_MsgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/state/message/msg with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_Message_MsgPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/state/message/msg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_Message_MsgPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_Message{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_Message", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_Message_MsgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/state/message/msg with a ONCE subscription.
func (n *System_Messages_Message_MsgPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message/msg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_Message_MsgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Messages_Message_MsgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Messages_Message{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages_Message", gs, queryPath, true, false)
		return convertSystem_Messages_Message_MsgPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message/msg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_Message_MsgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Messages_Message_MsgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages/state/message/msg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Messages_Message_MsgPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages/state/message/msg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages/state/message/msg to the batch object.
func (n *System_Messages_Message_MsgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message/msg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_Message_MsgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message/msg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_Message_MsgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Messages_Message_MsgPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages/state/message/msg to the batch object.
func (n *System_Messages_Message_MsgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Messages_Message_MsgPath extracts the value of the leaf Msg from its parent oc.System_Messages_Message
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Messages_Message_MsgPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages_Message) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Msg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/messages/state/message/msgid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_Message_MsgidPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Messages_Message{}
	md, ok := oc.Lookup(t, n, "System_Messages_Message", goStruct, true, false)
	if ok {
		return convertSystem_Messages_Message_MsgidPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/state/message/msgid with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_Message_MsgidPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/state/message/msgid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_Message_MsgidPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_Message{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_Message", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_Message_MsgidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/state/message/msgid with a ONCE subscription.
func (n *System_Messages_Message_MsgidPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message/msgid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_Message_MsgidPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Messages_Message_MsgidPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Messages_Message{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages_Message", gs, queryPath, true, false)
		return convertSystem_Messages_Message_MsgidPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message/msgid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_Message_MsgidPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Messages_Message_MsgidPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages/state/message/msgid with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Messages_Message_MsgidPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages/state/message/msgid failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages/state/message/msgid to the batch object.
func (n *System_Messages_Message_MsgidPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message/msgid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_Message_MsgidPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message/msgid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_Message_MsgidPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Messages_Message_MsgidPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages/state/message/msgid to the batch object.
func (n *System_Messages_Message_MsgidPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Messages_Message_MsgidPath extracts the value of the leaf Msgid from its parent oc.System_Messages_Message
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Messages_Message_MsgidPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages_Message) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Msgid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/messages/state/message/priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_Message_PriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.System_Messages_Message{}
	md, ok := oc.Lookup(t, n, "System_Messages_Message", goStruct, true, false)
	if ok {
		return convertSystem_Messages_Message_PriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/state/message/priority with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_Message_PriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/state/message/priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_Message_PriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_Message{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_Message", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_Message_PriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/state/message/priority with a ONCE subscription.
func (n *System_Messages_Message_PriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message/priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_Message_PriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Messages_Message_PriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.System_Messages_Message{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages_Message", gs, queryPath, true, false)
		return convertSystem_Messages_Message_PriorityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message/priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_Message_PriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Messages_Message_PriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages/state/message/priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Messages_Message_PriorityPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages/state/message/priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages/state/message/priority to the batch object.
func (n *System_Messages_Message_PriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message/priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_Message_PriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message/priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_Message_PriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_System_Messages_Message_PriorityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages/state/message/priority to the batch object.
func (n *System_Messages_Message_PriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Messages_Message_PriorityPath extracts the value of the leaf Priority from its parent oc.System_Messages_Message
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertSystem_Messages_Message_PriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages_Message) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Priority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/messages/state/message/procid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_Message_ProcidPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Messages_Message{}
	md, ok := oc.Lookup(t, n, "System_Messages_Message", goStruct, true, false)
	if ok {
		return convertSystem_Messages_Message_ProcidPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/state/message/procid with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_Message_ProcidPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/state/message/procid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_Message_ProcidPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_Message{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_Message", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_Message_ProcidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/state/message/procid with a ONCE subscription.
func (n *System_Messages_Message_ProcidPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message/procid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_Message_ProcidPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Messages_Message_ProcidPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Messages_Message{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages_Message", gs, queryPath, true, false)
		return convertSystem_Messages_Message_ProcidPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message/procid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_Message_ProcidPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Messages_Message_ProcidPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages/state/message/procid with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Messages_Message_ProcidPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages/state/message/procid failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages/state/message/procid to the batch object.
func (n *System_Messages_Message_ProcidPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/message/procid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_Message_ProcidPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/message/procid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_Message_ProcidPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Messages_Message_ProcidPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages/state/message/procid to the batch object.
func (n *System_Messages_Message_ProcidPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Messages_Message_ProcidPath extracts the value of the leaf Procid from its parent oc.System_Messages_Message
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Messages_Message_ProcidPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages_Message) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Procid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/messages/state/severity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_SeverityPath) Lookup(t testing.TB) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	goStruct := &oc.System_Messages{}
	md, ok := oc.Lookup(t, n, "System_Messages", goStruct, true, false)
	if ok {
		return convertSystem_Messages_SeverityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/state/severity with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_SeverityPath) Get(t testing.TB) oc.E_SystemLogging_SyslogSeverity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/state/severity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_SeverityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemLogging_SyslogSeverity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_SeverityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/state/severity with a ONCE subscription.
func (n *System_Messages_SeverityPathAny) Get(t testing.TB) []oc.E_SystemLogging_SyslogSeverity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SystemLogging_SyslogSeverity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/severity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_SeverityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SystemLogging_SyslogSeverity {
	t.Helper()
	c := &oc.CollectionE_SystemLogging_SyslogSeverity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SystemLogging_SyslogSeverity) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Messages_SeverityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SyslogSeverity) bool) *oc.E_SystemLogging_SyslogSeverityWatcher {
	t.Helper()
	w := &oc.E_SystemLogging_SyslogSeverityWatcher{}
	gs := &oc.System_Messages{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Messages", gs, queryPath, true, false)
		return convertSystem_Messages_SeverityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SystemLogging_SyslogSeverity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_SeverityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SyslogSeverity) bool) *oc.E_SystemLogging_SyslogSeverityWatcher {
	t.Helper()
	return watch_System_Messages_SeverityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/messages/state/severity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Messages_SeverityPath) Await(t testing.TB, timeout time.Duration, val oc.E_SystemLogging_SyslogSeverity) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_SystemLogging_SyslogSeverity) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/messages/state/severity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/messages/state/severity to the batch object.
func (n *System_Messages_SeverityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/messages/state/severity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Messages_SeverityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SystemLogging_SyslogSeverity {
	t.Helper()
	c := &oc.CollectionE_SystemLogging_SyslogSeverity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SystemLogging_SyslogSeverity) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/messages/state/severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Messages_SeverityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SyslogSeverity) bool) *oc.E_SystemLogging_SyslogSeverityWatcher {
	t.Helper()
	return watch_System_Messages_SeverityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/messages/state/severity to the batch object.
func (n *System_Messages_SeverityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Messages_SeverityPath extracts the value of the leaf Severity from its parent oc.System_Messages
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemLogging_SyslogSeverity.
func convertSystem_Messages_SeverityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	qv := &oc.QualifiedE_SystemLogging_SyslogSeverity{
		Metadata: md,
	}
	val := parent.Severity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/state/motd-banner with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MotdBannerPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, true, false)
	if ok {
		return convertSystem_MotdBannerPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/state/motd-banner with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MotdBannerPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/state/motd-banner with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MotdBannerPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_MotdBannerPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/state/motd-banner with a ONCE subscription.
func (n *System_MotdBannerPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/motd-banner with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MotdBannerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_MotdBannerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System", gs, queryPath, true, false)
		return convertSystem_MotdBannerPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/motd-banner with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MotdBannerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_MotdBannerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/state/motd-banner with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_MotdBannerPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/state/motd-banner failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/state/motd-banner to the batch object.
func (n *System_MotdBannerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/motd-banner with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MotdBannerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/motd-banner with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MotdBannerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_MotdBannerPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/state/motd-banner to the batch object.
func (n *System_MotdBannerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_MotdBannerPath extracts the value of the leaf MotdBanner from its parent oc.System
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_MotdBannerPath(t testing.TB, md *genutil.Metadata, parent *oc.System) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MotdBanner
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/mount-points/mount-point with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MountPointPath) Lookup(t testing.TB) *oc.QualifiedSystem_MountPoint {
	t.Helper()
	goStruct := &oc.System_MountPoint{}
	md, ok := oc.Lookup(t, n, "System_MountPoint", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_MountPoint{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/mount-points/mount-point with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MountPointPath) Get(t testing.TB) *oc.System_MountPoint {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/mount-points/mount-point with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MountPointPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_MountPoint {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_MountPoint
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_MountPoint{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_MountPoint", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_MountPoint{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/mount-points/mount-point with a ONCE subscription.
func (n *System_MountPointPathAny) Get(t testing.TB) []*oc.System_MountPoint {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_MountPoint
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPointPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_MountPoint {
	t.Helper()
	c := &oc.CollectionSystem_MountPoint{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_MountPoint) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_MountPoint{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_MountPoint)))
		return false
	})
	return c
}

func watch_System_MountPointPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_MountPoint) bool) *oc.System_MountPointWatcher {
	t.Helper()
	w := &oc.System_MountPointWatcher{}
	gs := &oc.System_MountPoint{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_MountPoint", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_MountPoint{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_MountPoint)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPointPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_MountPoint) bool) *oc.System_MountPointWatcher {
	t.Helper()
	return watch_System_MountPointPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/mount-points/mount-point with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_MountPointPath) Await(t testing.TB, timeout time.Duration, val *oc.System_MountPoint) *oc.QualifiedSystem_MountPoint {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_MountPoint) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/mount-points/mount-point failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/mount-points/mount-point to the batch object.
func (n *System_MountPointPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPointPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_MountPoint {
	t.Helper()
	c := &oc.CollectionSystem_MountPoint{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_MountPoint) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPointPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_MountPoint) bool) *oc.System_MountPointWatcher {
	t.Helper()
	return watch_System_MountPointPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/mount-points/mount-point to the batch object.
func (n *System_MountPointPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/mount-points/mount-point/state/available with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MountPoint_AvailablePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_MountPoint{}
	md, ok := oc.Lookup(t, n, "System_MountPoint", goStruct, true, false)
	if ok {
		return convertSystem_MountPoint_AvailablePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/mount-points/mount-point/state/available with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MountPoint_AvailablePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/mount-points/mount-point/state/available with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MountPoint_AvailablePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_MountPoint{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_MountPoint", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_MountPoint_AvailablePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/mount-points/mount-point/state/available with a ONCE subscription.
func (n *System_MountPoint_AvailablePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point/state/available with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPoint_AvailablePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_MountPoint_AvailablePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_MountPoint{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_MountPoint", gs, queryPath, true, false)
		return convertSystem_MountPoint_AvailablePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point/state/available with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPoint_AvailablePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_MountPoint_AvailablePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/mount-points/mount-point/state/available with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_MountPoint_AvailablePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/mount-points/mount-point/state/available failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/mount-points/mount-point/state/available to the batch object.
func (n *System_MountPoint_AvailablePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point/state/available with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPoint_AvailablePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point/state/available with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPoint_AvailablePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_MountPoint_AvailablePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/mount-points/mount-point/state/available to the batch object.
func (n *System_MountPoint_AvailablePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_MountPoint_AvailablePath extracts the value of the leaf Available from its parent oc.System_MountPoint
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_MountPoint_AvailablePath(t testing.TB, md *genutil.Metadata, parent *oc.System_MountPoint) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Available
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/mount-points/mount-point/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MountPoint_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_MountPoint{}
	md, ok := oc.Lookup(t, n, "System_MountPoint", goStruct, true, false)
	if ok {
		return convertSystem_MountPoint_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/mount-points/mount-point/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MountPoint_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/mount-points/mount-point/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MountPoint_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_MountPoint{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_MountPoint", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_MountPoint_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/mount-points/mount-point/state/name with a ONCE subscription.
func (n *System_MountPoint_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPoint_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_MountPoint_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_MountPoint{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_MountPoint", gs, queryPath, true, false)
		return convertSystem_MountPoint_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPoint_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_MountPoint_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/mount-points/mount-point/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_MountPoint_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/mount-points/mount-point/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/mount-points/mount-point/state/name to the batch object.
func (n *System_MountPoint_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/mount-points/mount-point/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_MountPoint_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/mount-points/mount-point/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_MountPoint_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_MountPoint_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/mount-points/mount-point/state/name to the batch object.
func (n *System_MountPoint_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_MountPoint_NamePath extracts the value of the leaf Name from its parent oc.System_MountPoint
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_MountPoint_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_MountPoint) *oc.QualifiedString {
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
