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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_SystemPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem) bool) *oc.SystemWatcher {
	t.Helper()
	w := &oc.SystemWatcher{}
	structs := map[string]*oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *SystemPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem) bool) *oc.SystemWatcher {
	t.Helper()
	return watch_SystemPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_System_AaaPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa) bool) *oc.System_AaaWatcher {
	t.Helper()
	w := &oc.System_AaaWatcher{}
	structs := map[string]*oc.System_Aaa{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *System_AaaPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa) bool) *oc.System_AaaWatcher {
	t.Helper()
	return watch_System_AaaPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Accounting", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_Accounting{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_System_Aaa_AccountingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting) bool) *oc.System_Aaa_AccountingWatcher {
	t.Helper()
	w := &oc.System_Aaa_AccountingWatcher{}
	structs := map[string]*oc.System_Aaa_Accounting{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Accounting{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Accounting", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_Accounting{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *System_Aaa_AccountingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting) bool) *oc.System_Aaa_AccountingWatcher {
	t.Helper()
	return watch_System_Aaa_AccountingPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Accounting", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Accounting_AccountingMethodPath(t, md, gs)}, nil
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

func watch_System_Aaa_Accounting_AccountingMethodPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) bool) *oc.System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher {
	t.Helper()
	w := &oc.System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher{}
	structs := map[string]*oc.System_Aaa_Accounting{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Accounting{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Accounting", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Accounting_AccountingMethodPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *System_Aaa_Accounting_AccountingMethodPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) bool) *oc.System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_AccountingMethodPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Accounting_Event", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_Accounting_Event{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_System_Aaa_Accounting_EventPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting_Event) bool) *oc.System_Aaa_Accounting_EventWatcher {
	t.Helper()
	w := &oc.System_Aaa_Accounting_EventWatcher{}
	structs := map[string]*oc.System_Aaa_Accounting_Event{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Accounting_Event{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Accounting_Event", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_Accounting_Event{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *System_Aaa_Accounting_EventPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Accounting_Event) bool) *oc.System_Aaa_Accounting_EventWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_EventPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Accounting_Event", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Accounting_Event_EventTypePath(t, md, gs)}, nil
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

func watch_System_Aaa_Accounting_Event_EventTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) bool) *oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher {
	t.Helper()
	w := &oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher{}
	structs := map[string]*oc.System_Aaa_Accounting_Event{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Accounting_Event{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Accounting_Event", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Accounting_Event_EventTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *System_Aaa_Accounting_Event_EventTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) bool) *oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_Event_EventTypePathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Accounting_Event", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Accounting_Event_RecordPath(t, md, gs)}, nil
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

func watch_System_Aaa_Accounting_Event_RecordPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Event_Record) bool) *oc.E_Event_RecordWatcher {
	t.Helper()
	w := &oc.E_Event_RecordWatcher{}
	structs := map[string]*oc.System_Aaa_Accounting_Event{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Accounting_Event{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Accounting_Event", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Accounting_Event_RecordPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *System_Aaa_Accounting_Event_RecordPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Event_Record) bool) *oc.E_Event_RecordWatcher {
	t.Helper()
	return watch_System_Aaa_Accounting_Event_RecordPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_Authentication{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_System_Aaa_AuthenticationPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication) bool) *oc.System_Aaa_AuthenticationWatcher {
	t.Helper()
	w := &oc.System_Aaa_AuthenticationWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_Authentication{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *System_Aaa_AuthenticationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication) bool) *oc.System_Aaa_AuthenticationWatcher {
	t.Helper()
	return watch_System_Aaa_AuthenticationPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_AdminUser", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_Authentication_AdminUser{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_System_Aaa_Authentication_AdminUserPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_AdminUser) bool) *oc.System_Aaa_Authentication_AdminUserWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authentication_AdminUserWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_AdminUser{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_AdminUser{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_Authentication_AdminUser{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *System_Aaa_Authentication_AdminUserPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_AdminUser) bool) *oc.System_Aaa_Authentication_AdminUserWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUserPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_AdminUser", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t, md, gs)}, nil
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

func watch_System_Aaa_Authentication_AdminUser_AdminPasswordHashedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_AdminUser{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_AdminUser{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password-hashed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUser_AdminPasswordHashedPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_AdminUser", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath(t, md, gs)}, nil
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

func watch_System_Aaa_Authentication_AdminUser_AdminPasswordPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_AdminUser{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_AdminUser{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-password with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUser_AdminPasswordPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_AdminUser", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_AdminUser_AdminUsernamePath(t, md, gs)}, nil
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

func watch_System_Aaa_Authentication_AdminUser_AdminUsernamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_AdminUser{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_AdminUser{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_AdminUser_AdminUsernamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/admin-user/state/admin-username with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AdminUser_AdminUsernamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AdminUser_AdminUsernamePathAny(t, n, timeout, predicate)
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

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/state/authentication-method with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_AuthenticationMethodPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/state/authentication-method with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Get(t testing.TB) []oc.System_Aaa_Authentication_AuthenticationMethod_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/state/authentication-method with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_AuthenticationMethodPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_AuthenticationMethodPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/state/authentication-method with a ONCE subscription.
func (n *System_Aaa_Authentication_AuthenticationMethodPathAny) Get(t testing.TB) [][]oc.System_Aaa_Authentication_AuthenticationMethod_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.System_Aaa_Authentication_AuthenticationMethod_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/state/authentication-method with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_AuthenticationMethodPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) bool) *oc.System_Aaa_Authentication_AuthenticationMethod_UnionSliceWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authentication_AuthenticationMethod_UnionSliceWatcher{}
	gs := &oc.System_Aaa_Authentication{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_AuthenticationMethodPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/state/authentication-method with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) bool) *oc.System_Aaa_Authentication_AuthenticationMethod_UnionSliceWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AuthenticationMethodPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/state/authentication-method with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Await(t testing.TB, timeout time.Duration, val []oc.System_Aaa_Authentication_AuthenticationMethod_Union) *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/state/authentication-method failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/state/authentication-method to the batch object.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/state/authentication-method with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_AuthenticationMethodPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_AuthenticationMethodPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) bool) *oc.System_Aaa_Authentication_AuthenticationMethod_UnionSliceWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authentication_AuthenticationMethod_UnionSliceWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_AuthenticationMethodPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/state/authentication-method with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_AuthenticationMethodPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) bool) *oc.System_Aaa_Authentication_AuthenticationMethod_UnionSliceWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_AuthenticationMethodPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/state/authentication-method to the batch object.
func (n *System_Aaa_Authentication_AuthenticationMethodPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_AuthenticationMethodPath extracts the value of the leaf AuthenticationMethod from its parent oc.System_Aaa_Authentication
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice.
func convertSystem_Aaa_Authentication_AuthenticationMethodPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication) *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice{
		Metadata: md,
	}
	val := parent.AuthenticationMethod
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_UserPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authentication_User {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Authentication_User{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_UserPath) Get(t testing.TB) *oc.System_Aaa_Authentication_User {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_UserPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authentication_User {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authentication_User
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Authentication_User{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user with a ONCE subscription.
func (n *System_Aaa_Authentication_UserPathAny) Get(t testing.TB) []*oc.System_Aaa_Authentication_User {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Authentication_User
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_UserPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authentication_User {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authentication_User{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authentication_User) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_Authentication_User{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_Authentication_User)))
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_UserPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_User) bool) *oc.System_Aaa_Authentication_UserWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authentication_UserWatcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_Authentication_User{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authentication_User)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_UserPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_User) bool) *oc.System_Aaa_Authentication_UserWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_UserPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_UserPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_Authentication_User) *oc.QualifiedSystem_Aaa_Authentication_User {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Authentication_User) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user to the batch object.
func (n *System_Aaa_Authentication_UserPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_UserPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authentication_User {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authentication_User{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authentication_User) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_UserPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_User) bool) *oc.System_Aaa_Authentication_UserWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authentication_UserWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_Authentication_User{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authentication_User)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_UserPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_User) bool) *oc.System_Aaa_Authentication_UserWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_UserPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user to the batch object.
func (n *System_Aaa_Authentication_UserPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on with a ONCE subscription.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on to the batch object.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-created-on to the batch object.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath extracts the value of the leaf AuthorizedKeysListCreatedOn from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_Authentication_User_AuthorizedKeysListCreatedOnPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AuthorizedKeysListCreatedOn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_User_AuthorizedKeysListVersionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_AuthorizedKeysListVersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version with a ONCE subscription.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_AuthorizedKeysListVersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_User_AuthorizedKeysListVersionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_AuthorizedKeysListVersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version to the batch object.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_AuthorizedKeysListVersionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_User_AuthorizedKeysListVersionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_AuthorizedKeysListVersionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/authorized-keys-list-version to the batch object.
func (n *System_Aaa_Authentication_User_AuthorizedKeysListVersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_User_AuthorizedKeysListVersionPath extracts the value of the leaf AuthorizedKeysListVersion from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_User_AuthorizedKeysListVersionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AuthorizedKeysListVersion
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on with a ONCE subscription.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on to the batch object.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-created-on to the batch object.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath extracts the value of the leaf AuthorizedUsersListCreatedOn from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_Authentication_User_AuthorizedUsersListCreatedOnPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AuthorizedUsersListCreatedOn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_User_AuthorizedUsersListVersionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_AuthorizedUsersListVersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version with a ONCE subscription.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_AuthorizedUsersListVersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_User_AuthorizedUsersListVersionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_AuthorizedUsersListVersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version to the batch object.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_AuthorizedUsersListVersionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_User_AuthorizedUsersListVersionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_AuthorizedUsersListVersionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/authorized-users-list-version to the batch object.
func (n *System_Aaa_Authentication_User_AuthorizedUsersListVersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_User_AuthorizedUsersListVersionPath extracts the value of the leaf AuthorizedUsersListVersion from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_User_AuthorizedUsersListVersionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AuthorizedUsersListVersion
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/password-created-on with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_User_PasswordCreatedOnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/password-created-on with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/password-created-on with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_PasswordCreatedOnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/password-created-on with a ONCE subscription.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_PasswordCreatedOnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_User_PasswordCreatedOnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_PasswordCreatedOnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user/state/password-created-on with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user/state/password-created-on failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/password-created-on to the batch object.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_PasswordCreatedOnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_User_PasswordCreatedOnPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_PasswordCreatedOnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/password-created-on to the batch object.
func (n *System_Aaa_Authentication_User_PasswordCreatedOnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_User_PasswordCreatedOnPath extracts the value of the leaf PasswordCreatedOn from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_Aaa_Authentication_User_PasswordCreatedOnPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.PasswordCreatedOn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/password-hashed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_User_PasswordHashedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/password-hashed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/password-hashed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_PasswordHashedPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_PasswordHashedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/password-hashed with a ONCE subscription.
func (n *System_Aaa_Authentication_User_PasswordHashedPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-hashed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_PasswordHashedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_User_PasswordHashedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-hashed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_PasswordHashedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user/state/password-hashed with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user/state/password-hashed failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/password-hashed to the batch object.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-hashed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_PasswordHashedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_PasswordHashedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_User_PasswordHashedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-hashed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_PasswordHashedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_PasswordHashedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/password-hashed to the batch object.
func (n *System_Aaa_Authentication_User_PasswordHashedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_User_PasswordHashedPath extracts the value of the leaf PasswordHashed from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_User_PasswordHashedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.PasswordHashed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/password with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_PasswordPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_User_PasswordPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/password with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_PasswordPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/password with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_PasswordPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_PasswordPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/password with a ONCE subscription.
func (n *System_Aaa_Authentication_User_PasswordPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/password with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_PasswordPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_PasswordPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_User_PasswordPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/password with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_PasswordPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_PasswordPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user/state/password with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_User_PasswordPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user/state/password failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/password to the batch object.
func (n *System_Aaa_Authentication_User_PasswordPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/password with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_PasswordPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_PasswordPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_User_PasswordPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/password with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_PasswordPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_PasswordPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/password to the batch object.
func (n *System_Aaa_Authentication_User_PasswordPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_User_PasswordPath extracts the value of the leaf Password from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_User_PasswordPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Password
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/password-version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_PasswordVersionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_User_PasswordVersionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/password-version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_PasswordVersionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/password-version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_PasswordVersionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_PasswordVersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/password-version with a ONCE subscription.
func (n *System_Aaa_Authentication_User_PasswordVersionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_PasswordVersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_PasswordVersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_User_PasswordVersionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_PasswordVersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_PasswordVersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user/state/password-version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_User_PasswordVersionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user/state/password-version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/password-version to the batch object.
func (n *System_Aaa_Authentication_User_PasswordVersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_PasswordVersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_PasswordVersionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_User_PasswordVersionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/password-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_PasswordVersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_PasswordVersionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/password-version to the batch object.
func (n *System_Aaa_Authentication_User_PasswordVersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_User_PasswordVersionPath extracts the value of the leaf PasswordVersion from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_User_PasswordVersionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.PasswordVersion
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/role with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_RolePath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_User_RolePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/role with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_RolePath) Get(t testing.TB) oc.System_Aaa_Authentication_User_Role_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/role with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_RolePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authentication_User_Role_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authentication_User_Role_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_RolePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/role with a ONCE subscription.
func (n *System_Aaa_Authentication_User_RolePathAny) Get(t testing.TB) []oc.System_Aaa_Authentication_User_Role_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.System_Aaa_Authentication_User_Role_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/role with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_RolePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authentication_User_Role_Union {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authentication_User_Role_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_RolePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union) bool) *oc.System_Aaa_Authentication_User_Role_UnionWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authentication_User_Role_UnionWatcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_User_RolePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authentication_User_Role_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/role with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_RolePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union) bool) *oc.System_Aaa_Authentication_User_Role_UnionWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_RolePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user/state/role with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_User_RolePath) Await(t testing.TB, timeout time.Duration, val oc.System_Aaa_Authentication_User_Role_Union) *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user/state/role failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/role to the batch object.
func (n *System_Aaa_Authentication_User_RolePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/role with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_RolePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authentication_User_Role_Union {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authentication_User_Role_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_RolePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union) bool) *oc.System_Aaa_Authentication_User_Role_UnionWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authentication_User_Role_UnionWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_User_RolePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authentication_User_Role_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/role with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_RolePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union) bool) *oc.System_Aaa_Authentication_User_Role_UnionWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_RolePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/role to the batch object.
func (n *System_Aaa_Authentication_User_RolePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_User_RolePath extracts the value of the leaf Role from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union.
func convertSystem_Aaa_Authentication_User_RolePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union {
	t.Helper()
	qv := &oc.QualifiedSystem_Aaa_Authentication_User_Role_Union{
		Metadata: md,
	}
	val := parent.Role
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/username with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_UsernamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authentication_User_UsernamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/state/username with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_UsernamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/username with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_UsernamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_UsernamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/state/username with a ONCE subscription.
func (n *System_Aaa_Authentication_User_UsernamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/username with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_UsernamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_UsernamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authentication_User", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authentication_User_UsernamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/username with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_UsernamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_UsernamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authentication/users/user/state/username with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authentication_User_UsernamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authentication/users/user/state/username failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/username to the batch object.
func (n *System_Aaa_Authentication_User_UsernamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authentication/users/user/state/username with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authentication_User_UsernamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authentication_User_UsernamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Aaa_Authentication_User{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authentication_User{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authentication_User", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authentication_User_UsernamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authentication/users/user/state/username with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authentication_User_UsernamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Aaa_Authentication_User_UsernamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authentication/users/user/state/username to the batch object.
func (n *System_Aaa_Authentication_User_UsernamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authentication_User_UsernamePath extracts the value of the leaf Username from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_User_UsernamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Username
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authorization with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_AuthorizationPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authorization {
	t.Helper()
	goStruct := &oc.System_Aaa_Authorization{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authorization", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Authorization{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authorization with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_AuthorizationPath) Get(t testing.TB) *oc.System_Aaa_Authorization {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authorization with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_AuthorizationPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authorization {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authorization
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authorization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authorization", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Authorization{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authorization with a ONCE subscription.
func (n *System_Aaa_AuthorizationPathAny) Get(t testing.TB) []*oc.System_Aaa_Authorization {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Authorization
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authorization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_AuthorizationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authorization {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authorization{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authorization) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_Authorization{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_Authorization)))
		return false
	})
	return c
}

func watch_System_Aaa_AuthorizationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization) bool) *oc.System_Aaa_AuthorizationWatcher {
	t.Helper()
	w := &oc.System_Aaa_AuthorizationWatcher{}
	gs := &oc.System_Aaa_Authorization{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authorization", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_Authorization{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authorization)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authorization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_AuthorizationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization) bool) *oc.System_Aaa_AuthorizationWatcher {
	t.Helper()
	return watch_System_Aaa_AuthorizationPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authorization with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_AuthorizationPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_Authorization) *oc.QualifiedSystem_Aaa_Authorization {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Authorization) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authorization failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authorization to the batch object.
func (n *System_Aaa_AuthorizationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authorization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_AuthorizationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authorization {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authorization{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authorization) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_AuthorizationPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization) bool) *oc.System_Aaa_AuthorizationWatcher {
	t.Helper()
	w := &oc.System_Aaa_AuthorizationWatcher{}
	structs := map[string]*oc.System_Aaa_Authorization{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authorization{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authorization", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_Authorization{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authorization)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authorization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_AuthorizationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization) bool) *oc.System_Aaa_AuthorizationWatcher {
	t.Helper()
	return watch_System_Aaa_AuthorizationPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authorization to the batch object.
func (n *System_Aaa_AuthorizationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/authorization/state/authorization-method with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	t.Helper()
	goStruct := &oc.System_Aaa_Authorization{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authorization", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authorization_AuthorizationMethodPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authorization/state/authorization-method with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Get(t testing.TB) []oc.System_Aaa_Authorization_AuthorizationMethod_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authorization/state/authorization-method with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authorization_AuthorizationMethodPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authorization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authorization", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authorization_AuthorizationMethodPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authorization/state/authorization-method with a ONCE subscription.
func (n *System_Aaa_Authorization_AuthorizationMethodPathAny) Get(t testing.TB) [][]oc.System_Aaa_Authorization_AuthorizationMethod_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.System_Aaa_Authorization_AuthorizationMethod_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authorization/state/authorization-method with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authorization_AuthorizationMethodPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) bool) *oc.System_Aaa_Authorization_AuthorizationMethod_UnionSliceWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authorization_AuthorizationMethod_UnionSliceWatcher{}
	gs := &oc.System_Aaa_Authorization{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authorization", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authorization_AuthorizationMethodPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authorization/state/authorization-method with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) bool) *oc.System_Aaa_Authorization_AuthorizationMethod_UnionSliceWatcher {
	t.Helper()
	return watch_System_Aaa_Authorization_AuthorizationMethodPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authorization/state/authorization-method with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Await(t testing.TB, timeout time.Duration, val []oc.System_Aaa_Authorization_AuthorizationMethod_Union) *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authorization/state/authorization-method failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authorization/state/authorization-method to the batch object.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authorization/state/authorization-method with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authorization_AuthorizationMethodPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authorization_AuthorizationMethodPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) bool) *oc.System_Aaa_Authorization_AuthorizationMethod_UnionSliceWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authorization_AuthorizationMethod_UnionSliceWatcher{}
	structs := map[string]*oc.System_Aaa_Authorization{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authorization{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authorization", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authorization_AuthorizationMethodPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authorization/state/authorization-method with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authorization_AuthorizationMethodPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) bool) *oc.System_Aaa_Authorization_AuthorizationMethod_UnionSliceWatcher {
	t.Helper()
	return watch_System_Aaa_Authorization_AuthorizationMethodPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authorization/state/authorization-method to the batch object.
func (n *System_Aaa_Authorization_AuthorizationMethodPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authorization_AuthorizationMethodPath extracts the value of the leaf AuthorizationMethod from its parent oc.System_Aaa_Authorization
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice.
func convertSystem_Aaa_Authorization_AuthorizationMethodPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authorization) *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice{
		Metadata: md,
	}
	val := parent.AuthorizationMethod
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authorization/events/event with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authorization_EventPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authorization_Event {
	t.Helper()
	goStruct := &oc.System_Aaa_Authorization_Event{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authorization_Event", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Authorization_Event{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authorization/events/event with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authorization_EventPath) Get(t testing.TB) *oc.System_Aaa_Authorization_Event {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authorization/events/event with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authorization_EventPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authorization_Event {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authorization_Event
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authorization_Event{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authorization_Event", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Authorization_Event{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authorization/events/event with a ONCE subscription.
func (n *System_Aaa_Authorization_EventPathAny) Get(t testing.TB) []*oc.System_Aaa_Authorization_Event {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Authorization_Event
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authorization/events/event with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authorization_EventPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authorization_Event {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authorization_Event{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authorization_Event) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_Authorization_Event{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_Authorization_Event)))
		return false
	})
	return c
}

func watch_System_Aaa_Authorization_EventPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization_Event) bool) *oc.System_Aaa_Authorization_EventWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authorization_EventWatcher{}
	gs := &oc.System_Aaa_Authorization_Event{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authorization_Event", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_Authorization_Event{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authorization_Event)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authorization/events/event with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authorization_EventPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization_Event) bool) *oc.System_Aaa_Authorization_EventWatcher {
	t.Helper()
	return watch_System_Aaa_Authorization_EventPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authorization/events/event with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authorization_EventPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_Authorization_Event) *oc.QualifiedSystem_Aaa_Authorization_Event {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_Authorization_Event) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authorization/events/event failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authorization/events/event to the batch object.
func (n *System_Aaa_Authorization_EventPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authorization/events/event with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authorization_EventPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_Authorization_Event {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_Authorization_Event{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_Authorization_Event) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authorization_EventPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization_Event) bool) *oc.System_Aaa_Authorization_EventWatcher {
	t.Helper()
	w := &oc.System_Aaa_Authorization_EventWatcher{}
	structs := map[string]*oc.System_Aaa_Authorization_Event{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authorization_Event{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authorization_Event", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_Authorization_Event{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_Authorization_Event)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authorization/events/event with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authorization_EventPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_Authorization_Event) bool) *oc.System_Aaa_Authorization_EventWatcher {
	t.Helper()
	return watch_System_Aaa_Authorization_EventPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authorization/events/event to the batch object.
func (n *System_Aaa_Authorization_EventPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/aaa/authorization/events/event/state/event-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authorization_Event_EventTypePath) Lookup(t testing.TB) *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	goStruct := &oc.System_Aaa_Authorization_Event{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authorization_Event", goStruct, true, false)
	if ok {
		return convertSystem_Aaa_Authorization_Event_EventTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authorization/events/event/state/event-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authorization_Event_EventTypePath) Get(t testing.TB) oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authorization/events/event/state/event-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authorization_Event_EventTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authorization_Event{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authorization_Event", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authorization_Event_EventTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authorization/events/event/state/event-type with a ONCE subscription.
func (n *System_Aaa_Authorization_Event_EventTypePathAny) Get(t testing.TB) []oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authorization/events/event/state/event-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authorization_Event_EventTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	c := &oc.CollectionE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authorization_Event_EventTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) bool) *oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPEWatcher {
	t.Helper()
	w := &oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPEWatcher{}
	gs := &oc.System_Aaa_Authorization_Event{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_Authorization_Event", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Aaa_Authorization_Event_EventTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authorization/events/event/state/event-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authorization_Event_EventTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) bool) *oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPEWatcher {
	t.Helper()
	return watch_System_Aaa_Authorization_Event_EventTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/authorization/events/event/state/event-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_Authorization_Event_EventTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/authorization/events/event/state/event-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/authorization/events/event/state/event-type to the batch object.
func (n *System_Aaa_Authorization_Event_EventTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/authorization/events/event/state/event-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_Authorization_Event_EventTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	c := &oc.CollectionE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_Authorization_Event_EventTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) bool) *oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPEWatcher {
	t.Helper()
	w := &oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPEWatcher{}
	structs := map[string]*oc.System_Aaa_Authorization_Event{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_Authorization_Event{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_Authorization_Event", structs[pre], queryPath, true, false)
			qv := convertSystem_Aaa_Authorization_Event_EventTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/authorization/events/event/state/event-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_Authorization_Event_EventTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) bool) *oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPEWatcher {
	t.Helper()
	return watch_System_Aaa_Authorization_Event_EventTypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/authorization/events/event/state/event-type to the batch object.
func (n *System_Aaa_Authorization_Event_EventTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Aaa_Authorization_Event_EventTypePath extracts the value of the leaf EventType from its parent oc.System_Aaa_Authorization_Event
// and combines the update with an existing Metadata to return a *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE.
func convertSystem_Aaa_Authorization_Event_EventTypePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authorization_Event) *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE{
		Metadata: md,
	}
	val := parent.EventType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/server-groups/server-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_ServerGroupPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_ServerGroup {
	t.Helper()
	goStruct := &oc.System_Aaa_ServerGroup{}
	md, ok := oc.Lookup(t, n, "System_Aaa_ServerGroup", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Aaa_ServerGroup{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/server-groups/server-group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_ServerGroupPath) Get(t testing.TB) *oc.System_Aaa_ServerGroup {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/server-groups/server-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_ServerGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_ServerGroup {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_ServerGroup
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_ServerGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_ServerGroup", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/server-groups/server-group with a ONCE subscription.
func (n *System_Aaa_ServerGroupPathAny) Get(t testing.TB) []*oc.System_Aaa_ServerGroup {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_ServerGroup
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroupPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_ServerGroup {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_ServerGroup{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_ServerGroup) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Aaa_ServerGroup{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Aaa_ServerGroup)))
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroupPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup) bool) *oc.System_Aaa_ServerGroupWatcher {
	t.Helper()
	w := &oc.System_Aaa_ServerGroupWatcher{}
	gs := &oc.System_Aaa_ServerGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Aaa_ServerGroup", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Aaa_ServerGroup{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_ServerGroup)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroupPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup) bool) *oc.System_Aaa_ServerGroupWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroupPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/aaa/server-groups/server-group with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Aaa_ServerGroupPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Aaa_ServerGroup) *oc.QualifiedSystem_Aaa_ServerGroup {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Aaa_ServerGroup) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/aaa/server-groups/server-group failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group to the batch object.
func (n *System_Aaa_ServerGroupPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/aaa/server-groups/server-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Aaa_ServerGroupPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Aaa_ServerGroup {
	t.Helper()
	c := &oc.CollectionSystem_Aaa_ServerGroup{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Aaa_ServerGroup) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Aaa_ServerGroupPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup) bool) *oc.System_Aaa_ServerGroupWatcher {
	t.Helper()
	w := &oc.System_Aaa_ServerGroupWatcher{}
	structs := map[string]*oc.System_Aaa_ServerGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Aaa_ServerGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Aaa_ServerGroup", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Aaa_ServerGroup{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Aaa_ServerGroup)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/aaa/server-groups/server-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Aaa_ServerGroupPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Aaa_ServerGroup) bool) *oc.System_Aaa_ServerGroupWatcher {
	t.Helper()
	return watch_System_Aaa_ServerGroupPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/aaa/server-groups/server-group to the batch object.
func (n *System_Aaa_ServerGroupPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
