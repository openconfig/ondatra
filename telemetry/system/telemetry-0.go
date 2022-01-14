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

// Lookup fetches the value at /openconfig-system/system with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *SystemPath) Lookup(t testing.TB) *oc.QualifiedSystem {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *SystemPath) Get(t testing.TB) *oc.System {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *SystemPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system with a ONCE subscription.
func (n *SystemPathAny) Get(t testing.TB) []*oc.System {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *SystemPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem {
	t.Helper()
	c := &oc.CollectionSystem{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System)))
		return false
	})
	return c
}

func watch_SystemPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem) bool) *oc.SystemWatcher {
	t.Helper()
	w := &oc.SystemWatcher{}
	gs := &oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System", gs, queryPath, false, false)
		return (&oc.QualifiedSystem{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *SystemPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem) bool) *oc.SystemWatcher {
	t.Helper()
	return watch_SystemPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *SystemPath) Await(t testing.TB, timeout time.Duration, val *oc.System) *oc.QualifiedSystem {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system to the batch object.
func (n *SystemPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *SystemPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem {
	t.Helper()
	c := &oc.CollectionSystem{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *SystemPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem) bool) *oc.SystemWatcher {
	t.Helper()
	return watch_SystemPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system to the batch object.
func (n *SystemPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_AaaPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa {
	t.Helper()
	goStruct := &oc.System_Aaa{}
	md, ok := oc.Lookup(t, n, "System_Aaa", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_AaaPath) Get(t testing.TB) *oc.System_Aaa {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_AaaPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa with a ONCE subscription.
func (n *System_AaaPathAny) Get(t testing.TB) []*oc.System_Aaa {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_AaaPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa {
	t.Helper()
	c := &oc.CollectionSystem_Aaa{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa)))
		return false
	})
	return c
}

func watch_System_AaaPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa) bool) *oc.System_AaaWatcher {
	t.Helper()
	w := &oc.System_AaaWatcher{}
	gs := &oc.System_Aaa{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Aaa{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_AaaPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa) bool) *oc.System_AaaWatcher {
	t.Helper()
	return watch_System_AaaPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_AaaPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa) *oc.QualifiedSystem_Aaa {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa to the batch object.
func (n *System_AaaPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_AaaPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa {
	t.Helper()
	c := &oc.CollectionSystem_Aaa{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_AaaPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa) bool) *oc.System_AaaWatcher {
	t.Helper()
	return watch_System_AaaPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa to the batch object.
func (n *System_AaaPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/accounting with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_AccountingPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Accounting {
	t.Helper()
	goStruct := &oc.System_Aaa_Accounting{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Accounting", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Accounting{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/accounting with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_AccountingPath) Get(t testing.TB) *oc.System_Aaa_Accounting {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/accounting with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_AccountingPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Accounting {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Accounting
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Accounting{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Accounting", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Accounting{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/accounting with a ONCE subscription.
func (n *System_Aaa_AccountingPathAny) Get(t testing.TB) []*oc.System_Aaa_Accounting {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Accounting
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/accounting with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_AccountingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Accounting {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Accounting{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Accounting) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_Accounting{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_Accounting)))
		return false
	})
	return c
}

func watch_System_Aaa_AccountingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting) bool) *oc.System_Aaa_AccountingWatcher {
	t.Helper()
	w := &oc.System_Aaa_AccountingWatcher{}
	gs := &oc.System_Aaa_Accounting{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Accounting", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Aaa_Accounting{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Accounting)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/accounting with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_AccountingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting) bool) *oc.System_Aaa_AccountingWatcher {
	t.Helper()
	return watch_System_Aaa_AccountingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/accounting with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_AccountingPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_Accounting) *oc.QualifiedSystem_Aaa_Accounting {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Accounting) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/accounting failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/accounting to the batch object.
func (n *System_Aaa_AccountingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/accounting with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_AccountingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Accounting {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Accounting{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Accounting) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/accounting with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_AccountingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting) bool) *oc.System_Aaa_AccountingWatcher {
	t.Helper()
	return watch_System_Aaa_AccountingPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/accounting to the batch object.
func (n *System_Aaa_AccountingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/accounting/state/accounting-method with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Accounting_AccountingMethodPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	t.Helper()
	goStruct := &oc.System_Aaa_Accounting{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Accounting", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Accounting_AccountingMethodPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/accounting/state/accounting-method with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Accounting_AccountingMethodPath) Get(t testing.TB) []oc.System_Aaa_Accounting_AccountingMethod_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/accounting/state/accounting-method with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Accounting_AccountingMethodPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Accounting{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Accounting", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Accounting_AccountingMethodPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/accounting/state/accounting-method with a ONCE subscription.
func (n *System_Aaa_Accounting_AccountingMethodPathAny) Get(t testing.TB) [][]oc.System_Aaa_Accounting_AccountingMethod_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.System_Aaa_Accounting_AccountingMethod_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/accounting/state/accounting-method with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Accounting_AccountingMethodPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Accounting_AccountingMethod_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Accounting_AccountingMethodPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) bool) *oc.System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher {
	t.Helper()
	w := &oc.System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher{}
	gs := &oc.System_Aaa_Accounting{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Accounting", gs, queryPath, true, false)
		return convertSystem_Aaa_Accounting_AccountingMethodPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/accounting/state/accounting-method with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Accounting_AccountingMethodPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) bool) *oc.System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_AccountingMethodPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/accounting/state/accounting-method with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Accounting_AccountingMethodPath) Await(t testing.TB, timeout time.Duration, val []oc.System_Aaa_Accounting_AccountingMethod_Union) *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/accounting/state/accounting-method failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/accounting/state/accounting-method to the batch object.
func (n *System_Aaa_Accounting_AccountingMethodPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/accounting/state/accounting-method with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Accounting_AccountingMethodPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Accounting_AccountingMethod_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/accounting/state/accounting-method with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Accounting_AccountingMethodPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) bool) *oc.System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_AccountingMethodPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/accounting/state/accounting-method to the batch object.
func (n *System_Aaa_Accounting_AccountingMethodPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Accounting_AccountingMethodPath extracts the value of the leaf AccountingMethod from its parent oc.System_Aaa_Accounting
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice.
func convertSystem_Aaa_Accounting_AccountingMethodPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Accounting) *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice{
		Metadata: md,
	}
	val := parent.AccountingMethod
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/accounting/events/event with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Accounting_EventPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Accounting_Event {
	t.Helper()
	goStruct := &oc.System_Aaa_Accounting_Event{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Accounting_Event", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Accounting_Event{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/accounting/events/event with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Accounting_EventPath) Get(t testing.TB) *oc.System_Aaa_Accounting_Event {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/accounting/events/event with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Accounting_EventPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Accounting_Event {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Accounting_Event
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Accounting_Event{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Accounting_Event", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Accounting_Event{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/accounting/events/event with a ONCE subscription.
func (n *System_Aaa_Accounting_EventPathAny) Get(t testing.TB) []*oc.System_Aaa_Accounting_Event {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Accounting_Event
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/accounting/events/event with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Accounting_EventPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Accounting_Event {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Accounting_Event{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Accounting_Event) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_Accounting_Event{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_Accounting_Event)))
		return false
	})
	return c
}

func watch_System_Aaa_Accounting_EventPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting_Event) bool) *oc.System_Aaa_Accounting_EventWatcher {
	t.Helper()
	w := &oc.System_Aaa_Accounting_EventWatcher{}
	gs := &oc.System_Aaa_Accounting_Event{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Accounting_Event", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Aaa_Accounting_Event{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Accounting_Event)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/accounting/events/event with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Accounting_EventPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting_Event) bool) *oc.System_Aaa_Accounting_EventWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_EventPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/accounting/events/event with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Accounting_EventPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_Accounting_Event) *oc.QualifiedSystem_Aaa_Accounting_Event {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Accounting_Event) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/accounting/events/event failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/accounting/events/event to the batch object.
func (n *System_Aaa_Accounting_EventPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/accounting/events/event with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Accounting_EventPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Accounting_Event {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Accounting_Event{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Accounting_Event) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/accounting/events/event with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Accounting_EventPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting_Event) bool) *oc.System_Aaa_Accounting_EventWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_EventPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/accounting/events/event to the batch object.
func (n *System_Aaa_Accounting_EventPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/accounting/events/event/state/event-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Accounting_Event_EventTypePath) Lookup(t testing.TB) *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	goStruct := &oc.System_Aaa_Accounting_Event{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Accounting_Event", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Accounting_Event_EventTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/accounting/events/event/state/event-type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Accounting_Event_EventTypePath) Get(t testing.TB) oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/accounting/events/event/state/event-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Accounting_Event_EventTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Accounting_Event{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Accounting_Event", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Accounting_Event_EventTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/accounting/events/event/state/event-type with a ONCE subscription.
func (n *System_Aaa_Accounting_Event_EventTypePathAny) Get(t testing.TB) []oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/accounting/events/event/state/event-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Accounting_Event_EventTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	c := &oc.CollectionE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Accounting_Event_EventTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) bool) *oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher {
	t.Helper()
	w := &oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher{}
	gs := &oc.System_Aaa_Accounting_Event{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Accounting_Event", gs, queryPath, true, false)
		return convertSystem_Aaa_Accounting_Event_EventTypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/accounting/events/event/state/event-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Accounting_Event_EventTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) bool) *oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_Event_EventTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/accounting/events/event/state/event-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Accounting_Event_EventTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/accounting/events/event/state/event-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/accounting/events/event/state/event-type to the batch object.
func (n *System_Aaa_Accounting_Event_EventTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/accounting/events/event/state/event-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Accounting_Event_EventTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	c := &oc.CollectionE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/accounting/events/event/state/event-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Accounting_Event_EventTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) bool) *oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_Event_EventTypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/accounting/events/event/state/event-type to the batch object.
func (n *System_Aaa_Accounting_Event_EventTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Accounting_Event_EventTypePath extracts the value of the leaf EventType from its parent oc.System_Aaa_Accounting_Event
// and combines the update with an existing Metadata to return a *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE.
func convertSystem_Aaa_Accounting_Event_EventTypePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Accounting_Event) *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE{
		Metadata: md,
	}
	val := parent.EventType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/accounting/events/event/state/record with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Accounting_Event_RecordPath) Lookup(t testing.TB) *oc.QualifiedE_Event_Record {
	t.Helper()
	goStruct := &oc.System_Aaa_Accounting_Event{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Accounting_Event", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Accounting_Event_RecordPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/accounting/events/event/state/record with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Accounting_Event_RecordPath) Get(t testing.TB) oc.E_Event_Record {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/accounting/events/event/state/record with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Accounting_Event_RecordPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Event_Record {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Event_Record
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Accounting_Event{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Accounting_Event", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Accounting_Event_RecordPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/accounting/events/event/state/record with a ONCE subscription.
func (n *System_Aaa_Accounting_Event_RecordPathAny) Get(t testing.TB) []oc.E_Event_Record {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Event_Record
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/accounting/events/event/state/record with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Accounting_Event_RecordPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Event_Record {
	t.Helper()
	c := &oc.CollectionE_Event_Record{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Event_Record) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Accounting_Event_RecordPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Event_Record) bool) *oc.E_Event_RecordWatcher {
	t.Helper()
	w := &oc.E_Event_RecordWatcher{}
	gs := &oc.System_Aaa_Accounting_Event{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Accounting_Event", gs, queryPath, true, false)
		return convertSystem_Aaa_Accounting_Event_RecordPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Event_Record)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/accounting/events/event/state/record with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Accounting_Event_RecordPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Event_Record) bool) *oc.E_Event_RecordWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_Event_RecordPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/accounting/events/event/state/record with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Accounting_Event_RecordPath) Await(t testing.TB, timeout time.Duration, val oc.E_Event_Record) *oc.QualifiedE_Event_Record {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Event_Record) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/accounting/events/event/state/record failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/accounting/events/event/state/record to the batch object.
func (n *System_Aaa_Accounting_Event_RecordPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/accounting/events/event/state/record with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Accounting_Event_RecordPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Event_Record {
	t.Helper()
	c := &oc.CollectionE_Event_Record{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Event_Record) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/accounting/events/event/state/record with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Accounting_Event_RecordPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Event_Record) bool) *oc.E_Event_RecordWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_Event_RecordPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/accounting/events/event/state/record to the batch object.
func (n *System_Aaa_Accounting_Event_RecordPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Accounting_Event_RecordPath extracts the value of the leaf Record from its parent oc.System_Aaa_Accounting_Event
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Event_Record.
func convertSystem_Aaa_Accounting_Event_RecordPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Accounting_Event) *oc.QualifiedE_Event_Record {
	t.Helper()
	qv := &oc.QualifiedE_Event_Record{
		Metadata: md,
	}
	val := parent.Record
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_AuthenticationPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authentication {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Authentication{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_AuthenticationPath) Get(t testing.TB) *oc.System_Aaa_Authentication {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_AuthenticationPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authentication {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authentication
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Authentication{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication with a ONCE subscription.
func (n *System_Aaa_AuthenticationPathAny) Get(t testing.TB) []*oc.System_Aaa_Authentication {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Authentication
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_AuthenticationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authentication {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authentication{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authentication) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_Authentication{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_Authentication)))
		return false
	})
	return c
}

func watch_System_Aaa_AuthenticationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication) bool) *oc.System_Aaa_AuthenticationWatcher {
	t.Helper()
	w := &oc.System_Aaa_AuthenticationWatcher{}
	gs := &oc.System_Aaa_Authentication{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Aaa_Authentication{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authentication)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_AuthenticationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication) bool) *oc.System_Aaa_AuthenticationWatcher {
	t.Helper()
	return watch_System_Aaa_AuthenticationPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_AuthenticationPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_Authentication) *oc.QualifiedSystem_Aaa_Authentication {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Authentication) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication to the batch object.
func (n *System_Aaa_AuthenticationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_AuthenticationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authentication {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authentication{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authentication) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_AuthenticationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication) bool) *oc.System_Aaa_AuthenticationWatcher {
	t.Helper()
	return watch_System_Aaa_AuthenticationPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication to the batch object.
func (n *System_Aaa_AuthenticationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/admin-user with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_AdminUserPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authentication_AdminUser {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_AdminUser{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_AdminUser", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Authentication_AdminUser{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/admin-user with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_AdminUserPath) Get(t testing.TB) *oc.System_Aaa_Authentication_AdminUser {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/admin-user with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_AdminUserPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authentication_AdminUser {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authentication_AdminUser
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_AdminUser{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Authentication_AdminUser{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/admin-user with a ONCE subscription.
func (n *System_Aaa_Authentication_AdminUserPathAny) Get(t testing.TB) []*oc.System_Aaa_Authentication_AdminUser {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Authentication_AdminUser
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/admin-user with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_AdminUserPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authentication_AdminUser {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authentication_AdminUser{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authentication_AdminUser) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_Authentication_AdminUser{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_Authentication_AdminUser)))
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_AdminUserPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_AdminUser) bool) *oc.System_Aaa_Authentication_AdminUserWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authentication_AdminUserWatcher{}
	gs := &oc.System_Aaa_Authentication_AdminUser{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_AdminUser", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_Aaa_Authentication_AdminUser{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authentication_AdminUser)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUserPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_AdminUser) bool) *oc.System_Aaa_Authentication_AdminUserWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUserPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/admin-user with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_AdminUserPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_Authentication_AdminUser) *oc.QualifiedSystem_Aaa_Authentication_AdminUser {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Authentication_AdminUser) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/admin-user failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/admin-user to the batch object.
func (n *System_Aaa_Authentication_AdminUserPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/admin-user with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_AdminUserPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authentication_AdminUser {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authentication_AdminUser{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authentication_AdminUser) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUserPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_AdminUser) bool) *oc.System_Aaa_Authentication_AdminUserWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUserPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/admin-user to the batch object.
func (n *System_Aaa_Authentication_AdminUserPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_AdminUser{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_AdminUser", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_AdminUser{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed with a ONCE subscription.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_Authentication_AdminUser{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_AdminUser", gs, queryPath, true, false)
		return convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed to the batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed to the batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath extracts the value of the leaf AdminPasswordHashed from its parent oc.System_Aaa_Authentication_AdminUser
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_AdminUser) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AdminPasswordHashed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_AdminUser{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_AdminUser", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_AdminUser{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password with a ONCE subscription.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_AdminUser_AdminPasswordPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_Authentication_AdminUser{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_AdminUser", gs, queryPath, true, false)
		return convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUser_AdminPasswordPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/admin-user/state/admin-password to the batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUser_AdminPasswordPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/admin-user/state/admin-password to the batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath extracts the value of the leaf AdminPassword from its parent oc.System_Aaa_Authentication_AdminUser
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_AdminUser) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AdminPassword
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_AdminUser{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_AdminUser", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_AdminUser_AdminUsernamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_AdminUser{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_AdminUser_AdminUsernamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username with a ONCE subscription.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_AdminUser_AdminUsernamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_Authentication_AdminUser{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_AdminUser", gs, queryPath, true, false)
		return convertSystem_Aaa_Authentication_AdminUser_AdminUsernamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUser_AdminUsernamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/admin-user/state/admin-username to the batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUser_AdminUsernamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/admin-user/state/admin-username to the batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_AdminUser_AdminUsernamePath extracts the value of the leaf AdminUsername from its parent oc.System_Aaa_Authentication_AdminUser
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_AdminUser_AdminUsernamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_AdminUser) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AdminUsername
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
