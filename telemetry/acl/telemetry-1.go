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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_InputInterface {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_InputInterface{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_InputInterface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_InputInterface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_InputInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_InputInterface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_InputInterface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_InputInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_InputInterfacePathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_InputInterface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_InputInterface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_InputInterface {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_InputInterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_InputInterface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_AclSet_AclEntry_InputInterface)))
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_InputInterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_InputInterface) bool) *oc.Acl_AclSet_AclEntry_InputInterfaceWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_InputInterfaceWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_InputInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_InputInterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_InputInterface) bool) *oc.Acl_AclSet_AclEntry_InputInterfaceWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_InputInterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_AclSet_AclEntry_InputInterface) *oc.QualifiedAcl_AclSet_AclEntry_InputInterface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_InputInterface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface to the batch object.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_InputInterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_InputInterface {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_InputInterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_InputInterface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_InputInterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_InputInterface) bool) *oc.Acl_AclSet_AclEntry_InputInterfaceWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_InputInterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface to the batch object.
func (n *Acl_AclSet_AclEntry_InputInterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef)))
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef) bool) *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRefWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRefWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef) bool) *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRefWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef) *oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref to the batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef) bool) *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRefWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref to the batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", gs, queryPath, true, false)
		return convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface to the batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/interface to the batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", gs, queryPath, true, false)
		return convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface to the batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/state/subinterface to the batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4PathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Ipv4
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4PathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_Ipv4
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Ipv4{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Ipv4) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_AclSet_AclEntry_Ipv4{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_AclSet_AclEntry_Ipv4)))
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv4) bool) *oc.Acl_AclSet_AclEntry_Ipv4Watcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Ipv4Watcher{}
	gs := &oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_AclSet_AclEntry_Ipv4{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Ipv4)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv4) bool) *oc.Acl_AclSet_AclEntry_Ipv4Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Await(t testing.TB, timeout time.Duration, val *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedAcl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_Ipv4) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Ipv4{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Ipv4) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Ipv4) bool) *oc.Acl_AclSet_AclEntry_Ipv4Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4Path(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_DestinationAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_DestinationAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", gs, queryPath, true, false)
		return convertAcl_AclSet_AclEntry_Ipv4_DestinationAddressPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_DestinationAddressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/destination-address to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DestinationAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Ipv4_DestinationAddressPath extracts the value of the leaf DestinationAddress from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_Ipv4_DestinationAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Ipv4_DscpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Ipv4_DscpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Ipv4_DscpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Acl_AclSet_AclEntry_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", gs, queryPath, true, false)
		return convertAcl_AclSet_AclEntry_Ipv4_DscpPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_DscpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Ipv4_DscpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4/state/dscp to the batch object.
func (n *Acl_AclSet_AclEntry_Ipv4_DscpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Ipv4_DscpPath extracts the value of the leaf Dscp from its parent oc.Acl_AclSet_AclEntry_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertAcl_AclSet_AclEntry_Ipv4_DscpPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Ipv4) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Dscp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
