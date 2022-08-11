package isis

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry/otg"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/default-metric to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath extracts the value of the leaf DefaultMetric from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_DefaultMetricPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.DefaultMetric
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath) Lookup(t testing.TB) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath) Get(t testing.TB) oc.E_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePathAny) Get(t testing.TB) []oc.E_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ipv4ExternalReachability_Prefix_OriginType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	c := &oc.CollectionE_Ipv4ExternalReachability_Prefix_OriginType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool) *oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool) *oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Ipv4ExternalReachability_Prefix_OriginType) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	c := &oc.CollectionE_Ipv4ExternalReachability_Prefix_OriginType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool) *oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool) *oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/origin-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath extracts the value of the leaf OriginType from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_OriginTypePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	qv := &oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType{
		Metadata: md,
	}
	val := parent.OriginType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix-length to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.PrefixLength
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath extracts the value of the leaf Prefix from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_PrefixPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Prefix
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath) Lookup(t testing.TB) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath) Get(t testing.TB) oc.E_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePathAny) Get(t testing.TB) []oc.E_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ipv4ExternalReachability_Prefix_RedistributionType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	c := &oc.CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool) *oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool) *oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Ipv4ExternalReachability_Prefix_RedistributionType) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	c := &oc.CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool) *oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool) *oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability/prefixes/prefix/state/redistribution-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath extracts the value of the leaf RedistributionType from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix_RedistributionTypePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	qv := &oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType{
		Metadata: md,
	}
	val := parent.RedistributionType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/default-metric to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath extracts the value of the leaf DefaultMetric from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_DefaultMetricPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.DefaultMetric
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath) Lookup(t testing.TB) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath) Get(t testing.TB) oc.E_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePathAny) Get(t testing.TB) []oc.E_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ipv4ExternalReachability_Prefix_OriginType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	c := &oc.CollectionE_Ipv4ExternalReachability_Prefix_OriginType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool) *oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool) *oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Ipv4ExternalReachability_Prefix_OriginType) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	c := &oc.CollectionE_Ipv4ExternalReachability_Prefix_OriginType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool) *oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) bool) *oc.E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/origin-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath extracts the value of the leaf OriginType from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_OriginTypePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	qv := &oc.QualifiedE_Ipv4ExternalReachability_Prefix_OriginType{
		Metadata: md,
	}
	val := parent.OriginType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix-length to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.PrefixLength
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath extracts the value of the leaf Prefix from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_PrefixPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Prefix
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath) Lookup(t testing.TB) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath) Get(t testing.TB) oc.E_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePathAny) Get(t testing.TB) []oc.E_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ipv4ExternalReachability_Prefix_RedistributionType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	c := &oc.CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool) *oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool) *oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Ipv4ExternalReachability_Prefix_RedistributionType) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	c := &oc.CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool) *oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) bool) *oc.E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-internal-reachability/prefixes/prefix/state/redistribution-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath extracts the value of the leaf RedistributionType from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix_RedistributionTypePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) *oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	qv := &oc.QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType{
		Metadata: md,
	}
	val := parent.RedistributionType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/metric to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath extracts the value of the leaf Metric from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_MetricPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Metric
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath) Lookup(t testing.TB) *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath) Get(t testing.TB) oc.E_Ipv6Reachability_Prefix_OriginType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ipv6Reachability_Prefix_OriginType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ipv6Reachability_Prefix_OriginType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePathAny) Get(t testing.TB) []oc.E_Ipv6Reachability_Prefix_OriginType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ipv6Reachability_Prefix_OriginType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv6Reachability_Prefix_OriginType {
	t.Helper()
	c := &oc.CollectionE_Ipv6Reachability_Prefix_OriginType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType) bool) *oc.E_Ipv6Reachability_Prefix_OriginTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv6Reachability_Prefix_OriginTypeWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv6Reachability_Prefix_OriginType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType) bool) *oc.E_Ipv6Reachability_Prefix_OriginTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Ipv6Reachability_Prefix_OriginType) *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv6Reachability_Prefix_OriginType {
	t.Helper()
	c := &oc.CollectionE_Ipv6Reachability_Prefix_OriginType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType) bool) *oc.E_Ipv6Reachability_Prefix_OriginTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv6Reachability_Prefix_OriginTypeWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv6Reachability_Prefix_OriginType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType) bool) *oc.E_Ipv6Reachability_Prefix_OriginTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/origin-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath extracts the value of the leaf OriginType from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_OriginTypePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) *oc.QualifiedE_Ipv6Reachability_Prefix_OriginType {
	t.Helper()
	qv := &oc.QualifiedE_Ipv6Reachability_Prefix_OriginType{
		Metadata: md,
	}
	val := parent.OriginType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath) Lookup(t testing.TB) *oc.QualifiedE_State_FlagsSlice {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath) Get(t testing.TB) []oc.E_State_Flags {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_State_FlagsSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_State_FlagsSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPathAny) Get(t testing.TB) [][]oc.E_State_Flags {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_State_Flags
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_State_FlagsSlice {
	t.Helper()
	c := &oc.CollectionE_State_FlagsSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_State_FlagsSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_State_FlagsSlice) bool) *oc.E_State_FlagsSliceWatcher {
	t.Helper()
	w := &oc.E_State_FlagsSliceWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_State_FlagsSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_State_FlagsSlice) bool) *oc.E_State_FlagsSliceWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath) Await(t testing.TB, timeout time.Duration, val []oc.E_State_Flags) *oc.QualifiedE_State_FlagsSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_State_FlagsSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_State_FlagsSlice {
	t.Helper()
	c := &oc.CollectionE_State_FlagsSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_State_FlagsSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_State_FlagsSlice) bool) *oc.E_State_FlagsSliceWatcher {
	t.Helper()
	w := &oc.E_State_FlagsSliceWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_State_FlagsSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_State_FlagsSlice) bool) *oc.E_State_FlagsSliceWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-attributes/flags to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath extracts the value of the leaf Flags from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes
// and combines the update with an existing Metadata to return a *oc.QualifiedE_State_FlagsSlice.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes_FlagsPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) *oc.QualifiedE_State_FlagsSlice {
	t.Helper()
	qv := &oc.QualifiedE_State_FlagsSlice{
		Metadata: md,
	}
	val := parent.Flags
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix-length to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.PrefixLength
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath extracts the value of the leaf Prefix from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Prefix
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath) Lookup(t testing.TB) *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath) Get(t testing.TB) oc.E_Ipv6Reachability_Prefix_RedistributionType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePathAny) Get(t testing.TB) []oc.E_Ipv6Reachability_Prefix_RedistributionType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ipv6Reachability_Prefix_RedistributionType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv6Reachability_Prefix_RedistributionType {
	t.Helper()
	c := &oc.CollectionE_Ipv6Reachability_Prefix_RedistributionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType) bool) *oc.E_Ipv6Reachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv6Reachability_Prefix_RedistributionTypeWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType) bool) *oc.E_Ipv6Reachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Ipv6Reachability_Prefix_RedistributionType) *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Ipv6Reachability_Prefix_RedistributionType {
	t.Helper()
	c := &oc.CollectionE_Ipv6Reachability_Prefix_RedistributionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType) bool) *oc.E_Ipv6Reachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	w := &oc.E_Ipv6Reachability_Prefix_RedistributionTypeWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType) bool) *oc.E_Ipv6Reachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv6-reachability/prefixes/prefix/state/redistribution-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath extracts the value of the leaf RedistributionType from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_RedistributionTypePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) *oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType {
	t.Helper()
	qv := &oc.QualifiedE_Ipv6Reachability_Prefix_RedistributionType{
		Metadata: md,
	}
	val := parent.RedistributionType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/is-reachability/neighbors/neighbor/state/system-id to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath extracts the value of the leaf SystemId from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor_SystemIdPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SystemId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.IsisRouter{}
	md, ok := oc.Lookup(t, n, "IsisRouter", goStruct, true, false)
	if ok {
		return convertIsisRouter_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a ONCE subscription.
func (n *IsisRouter_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.IsisRouter{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_NamePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/name to the batch object.
func (n *IsisRouter_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.IsisRouter{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/name to the batch object.
func (n *IsisRouter_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_NamePath extracts the value of the leaf Name from its parent oc.IsisRouter
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertIsisRouter_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter) *oc.QualifiedString {
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
