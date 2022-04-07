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

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_InterfacePath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface {
	t.Helper()
	goStruct := &oc.Acl_Interface{}
	md, ok := oc.Lookup(t, n, "Acl_Interface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_InterfacePath) Get(t testing.TB) *oc.Acl_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface with a ONCE subscription.
func (n *Acl_InterfacePathAny) Get(t testing.TB) []*oc.Acl_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface {
	t.Helper()
	c := &oc.CollectionAcl_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_Interface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_Interface)))
		return false
	})
	return c
}

func watch_Acl_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface) bool) *oc.Acl_InterfaceWatcher {
	t.Helper()
	w := &oc.Acl_InterfaceWatcher{}
	gs := &oc.Acl_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface", gs, queryPath, false, false)
		qv := (&oc.QualifiedAcl_Interface{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface) bool) *oc.Acl_InterfaceWatcher {
	t.Helper()
	return watch_Acl_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_InterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_Interface) *oc.QualifiedAcl_Interface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_Interface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface to the batch object.
func (n *Acl_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface {
	t.Helper()
	c := &oc.CollectionAcl_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface) bool) *oc.Acl_InterfaceWatcher {
	t.Helper()
	w := &oc.Acl_InterfaceWatcher{}
	structs := map[string]*oc.Acl_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_Interface", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedAcl_Interface{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface) bool) *oc.Acl_InterfaceWatcher {
	t.Helper()
	return watch_Acl_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface to the batch object.
func (n *Acl_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSetPath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface_EgressAclSet {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_Interface_EgressAclSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSetPath) Get(t testing.TB) *oc.Acl_Interface_EgressAclSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSetPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface_EgressAclSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface_EgressAclSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface_EgressAclSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription.
func (n *Acl_Interface_EgressAclSetPathAny) Get(t testing.TB) []*oc.Acl_Interface_EgressAclSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface_EgressAclSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_EgressAclSet {
	t.Helper()
	c := &oc.CollectionAcl_Interface_EgressAclSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_EgressAclSet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_Interface_EgressAclSet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_Interface_EgressAclSet)))
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet) bool) *oc.Acl_Interface_EgressAclSetWatcher {
	t.Helper()
	w := &oc.Acl_Interface_EgressAclSetWatcher{}
	gs := &oc.Acl_Interface_EgressAclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet", gs, queryPath, false, false)
		qv := (&oc.QualifiedAcl_Interface_EgressAclSet{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface_EgressAclSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet) bool) *oc.Acl_Interface_EgressAclSetWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSetPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_Interface_EgressAclSet) *oc.QualifiedAcl_Interface_EgressAclSet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_Interface_EgressAclSet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set to the batch object.
func (n *Acl_Interface_EgressAclSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_EgressAclSet {
	t.Helper()
	c := &oc.CollectionAcl_Interface_EgressAclSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_EgressAclSet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet) bool) *oc.Acl_Interface_EgressAclSetWatcher {
	t.Helper()
	w := &oc.Acl_Interface_EgressAclSetWatcher{}
	structs := map[string]*oc.Acl_Interface_EgressAclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_Interface_EgressAclSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_Interface_EgressAclSet", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedAcl_Interface_EgressAclSet{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface_EgressAclSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet) bool) *oc.Acl_Interface_EgressAclSetWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set to the batch object.
func (n *Acl_Interface_EgressAclSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet_AclEntry", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_Interface_EgressAclSet_AclEntry{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Get(t testing.TB) *oc.Acl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntryPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface_EgressAclSet_AclEntry
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface_EgressAclSet_AclEntry{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_AclEntryPathAny) Get(t testing.TB) []*oc.Acl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface_EgressAclSet_AclEntry
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	c := &oc.CollectionAcl_Interface_EgressAclSet_AclEntry{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_Interface_EgressAclSet_AclEntry{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_Interface_EgressAclSet_AclEntry)))
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntryPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool) *oc.Acl_Interface_EgressAclSet_AclEntryWatcher {
	t.Helper()
	w := &oc.Acl_Interface_EgressAclSet_AclEntryWatcher{}
	gs := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", gs, queryPath, false, false)
		qv := (&oc.QualifiedAcl_Interface_EgressAclSet_AclEntry{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface_EgressAclSet_AclEntry)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool) *oc.Acl_Interface_EgressAclSet_AclEntryWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntryPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_Interface_EgressAclSet_AclEntry) *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntryPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	c := &oc.CollectionAcl_Interface_EgressAclSet_AclEntry{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntryPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool) *oc.Acl_Interface_EgressAclSet_AclEntryWatcher {
	t.Helper()
	w := &oc.Acl_Interface_EgressAclSet_AclEntryWatcher{}
	structs := map[string]*oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_Interface_EgressAclSet_AclEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedAcl_Interface_EgressAclSet_AclEntry{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface_EgressAclSet_AclEntry)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntryPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool) *oc.Acl_Interface_EgressAclSet_AclEntryWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntryPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntryPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet_AclEntry", goStruct, true, false)
	if ok {
		return convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_Interface_EgressAclSet_AclEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", structs[pre], queryPath, true, false)
			qv := convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-octets to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath extracts the value of the leaf MatchedOctets from its parent oc.Acl_Interface_EgressAclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertAcl_Interface_EgressAclSet_AclEntry_MatchedOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet_AclEntry) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MatchedOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet_AclEntry", goStruct, true, false)
	if ok {
		return convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_Interface_EgressAclSet_AclEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", structs[pre], queryPath, true, false)
			qv := convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/matched-packets to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_MatchedPacketsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath extracts the value of the leaf MatchedPackets from its parent oc.Acl_Interface_EgressAclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertAcl_Interface_EgressAclSet_AclEntry_MatchedPacketsPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet_AclEntry) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MatchedPackets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet_AclEntry", goStruct, true, false)
	if ok {
		return convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_Interface_EgressAclSet_AclEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", structs[pre], queryPath, true, false)
			qv := convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry/state/sequence-id to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntry_SequenceIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath extracts the value of the leaf SequenceId from its parent oc.Acl_Interface_EgressAclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertAcl_Interface_EgressAclSet_AclEntry_SequenceIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet_AclEntry) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.SequenceId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet", goStruct, true, false)
	if ok {
		return convertAcl_Interface_EgressAclSet_SetNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_SetNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_SetNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_Interface_EgressAclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_Interface_EgressAclSet_SetNamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_SetNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name to the batch object.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_SetNamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Acl_Interface_EgressAclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_Interface_EgressAclSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_Interface_EgressAclSet", structs[pre], queryPath, true, false)
			qv := convertAcl_Interface_EgressAclSet_SetNamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_SetNamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/set-name to the batch object.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_Interface_EgressAclSet_SetNamePath extracts the value of the leaf SetName from its parent oc.Acl_Interface_EgressAclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_Interface_EgressAclSet_SetNamePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SetName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
