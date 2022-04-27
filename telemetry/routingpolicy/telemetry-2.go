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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_RoutingPolicy_DefinedSets_PrefixSet_PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) bool) *oc.RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher{}
	structs := map[string]*oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) bool) *oc.RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_PrefixSet_PrefixPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t, md, gs)}, nil
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

func watch_RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/ip-prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t, md, gs)}, nil
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

func watch_RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/state/masklength-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_TagSet{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
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

func watch_RoutingPolicy_DefinedSets_TagSetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_TagSet) bool) *oc.RoutingPolicy_DefinedSets_TagSetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_TagSetWatcher{}
	structs := map[string]*oc.RoutingPolicy_DefinedSets_TagSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_DefinedSets_TagSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_DefinedSets_TagSet{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
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
func (n *RoutingPolicy_DefinedSets_TagSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_TagSet) bool) *oc.RoutingPolicy_DefinedSets_TagSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_TagSetPathAny(t, n, timeout, predicate)
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
// failing the test fatally if no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_DefinedSets_TagSet_NamePath(t, md, gs)}, nil
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

func watch_RoutingPolicy_DefinedSets_TagSet_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_DefinedSets_TagSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_DefinedSets_TagSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_DefinedSets_TagSet_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_TagSet_NamePathAny(t, n, timeout, predicate)
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_TagSet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_DefinedSets_TagSet_TagValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Get(t testing.TB) []oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_TagSet_TagValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePathAny) Get(t testing.TB) [][]oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_TagSet_TagValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) bool) *oc.RoutingPolicy_DefinedSets_TagSet_TagValue_UnionSliceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_TagSet_TagValue_UnionSliceWatcher{}
	gs := &oc.RoutingPolicy_DefinedSets_TagSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_DefinedSets_TagSet_TagValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) bool) *oc.RoutingPolicy_DefinedSets_TagSet_TagValue_UnionSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_TagSet_TagValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Await(t testing.TB, timeout time.Duration, val []oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union) *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value to the batch object.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_DefinedSets_TagSet_TagValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) bool) *oc.RoutingPolicy_DefinedSets_TagSet_TagValue_UnionSliceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_DefinedSets_TagSet_TagValue_UnionSliceWatcher{}
	structs := map[string]*oc.RoutingPolicy_DefinedSets_TagSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_DefinedSets_TagSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_DefinedSets_TagSet_TagValuePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) bool) *oc.RoutingPolicy_DefinedSets_TagSet_TagValue_UnionSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_DefinedSets_TagSet_TagValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/state/tag-value to the batch object.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_DefinedSets_TagSet_TagValuePath extracts the value of the leaf TagValue from its parent oc.RoutingPolicy_DefinedSets_TagSet
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice.
func convertRoutingPolicy_DefinedSets_TagSet_TagValuePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_TagSet) *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice{
		Metadata: md,
	}
	val := parent.TagValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinitionPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinitionPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinitionPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinitionPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinitionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinitionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition) bool) *oc.RoutingPolicy_PolicyDefinitionWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinitionWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinitionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition) bool) *oc.RoutingPolicy_PolicyDefinitionWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinitionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinitionPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition) *oc.QualifiedRoutingPolicy_PolicyDefinition {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition to the batch object.
func (n *RoutingPolicy_PolicyDefinitionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinitionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinitionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition) bool) *oc.RoutingPolicy_PolicyDefinitionWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinitionWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinitionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition) bool) *oc.RoutingPolicy_PolicyDefinitionWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinitionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition to the batch object.
func (n *RoutingPolicy_PolicyDefinitionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name to the batch object.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/state/name to the batch object.
func (n *RoutingPolicy_PolicyDefinition_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_NamePath extracts the value of the leaf Name from its parent oc.RoutingPolicy_PolicyDefinition
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_StatementPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_StatementPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_StatementPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement) bool) *oc.RoutingPolicy_PolicyDefinition_StatementWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_StatementWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement) bool) *oc.RoutingPolicy_PolicyDefinition_StatementWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_StatementPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement to the batch object.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_StatementPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_StatementPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement) bool) *oc.RoutingPolicy_PolicyDefinition_StatementWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_StatementWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_StatementPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement) bool) *oc.RoutingPolicy_PolicyDefinition_StatementWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_StatementPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement to the batch object.
func (n *RoutingPolicy_PolicyDefinition_StatementPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_ActionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_ActionsWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_ActionsWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_ActionsWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_ActionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_ActionsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_ActionsWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_ActionsWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_ActionsWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_ActionsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
