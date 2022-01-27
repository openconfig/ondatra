package acl

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

// Lookup fetches the value at /openconfig-acl/acl with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *AclPath) Lookup(t testing.TB) *oc.QualifiedAcl {
	t.Helper()
	goStruct := &oc.Acl{}
	md, ok := oc.Lookup(t, n, "Acl", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *AclPath) Get(t testing.TB) *oc.Acl {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *AclPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl with a ONCE subscription.
func (n *AclPathAny) Get(t testing.TB) []*oc.Acl {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *AclPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl {
	t.Helper()
	c := &oc.CollectionAcl{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl)))
		return false
	})
	return c
}

func watch_AclPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl) bool) *oc.AclWatcher {
	t.Helper()
	w := &oc.AclWatcher{}
	gs := &oc.Acl{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl", gs, queryPath, false, false)
		return (&oc.QualifiedAcl{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *AclPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl) bool) *oc.AclWatcher {
	t.Helper()
	return watch_AclPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *AclPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl) *oc.QualifiedAcl {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl to the batch object.
func (n *AclPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *AclPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl {
	t.Helper()
	c := &oc.CollectionAcl{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *AclPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl) bool) *oc.AclWatcher {
	t.Helper()
	return watch_AclPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl to the batch object.
func (n *AclPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSetPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet {
	t.Helper()
	goStruct := &oc.Acl_AclSet{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_AclSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSetPath) Get(t testing.TB) *oc.Acl_AclSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSetPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set with a ONCE subscription.
func (n *Acl_AclSetPathAny) Get(t testing.TB) []*oc.Acl_AclSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet {
	t.Helper()
	c := &oc.CollectionAcl_AclSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_AclSet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_AclSet)))
		return false
	})
	return c
}

func watch_Acl_AclSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet) bool) *oc.Acl_AclSetWatcher {
	t.Helper()
	w := &oc.Acl_AclSetWatcher{}
	gs := &oc.Acl_AclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_AclSet{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet) bool) *oc.Acl_AclSetWatcher {
	t.Helper()
	return watch_Acl_AclSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSetPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_AclSet) *oc.QualifiedAcl_AclSet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set to the batch object.
func (n *Acl_AclSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet {
	t.Helper()
	c := &oc.CollectionAcl_AclSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet) bool) *oc.Acl_AclSetWatcher {
	t.Helper()
	return watch_Acl_AclSetPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set to the batch object.
func (n *Acl_AclSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntryPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntryPath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntryPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a ONCE subscription.
func (n *Acl_AclSet_AclEntryPathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntryPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_AclSet_AclEntry{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_AclSet_AclEntry)))
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntryPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry) bool) *oc.Acl_AclSet_AclEntryWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntryWatcher{}
	gs := &oc.Acl_AclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_AclSet_AclEntry{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntryPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry) bool) *oc.Acl_AclSet_AclEntryWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntryPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntryPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_AclSet_AclEntry) *oc.QualifiedAcl_AclSet_AclEntry {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry to the batch object.
func (n *Acl_AclSet_AclEntryPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntryPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntryPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry) bool) *oc.Acl_AclSet_AclEntryWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntryPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry to the batch object.
func (n *Acl_AclSet_AclEntryPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_ActionsPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Actions {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Actions", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_Actions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_ActionsPath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_Actions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_ActionsPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Actions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Actions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Actions", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Actions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_ActionsPathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_Actions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_Actions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_ActionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Actions {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Actions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Actions) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_AclSet_AclEntry_Actions{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_AclSet_AclEntry_Actions)))
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_ActionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Actions) bool) *oc.Acl_AclSet_AclEntry_ActionsWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_ActionsWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Actions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Actions", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_AclSet_AclEntry_Actions{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Actions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_ActionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Actions) bool) *oc.Acl_AclSet_AclEntry_ActionsWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_ActionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_ActionsPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_AclSet_AclEntry_Actions) *oc.QualifiedAcl_AclSet_AclEntry_Actions {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_Actions) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions to the batch object.
func (n *Acl_AclSet_AclEntry_ActionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_ActionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Actions {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Actions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Actions) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_ActionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Actions) bool) *oc.Acl_AclSet_AclEntry_ActionsWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_ActionsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions to the batch object.
func (n *Acl_AclSet_AclEntry_ActionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Lookup(t testing.TB) *oc.QualifiedE_Acl_FORWARDING_ACTION {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Actions", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Actions_ForwardingActionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Get(t testing.TB) oc.E_Acl_FORWARDING_ACTION {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_FORWARDING_ACTION {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_FORWARDING_ACTION
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Actions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Actions_ForwardingActionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPathAny) Get(t testing.TB) []oc.E_Acl_FORWARDING_ACTION {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_FORWARDING_ACTION
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_FORWARDING_ACTION {
	t.Helper()
	c := &oc.CollectionE_Acl_FORWARDING_ACTION{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_FORWARDING_ACTION) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Actions_ForwardingActionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Acl_FORWARDING_ACTION) bool) *oc.E_Acl_FORWARDING_ACTIONWatcher {
	t.Helper()
	w := &oc.E_Acl_FORWARDING_ACTIONWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Actions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Actions", gs, queryPath, true, false)
		return convertAcl_AclSet_AclEntry_Actions_ForwardingActionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Acl_FORWARDING_ACTION)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_FORWARDING_ACTION) bool) *oc.E_Acl_FORWARDING_ACTIONWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Actions_ForwardingActionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Await(t testing.TB, timeout time.Duration, val oc.E_Acl_FORWARDING_ACTION) *oc.QualifiedE_Acl_FORWARDING_ACTION {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Acl_FORWARDING_ACTION) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action to the batch object.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_FORWARDING_ACTION {
	t.Helper()
	c := &oc.CollectionE_Acl_FORWARDING_ACTION{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_FORWARDING_ACTION) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_FORWARDING_ACTION) bool) *oc.E_Acl_FORWARDING_ACTIONWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Actions_ForwardingActionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/forwarding-action to the batch object.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Actions_ForwardingActionPath extracts the value of the leaf ForwardingAction from its parent oc.Acl_AclSet_AclEntry_Actions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_FORWARDING_ACTION.
func convertAcl_AclSet_AclEntry_Actions_ForwardingActionPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Actions) *oc.QualifiedE_Acl_FORWARDING_ACTION {
	t.Helper()
	qv := &oc.QualifiedE_Acl_FORWARDING_ACTION{
		Metadata: md,
	}
	val := parent.ForwardingAction
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Lookup(t testing.TB) *oc.QualifiedE_Acl_LOG_ACTION {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Actions", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Actions_LogActionPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_Acl_LOG_ACTION{
		Metadata: md,
	}).SetVal(goStruct.GetLogAction())
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Get(t testing.TB) oc.E_Acl_LOG_ACTION {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_LOG_ACTION {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_LOG_ACTION
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Actions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Actions_LogActionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPathAny) Get(t testing.TB) []oc.E_Acl_LOG_ACTION {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_LOG_ACTION
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_LOG_ACTION {
	t.Helper()
	c := &oc.CollectionE_Acl_LOG_ACTION{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_LOG_ACTION) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Actions_LogActionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Acl_LOG_ACTION) bool) *oc.E_Acl_LOG_ACTIONWatcher {
	t.Helper()
	w := &oc.E_Acl_LOG_ACTIONWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Actions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Actions", gs, queryPath, true, false)
		return convertAcl_AclSet_AclEntry_Actions_LogActionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Acl_LOG_ACTION)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_LOG_ACTION) bool) *oc.E_Acl_LOG_ACTIONWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Actions_LogActionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Await(t testing.TB, timeout time.Duration, val oc.E_Acl_LOG_ACTION) *oc.QualifiedE_Acl_LOG_ACTION {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Acl_LOG_ACTION) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action to the batch object.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_LOG_ACTION {
	t.Helper()
	c := &oc.CollectionE_Acl_LOG_ACTION{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_LOG_ACTION) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_LOG_ACTION) bool) *oc.E_Acl_LOG_ACTIONWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Actions_LogActionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/state/log-action to the batch object.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Actions_LogActionPath extracts the value of the leaf LogAction from its parent oc.Acl_AclSet_AclEntry_Actions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_LOG_ACTION.
func convertAcl_AclSet_AclEntry_Actions_LogActionPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Actions) *oc.QualifiedE_Acl_LOG_ACTION {
	t.Helper()
	qv := &oc.QualifiedE_Acl_LOG_ACTION{
		Metadata: md,
	}
	val := parent.LogAction
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_DescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry", gs, queryPath, true, false)
		return convertAcl_AclSet_AclEntry_DescriptionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_DescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description to the batch object.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_DescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_DescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_DescriptionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/description to the batch object.
func (n *Acl_AclSet_AclEntry_DescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_DescriptionPath extracts the value of the leaf Description from its parent oc.Acl_AclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Description
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
