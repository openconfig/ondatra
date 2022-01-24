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
