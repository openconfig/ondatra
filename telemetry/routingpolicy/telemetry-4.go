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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix)))
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_PrefixSet_PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) bool) *oc.RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) bool) *oc.RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_PrefixSet_PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix) *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix to the batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) bool) *oc.RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_PrefixSet_PrefixPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix to the batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix to the batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix to the batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath extracts the value of the leaf IpPrefix from its parent oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.IpPrefix
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range to the batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range to the batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath extracts the value of the leaf MasklengthRange from its parent oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MasklengthRange
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_TagSet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_TagSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_TagSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_TagSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_TagSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_TagSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_TagSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_TagSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_TagSet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_DefinedSets_TagSet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_DefinedSets_TagSet)))
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_TagSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_TagSet) bool) *oc.RoutingPolicy_DefinedSets_TagSetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_TagSetWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_TagSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_DefinedSets_TagSet{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets_TagSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_TagSet) bool) *oc.RoutingPolicy_DefinedSets_TagSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_TagSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_DefinedSets_TagSet) *oc.QualifiedRoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_DefinedSets_TagSet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set to the batch object.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_TagSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_TagSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_TagSet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_TagSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_TagSet) bool) *oc.RoutingPolicy_DefinedSets_TagSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_TagSetPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set to the batch object.
func (n *RoutingPolicy_DefinedSets_TagSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_TagSet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_TagSet_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_TagSet_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_TagSet_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_TagSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", gs, queryPath, true, false)
		return convertRoutingPolicy_DefinedSets_TagSet_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_TagSet_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name to the batch object.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_TagSet_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name to the batch object.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_TagSet_NamePath extracts the value of the leaf Name from its parent oc.RoutingPolicy_DefinedSets_TagSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_TagSet_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_TagSet) *oc.QualifiedString {
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
