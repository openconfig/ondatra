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

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet", goStruct, true, false)
	if ok {
		return convertAcl_Interface_EgressAclSet_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_TypePath) Get(t testing.TB) oc.E_Acl_ACL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_ACL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_TypePathAny) Get(t testing.TB) []oc.E_Acl_ACL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_ACL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_TypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_ACL_TYPE {
	t.Helper()
	c := &oc.CollectionE_Acl_ACL_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_ACL_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_TypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_TYPE) bool) *oc.E_Acl_ACL_TYPEWatcher {
	t.Helper()
	w := &oc.E_Acl_ACL_TYPEWatcher{}
	gs := &oc.Acl_Interface_EgressAclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet", gs, queryPath, true, false)
		return convertAcl_Interface_EgressAclSet_TypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Acl_ACL_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_TypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_TYPE) bool) *oc.E_Acl_ACL_TYPEWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_TypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_TypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Acl_ACL_TYPE) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Acl_ACL_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type to the batch object.
func (n *Acl_Interface_EgressAclSet_TypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_TypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_ACL_TYPE {
	t.Helper()
	c := &oc.CollectionE_Acl_ACL_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_ACL_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_TYPE) bool) *oc.E_Acl_ACL_TYPEWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_TypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/state/type to the batch object.
func (n *Acl_Interface_EgressAclSet_TypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_Interface_EgressAclSet_TypePath extracts the value of the leaf Type from its parent oc.Acl_Interface_EgressAclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_ACL_TYPE.
func convertAcl_Interface_EgressAclSet_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_Acl_ACL_TYPE{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_Interface{}
	md, ok := oc.Lookup(t, n, "Acl_Interface", goStruct, true, false)
	if ok {
		return convertAcl_Interface_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/state/id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/state/id with a ONCE subscription.
func (n *Acl_Interface_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_Interface_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Acl_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface", gs, queryPath, true, false)
		return convertAcl_Interface_IdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_Interface_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_IdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/state/id to the batch object.
func (n *Acl_Interface_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Acl_Interface_IdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/state/id to the batch object.
func (n *Acl_Interface_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_Interface_IdPath extracts the value of the leaf Id from its parent oc.Acl_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_Interface_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_IngressAclSetPath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface_IngressAclSet {
	t.Helper()
	goStruct := &oc.Acl_Interface_IngressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_IngressAclSet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_Interface_IngressAclSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_IngressAclSetPath) Get(t testing.TB) *oc.Acl_Interface_IngressAclSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_IngressAclSetPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface_IngressAclSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface_IngressAclSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_IngressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_IngressAclSet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface_IngressAclSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a ONCE subscription.
func (n *Acl_Interface_IngressAclSetPathAny) Get(t testing.TB) []*oc.Acl_Interface_IngressAclSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface_IngressAclSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_IngressAclSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_IngressAclSet {
	t.Helper()
	c := &oc.CollectionAcl_Interface_IngressAclSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_IngressAclSet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_Interface_IngressAclSet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_Interface_IngressAclSet)))
		return false
	})
	return c
}

func watch_Acl_Interface_IngressAclSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface_IngressAclSet) bool) *oc.Acl_Interface_IngressAclSetWatcher {
	t.Helper()
	w := &oc.Acl_Interface_IngressAclSetWatcher{}
	gs := &oc.Acl_Interface_IngressAclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_IngressAclSet", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_Interface_IngressAclSet{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface_IngressAclSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_IngressAclSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_IngressAclSet) bool) *oc.Acl_Interface_IngressAclSetWatcher {
	t.Helper()
	return watch_Acl_Interface_IngressAclSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_IngressAclSetPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_Interface_IngressAclSet) *oc.QualifiedAcl_Interface_IngressAclSet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_Interface_IngressAclSet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set to the batch object.
func (n *Acl_Interface_IngressAclSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_IngressAclSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_IngressAclSet {
	t.Helper()
	c := &oc.CollectionAcl_Interface_IngressAclSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_IngressAclSet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_IngressAclSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_IngressAclSet) bool) *oc.Acl_Interface_IngressAclSetWatcher {
	t.Helper()
	return watch_Acl_Interface_IngressAclSetPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set to the batch object.
func (n *Acl_Interface_IngressAclSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_IngressAclSet_AclEntryPath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface_IngressAclSet_AclEntry {
	t.Helper()
	goStruct := &oc.Acl_Interface_IngressAclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_IngressAclSet_AclEntry", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_Interface_IngressAclSet_AclEntry{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_IngressAclSet_AclEntryPath) Get(t testing.TB) *oc.Acl_Interface_IngressAclSet_AclEntry {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_IngressAclSet_AclEntryPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface_IngressAclSet_AclEntry {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface_IngressAclSet_AclEntry
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_IngressAclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_IngressAclSet_AclEntry", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface_IngressAclSet_AclEntry{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry with a ONCE subscription.
func (n *Acl_Interface_IngressAclSet_AclEntryPathAny) Get(t testing.TB) []*oc.Acl_Interface_IngressAclSet_AclEntry {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface_IngressAclSet_AclEntry
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_IngressAclSet_AclEntryPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_IngressAclSet_AclEntry {
	t.Helper()
	c := &oc.CollectionAcl_Interface_IngressAclSet_AclEntry{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_IngressAclSet_AclEntry) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_Interface_IngressAclSet_AclEntry{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_Interface_IngressAclSet_AclEntry)))
		return false
	})
	return c
}

func watch_Acl_Interface_IngressAclSet_AclEntryPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface_IngressAclSet_AclEntry) bool) *oc.Acl_Interface_IngressAclSet_AclEntryWatcher {
	t.Helper()
	w := &oc.Acl_Interface_IngressAclSet_AclEntryWatcher{}
	gs := &oc.Acl_Interface_IngressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_IngressAclSet_AclEntry", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_Interface_IngressAclSet_AclEntry{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface_IngressAclSet_AclEntry)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_IngressAclSet_AclEntryPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_IngressAclSet_AclEntry) bool) *oc.Acl_Interface_IngressAclSet_AclEntryWatcher {
	t.Helper()
	return watch_Acl_Interface_IngressAclSet_AclEntryPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_IngressAclSet_AclEntryPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_Interface_IngressAclSet_AclEntry) *oc.QualifiedAcl_Interface_IngressAclSet_AclEntry {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_Interface_IngressAclSet_AclEntry) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry to the batch object.
func (n *Acl_Interface_IngressAclSet_AclEntryPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_IngressAclSet_AclEntryPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_IngressAclSet_AclEntry {
	t.Helper()
	c := &oc.CollectionAcl_Interface_IngressAclSet_AclEntry{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_IngressAclSet_AclEntry) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_IngressAclSet_AclEntryPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_IngressAclSet_AclEntry) bool) *oc.Acl_Interface_IngressAclSet_AclEntryWatcher {
	t.Helper()
	return watch_Acl_Interface_IngressAclSet_AclEntryPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/acl-entries/acl-entry to the batch object.
func (n *Acl_Interface_IngressAclSet_AclEntryPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
