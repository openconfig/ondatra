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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicyPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy {
	t.Helper()
	goStruct := &oc.RoutingPolicy{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicyPath) Get(t testing.TB) *oc.RoutingPolicy {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicyPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy with a ONCE subscription.
func (n *RoutingPolicyPathAny) Get(t testing.TB) []*oc.RoutingPolicy {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy {
	t.Helper()
	c := &oc.CollectionRoutingPolicy{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy)))
		return false
	})
	return c
}

func watch_RoutingPolicyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy) bool) *oc.RoutingPolicyWatcher {
	t.Helper()
	w := &oc.RoutingPolicyWatcher{}
	gs := &oc.RoutingPolicy{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy) bool) *oc.RoutingPolicyWatcher {
	t.Helper()
	return watch_RoutingPolicyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicyPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy) *oc.QualifiedRoutingPolicy {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy to the batch object.
func (n *RoutingPolicyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy {
	t.Helper()
	c := &oc.CollectionRoutingPolicy{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy) bool) *oc.RoutingPolicyWatcher {
	t.Helper()
	return watch_RoutingPolicyPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy to the batch object.
func (n *RoutingPolicyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSetsPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSetsPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSetsPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets with a ONCE subscription.
func (n *RoutingPolicy_DefinedSetsPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_DefinedSets{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_DefinedSets)))
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets) bool) *oc.RoutingPolicy_DefinedSetsWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSetsWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_DefinedSets{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets) bool) *oc.RoutingPolicy_DefinedSetsWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSetsPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_DefinedSets) *oc.QualifiedRoutingPolicy_DefinedSets {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_DefinedSets) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets to the batch object.
func (n *RoutingPolicy_DefinedSetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets) bool) *oc.RoutingPolicy_DefinedSetsWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets to the batch object.
func (n *RoutingPolicy_DefinedSetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_DefinedSets_BgpDefinedSets)))
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSetsWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_BgpDefinedSetsWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSetsWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSetsWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet)))
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-member to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath extracts the value of the leaf AsPathSetMember from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.AsPathSetMember
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/state/as-path-set-name to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath extracts the value of the leaf AsPathSetName from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AsPathSetName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet)))
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Get(t testing.TB) []oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPathAny) Get(t testing.TB) [][]oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSliceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSliceWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Await(t testing.TB, timeout time.Duration, val []oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) bool) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-member to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath extracts the value of the leaf CommunityMember from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice{
		Metadata: md,
	}
	val := parent.CommunityMember
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/community-set-name to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath extracts the value of the leaf CommunitySetName from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CommunitySetName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_PolicyTypes_MatchSetOptionsType{
		Metadata: md,
	}).SetVal(goStruct.GetMatchSetOptions())
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Get(t testing.TB) oc.E_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_MatchSetOptionsType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPathAny) Get(t testing.TB) []oc.E_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_MatchSetOptionsType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_MatchSetOptionsType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool) *oc.E_PolicyTypes_MatchSetOptionsTypeWatcher {
	t.Helper()
	w := &oc.E_PolicyTypes_MatchSetOptionsTypeWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PolicyTypes_MatchSetOptionsType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool) *oc.E_PolicyTypes_MatchSetOptionsTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Await(t testing.TB, timeout time.Duration, val oc.E_PolicyTypes_MatchSetOptionsType) *oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_MatchSetOptionsType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsType) bool) *oc.E_PolicyTypes_MatchSetOptionsTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/state/match-set-options to the batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath extracts the value of the leaf MatchSetOptions from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_MatchSetOptionsType.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) *oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
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
