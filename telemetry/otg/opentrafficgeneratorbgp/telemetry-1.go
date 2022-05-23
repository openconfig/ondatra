package opentrafficgeneratorbgp

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

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_InRouteWithdrawPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_InRouteWithdrawPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_InRouteWithdrawPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_InRouteWithdrawPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_Counters_InRouteWithdrawPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw with a ONCE subscription.
func (n *BgpPeer_Counters_InRouteWithdrawPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InRouteWithdrawPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InRouteWithdrawPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_InRouteWithdrawPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InRouteWithdrawPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InRouteWithdrawPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_InRouteWithdrawPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw to the batch object.
func (n *BgpPeer_Counters_InRouteWithdrawPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InRouteWithdrawPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InRouteWithdrawPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_Counters", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_Counters_InRouteWithdrawPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InRouteWithdrawPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InRouteWithdrawPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-route-withdraw to the batch object.
func (n *BgpPeer_Counters_InRouteWithdrawPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_InRouteWithdrawPath extracts the value of the leaf InRouteWithdraw from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_InRouteWithdrawPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InRouteWithdraw
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_InRoutesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_InRoutesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_InRoutesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_InRoutesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_Counters_InRoutesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes with a ONCE subscription.
func (n *BgpPeer_Counters_InRoutesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InRoutesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InRoutesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_InRoutesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InRoutesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InRoutesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_InRoutesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes to the batch object.
func (n *BgpPeer_Counters_InRoutesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InRoutesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InRoutesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_Counters", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_Counters_InRoutesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InRoutesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InRoutesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-routes to the batch object.
func (n *BgpPeer_Counters_InRoutesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_InRoutesPath extracts the value of the leaf InRoutes from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_InRoutesPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InRoutes
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_InUpdatesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_InUpdatesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_InUpdatesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_InUpdatesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_Counters_InUpdatesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates with a ONCE subscription.
func (n *BgpPeer_Counters_InUpdatesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InUpdatesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InUpdatesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_InUpdatesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InUpdatesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InUpdatesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_InUpdatesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates to the batch object.
func (n *BgpPeer_Counters_InUpdatesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InUpdatesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InUpdatesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_Counters", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_Counters_InUpdatesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InUpdatesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InUpdatesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-updates to the batch object.
func (n *BgpPeer_Counters_InUpdatesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_InUpdatesPath extracts the value of the leaf InUpdates from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_InUpdatesPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InUpdates
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_OutKeepalivesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_OutKeepalivesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_OutKeepalivesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_OutKeepalivesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_Counters_OutKeepalivesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives with a ONCE subscription.
func (n *BgpPeer_Counters_OutKeepalivesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutKeepalivesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutKeepalivesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_OutKeepalivesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutKeepalivesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutKeepalivesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_OutKeepalivesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives to the batch object.
func (n *BgpPeer_Counters_OutKeepalivesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutKeepalivesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutKeepalivesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_Counters", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_Counters_OutKeepalivesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutKeepalivesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutKeepalivesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-keepalives to the batch object.
func (n *BgpPeer_Counters_OutKeepalivesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_OutKeepalivesPath extracts the value of the leaf OutKeepalives from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_OutKeepalivesPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutKeepalives
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_OutNotificationsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_OutNotificationsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_OutNotificationsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_OutNotificationsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_Counters_OutNotificationsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications with a ONCE subscription.
func (n *BgpPeer_Counters_OutNotificationsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutNotificationsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutNotificationsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_OutNotificationsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutNotificationsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutNotificationsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_OutNotificationsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications to the batch object.
func (n *BgpPeer_Counters_OutNotificationsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutNotificationsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutNotificationsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_Counters", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_Counters_OutNotificationsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutNotificationsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutNotificationsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-notifications to the batch object.
func (n *BgpPeer_Counters_OutNotificationsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_OutNotificationsPath extracts the value of the leaf OutNotifications from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_OutNotificationsPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutNotifications
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_OutOpensPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_OutOpensPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_OutOpensPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_OutOpensPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_Counters_OutOpensPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens with a ONCE subscription.
func (n *BgpPeer_Counters_OutOpensPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutOpensPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutOpensPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_OutOpensPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutOpensPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutOpensPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_OutOpensPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens to the batch object.
func (n *BgpPeer_Counters_OutOpensPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutOpensPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutOpensPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_Counters", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_Counters_OutOpensPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutOpensPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutOpensPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-opens to the batch object.
func (n *BgpPeer_Counters_OutOpensPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_OutOpensPath extracts the value of the leaf OutOpens from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_OutOpensPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutOpens
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
