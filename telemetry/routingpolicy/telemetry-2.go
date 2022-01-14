package routingpolicy

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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet)))
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-member to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath extracts the value of the leaf ExtCommunityMember from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.ExtCommunityMember
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/ext-community-set-name to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath extracts the value of the leaf ExtCommunitySetName from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ExtCommunitySetName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_PolicyTypes_MatchSetOptionsType{
		Metadata: md,
	}).SetVal(goStruct.GetMatchSetOptions())
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Get(t testing.TB) oc.E_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_MatchSetOptionsType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPathAny) Get(t testing.TB) []oc.E_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_MatchSetOptionsType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_MatchSetOptionsType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool) *oc.E_PolicyTypes_MatchSetOptionsTypeWatcher {
	t.Helper()
	w := &oc.E_PolicyTypes_MatchSetOptionsTypeWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PolicyTypes_MatchSetOptionsType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool) *oc.E_PolicyTypes_MatchSetOptionsTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Await(t testing.TB, timeout time.Duration, val oc.E_PolicyTypes_MatchSetOptionsType) *oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_MatchSetOptionsType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool) *oc.E_PolicyTypes_MatchSetOptionsTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/state/match-set-options to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath extracts the value of the leaf MatchSetOptions from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_MatchSetOptionsType.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) *oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_MatchSetOptionsType{
		Metadata: md,
	}
	val := parent.MatchSetOptions
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_NeighborSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_NeighborSet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_NeighborSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_NeighborSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_NeighborSet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_NeighborSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_NeighborSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_NeighborSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_DefinedSets_NeighborSet)))
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_NeighborSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet) bool) *oc.RoutingPolicy_DefinedSets_NeighborSetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_NeighborSetWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_NeighborSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_NeighborSet", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet) bool) *oc.RoutingPolicy_DefinedSets_NeighborSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_NeighborSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_DefinedSets_NeighborSet) *oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set to the batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_NeighborSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_NeighborSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_NeighborSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet) bool) *oc.RoutingPolicy_DefinedSets_NeighborSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_NeighborSetPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set to the batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
