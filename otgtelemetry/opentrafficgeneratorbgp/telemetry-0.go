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

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeerPath) Lookup(t testing.TB) *oc.QualifiedBgpPeer {
	t.Helper()
	goStruct := &oc.BgpPeer{}
	md, ok := oc.Lookup(t, n, "BgpPeer", goStruct, false, false)
	if ok {
		return (&oc.QualifiedBgpPeer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeerPath) Get(t testing.TB) *oc.BgpPeer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeerPathAny) Lookup(t testing.TB) []*oc.QualifiedBgpPeer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBgpPeer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedBgpPeer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer with a ONCE subscription.
func (n *BgpPeerPathAny) Get(t testing.TB) []*oc.BgpPeer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.BgpPeer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer {
	t.Helper()
	c := &oc.CollectionBgpPeer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedBgpPeer{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.BgpPeer)))
		return false
	})
	return c
}

func watch_BgpPeerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer) bool) *oc.BgpPeerWatcher {
	t.Helper()
	w := &oc.BgpPeerWatcher{}
	gs := &oc.BgpPeer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer", gs, queryPath, false, false)
		qv := (&oc.QualifiedBgpPeer{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer) bool) *oc.BgpPeerWatcher {
	t.Helper()
	return watch_BgpPeerPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeerPath) Await(t testing.TB, timeout time.Duration, val *oc.BgpPeer) *oc.QualifiedBgpPeer {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBgpPeer) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer to the batch object.
func (n *BgpPeerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer {
	t.Helper()
	c := &oc.CollectionBgpPeer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer) bool) *oc.BgpPeerWatcher {
	t.Helper()
	w := &oc.BgpPeerWatcher{}
	structs := map[string]*oc.BgpPeer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedBgpPeer{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer) bool) *oc.BgpPeerWatcher {
	t.Helper()
	return watch_BgpPeerPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer to the batch object.
func (n *BgpPeerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_CountersPath) Lookup(t testing.TB) *oc.QualifiedBgpPeer_Counters {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedBgpPeer_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_CountersPath) Get(t testing.TB) *oc.BgpPeer_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedBgpPeer_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBgpPeer_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedBgpPeer_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters with a ONCE subscription.
func (n *BgpPeer_CountersPathAny) Get(t testing.TB) []*oc.BgpPeer_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.BgpPeer_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer_Counters {
	t.Helper()
	c := &oc.CollectionBgpPeer_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedBgpPeer_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.BgpPeer_Counters)))
		return false
	})
	return c
}

func watch_BgpPeer_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer_Counters) bool) *oc.BgpPeer_CountersWatcher {
	t.Helper()
	w := &oc.BgpPeer_CountersWatcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedBgpPeer_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer_Counters) bool) *oc.BgpPeer_CountersWatcher {
	t.Helper()
	return watch_BgpPeer_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.BgpPeer_Counters) *oc.QualifiedBgpPeer_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBgpPeer_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters to the batch object.
func (n *BgpPeer_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer_Counters {
	t.Helper()
	c := &oc.CollectionBgpPeer_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer_Counters) bool) *oc.BgpPeer_CountersWatcher {
	t.Helper()
	w := &oc.BgpPeer_CountersWatcher{}
	structs := map[string]*oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedBgpPeer_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer_Counters) bool) *oc.BgpPeer_CountersWatcher {
	t.Helper()
	return watch_BgpPeer_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters to the batch object.
func (n *BgpPeer_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_FlapsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_FlapsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_FlapsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_FlapsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertBgpPeer_Counters_FlapsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps with a ONCE subscription.
func (n *BgpPeer_Counters_FlapsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_FlapsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_FlapsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_FlapsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_FlapsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_FlapsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_FlapsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps to the batch object.
func (n *BgpPeer_Counters_FlapsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_FlapsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_FlapsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertBgpPeer_Counters_FlapsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_FlapsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_FlapsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/flaps to the batch object.
func (n *BgpPeer_Counters_FlapsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_FlapsPath extracts the value of the leaf Flaps from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_FlapsPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Flaps
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_InKeepalivesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_InKeepalivesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_InKeepalivesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_InKeepalivesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertBgpPeer_Counters_InKeepalivesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives with a ONCE subscription.
func (n *BgpPeer_Counters_InKeepalivesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InKeepalivesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InKeepalivesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_InKeepalivesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InKeepalivesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InKeepalivesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_InKeepalivesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives to the batch object.
func (n *BgpPeer_Counters_InKeepalivesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InKeepalivesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InKeepalivesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertBgpPeer_Counters_InKeepalivesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InKeepalivesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InKeepalivesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-keepalives to the batch object.
func (n *BgpPeer_Counters_InKeepalivesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_InKeepalivesPath extracts the value of the leaf InKeepalives from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_InKeepalivesPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InKeepalives
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_InNotificationsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_InNotificationsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_InNotificationsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_InNotificationsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertBgpPeer_Counters_InNotificationsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications with a ONCE subscription.
func (n *BgpPeer_Counters_InNotificationsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InNotificationsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InNotificationsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_InNotificationsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InNotificationsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InNotificationsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_InNotificationsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications to the batch object.
func (n *BgpPeer_Counters_InNotificationsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InNotificationsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InNotificationsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertBgpPeer_Counters_InNotificationsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InNotificationsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InNotificationsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-notifications to the batch object.
func (n *BgpPeer_Counters_InNotificationsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_InNotificationsPath extracts the value of the leaf InNotifications from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_InNotificationsPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InNotifications
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_Counters_InOpensPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.BgpPeer_Counters{}
	md, ok := oc.Lookup(t, n, "BgpPeer_Counters", goStruct, true, false)
	if ok {
		return convertBgpPeer_Counters_InOpensPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_Counters_InOpensPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_Counters_InOpensPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertBgpPeer_Counters_InOpensPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens with a ONCE subscription.
func (n *BgpPeer_Counters_InOpensPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InOpensPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InOpensPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.BgpPeer_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_Counters_InOpensPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InOpensPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InOpensPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_Counters_InOpensPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens to the batch object.
func (n *BgpPeer_Counters_InOpensPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_Counters_InOpensPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_Counters_InOpensPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertBgpPeer_Counters_InOpensPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_Counters_InOpensPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_BgpPeer_Counters_InOpensPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/state/counters/in-opens to the batch object.
func (n *BgpPeer_Counters_InOpensPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_Counters_InOpensPath extracts the value of the leaf InOpens from its parent oc.BgpPeer_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertBgpPeer_Counters_InOpensPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InOpens
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
