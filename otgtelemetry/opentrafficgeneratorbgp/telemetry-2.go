package opentrafficgeneratorbgp

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/otgtelemetry"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_OutRouteWithdrawPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_OutRouteWithdrawPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_OutRouteWithdrawPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_OutRouteWithdrawPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertBgpPeer_Counters_OutRouteWithdrawPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw with a ONCE subscription.
func (n *BgpPeer_Counters_OutRouteWithdrawPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutRouteWithdrawPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutRouteWithdrawPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_OutRouteWithdrawPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutRouteWithdrawPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutRouteWithdrawPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_OutRouteWithdrawPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw to the batch object.
func (n *BgpPeer_Counters_OutRouteWithdrawPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutRouteWithdrawPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutRouteWithdrawPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertBgpPeer_Counters_OutRouteWithdrawPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutRouteWithdrawPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutRouteWithdrawPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-route-withdraw to the batch object.
func (n *BgpPeer_Counters_OutRouteWithdrawPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_OutRouteWithdrawPath extracts the value of the leaf OutRouteWithdraw from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_OutRouteWithdrawPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutRouteWithdraw
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_OutRoutesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_OutRoutesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_OutRoutesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_OutRoutesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertBgpPeer_Counters_OutRoutesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes with a ONCE subscription.
func (n *BgpPeer_Counters_OutRoutesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutRoutesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutRoutesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_OutRoutesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutRoutesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutRoutesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_OutRoutesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes to the batch object.
func (n *BgpPeer_Counters_OutRoutesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutRoutesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutRoutesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertBgpPeer_Counters_OutRoutesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutRoutesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutRoutesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-routes to the batch object.
func (n *BgpPeer_Counters_OutRoutesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_OutRoutesPath extracts the value of the leaf OutRoutes from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_OutRoutesPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutRoutes
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_OutUpdatesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_OutUpdatesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_OutUpdatesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_OutUpdatesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertBgpPeer_Counters_OutUpdatesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates with a ONCE subscription.
func (n *BgpPeer_Counters_OutUpdatesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutUpdatesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutUpdatesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_OutUpdatesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutUpdatesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutUpdatesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_OutUpdatesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates to the batch object.
func (n *BgpPeer_Counters_OutUpdatesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_OutUpdatesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_OutUpdatesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertBgpPeer_Counters_OutUpdatesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_OutUpdatesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_OutUpdatesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/out-updates to the batch object.
func (n *BgpPeer_Counters_OutUpdatesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_OutUpdatesPath extracts the value of the leaf OutUpdates from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_OutUpdatesPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutUpdates
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.BgpPeer{}
	md, ok := oc.Lookup(t, n, "BgpPeer", goStruct, true, false)
	if ok {
		return convertBgpPeer_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name with a ONCE subscription.
func (n *BgpPeer_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.BgpPeer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_NamePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name to the batch object.
func (n *BgpPeer_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.BgpPeer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/name to the batch object.
func (n *BgpPeer_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_NamePath extracts the value of the leaf Name from its parent oc.BgpPeer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertBgpPeer_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer) *oc.QualifiedString {
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

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_SessionStatePath) Lookup(t testing.TB) *oc.QualifiedE_BgpPeer_SessionState {
	t.Helper()
	goStruct := &oc.BgpPeer{}
	md, ok := oc.Lookup(t, n, "BgpPeer", goStruct, true, false)
	if ok {
		return convertBgpPeer_SessionStatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_SessionStatePath) Get(t testing.TB) oc.E_BgpPeer_SessionState {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_SessionStatePathAny) Lookup(t testing.TB) []*oc.QualifiedE_BgpPeer_SessionState {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_BgpPeer_SessionState
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_SessionStatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state with a ONCE subscription.
func (n *BgpPeer_SessionStatePathAny) Get(t testing.TB) []oc.E_BgpPeer_SessionState {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_BgpPeer_SessionState
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_SessionStatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_BgpPeer_SessionState {
	t.Helper()
	c := &oc.CollectionE_BgpPeer_SessionState{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_BgpPeer_SessionState) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_SessionStatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_BgpPeer_SessionState) bool) *oc.E_BgpPeer_SessionStateWatcher {
	t.Helper()
	w := &oc.E_BgpPeer_SessionStateWatcher{}
	gs := &oc.BgpPeer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_SessionStatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_BgpPeer_SessionState)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_SessionStatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_BgpPeer_SessionState) bool) *oc.E_BgpPeer_SessionStateWatcher {
	t.Helper()
	return watch_BgpPeer_SessionStatePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_SessionStatePath) Await(t testing.TB, timeout time.Duration, val oc.E_BgpPeer_SessionState) *oc.QualifiedE_BgpPeer_SessionState {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_BgpPeer_SessionState) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state to the batch object.
func (n *BgpPeer_SessionStatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_SessionStatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_BgpPeer_SessionState {
	t.Helper()
	c := &oc.CollectionE_BgpPeer_SessionState{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_BgpPeer_SessionState) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_SessionStatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_BgpPeer_SessionState) bool) *oc.E_BgpPeer_SessionStateWatcher {
	t.Helper()
	w := &oc.E_BgpPeer_SessionStateWatcher{}
	structs := map[string]*oc.BgpPeer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_SessionStatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_BgpPeer_SessionState)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_SessionStatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_BgpPeer_SessionState) bool) *oc.E_BgpPeer_SessionStateWatcher {
	t.Helper()
	return watch_BgpPeer_SessionStatePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/session-state to the batch object.
func (n *BgpPeer_SessionStatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_SessionStatePath extracts the value of the leaf SessionState from its parent oc.BgpPeer
// and combines the update with an existing Metadata to return a *oc.QualifiedE_BgpPeer_SessionState.
func convertBgpPeer_SessionStatePath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer) *oc.QualifiedE_BgpPeer_SessionState {
	t.Helper()
	qv := &oc.QualifiedE_BgpPeer_SessionState{
		Metadata: md,
	}
	val := parent.SessionState
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
