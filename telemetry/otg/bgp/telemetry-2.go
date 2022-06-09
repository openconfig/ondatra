package bgp

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

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_PathIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_PathIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_PathIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_PathIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_PathIdPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_PathIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv4Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv4Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv4Prefix_PathIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_PathIdPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/path-id to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_PathIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_PathIdPath extracts the value of the leaf PathId from its parent oc.BgpPeer_UnicastIpv4Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertBgpPeer_UnicastIpv4Prefix_PathIdPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.PathId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_PrefixLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_PrefixLengthPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_PrefixLengthPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_PrefixLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv4Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv4Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv4Prefix_PrefixLengthPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_PrefixLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/prefix-length to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_PrefixLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.BgpPeer_UnicastIpv4Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertBgpPeer_UnicastIpv4Prefix_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv6PrefixPath) Lookup(t testing.TB) *oc.QualifiedBgpPeer_UnicastIpv6Prefix {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv6Prefix", goStruct, false, false)
	if ok {
		return (&oc.QualifiedBgpPeer_UnicastIpv6Prefix{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv6PrefixPath) Get(t testing.TB) *oc.BgpPeer_UnicastIpv6Prefix {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv6PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedBgpPeer_UnicastIpv6Prefix {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBgpPeer_UnicastIpv6Prefix
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedBgpPeer_UnicastIpv6Prefix{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix with a ONCE subscription.
func (n *BgpPeer_UnicastIpv6PrefixPathAny) Get(t testing.TB) []*oc.BgpPeer_UnicastIpv6Prefix {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.BgpPeer_UnicastIpv6Prefix
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer_UnicastIpv6Prefix {
	t.Helper()
	c := &oc.CollectionBgpPeer_UnicastIpv6Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer_UnicastIpv6Prefix) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedBgpPeer_UnicastIpv6Prefix{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.BgpPeer_UnicastIpv6Prefix)))
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv6Prefix) bool) *oc.BgpPeer_UnicastIpv6PrefixWatcher {
	t.Helper()
	w := &oc.BgpPeer_UnicastIpv6PrefixWatcher{}
	gs := &oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", gs, queryPath, false, false)
		qv := (&oc.QualifiedBgpPeer_UnicastIpv6Prefix{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer_UnicastIpv6Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv6Prefix) bool) *oc.BgpPeer_UnicastIpv6PrefixWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv6PrefixPath) Await(t testing.TB, timeout time.Duration, val *oc.BgpPeer_UnicastIpv6Prefix) *oc.QualifiedBgpPeer_UnicastIpv6Prefix {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBgpPeer_UnicastIpv6Prefix) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix to the batch object.
func (n *BgpPeer_UnicastIpv6PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer_UnicastIpv6Prefix {
	t.Helper()
	c := &oc.CollectionBgpPeer_UnicastIpv6Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer_UnicastIpv6Prefix) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv6Prefix) bool) *oc.BgpPeer_UnicastIpv6PrefixWatcher {
	t.Helper()
	w := &oc.BgpPeer_UnicastIpv6PrefixWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv6Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedBgpPeer_UnicastIpv6Prefix{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer_UnicastIpv6Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv6Prefix) bool) *oc.BgpPeer_UnicastIpv6PrefixWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix to the batch object.
func (n *BgpPeer_UnicastIpv6PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv6Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv6Prefix_AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv6Prefix_AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address with a ONCE subscription.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv6Prefix_AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_AddressPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv6Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv6Prefix_AddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/address to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv6Prefix_AddressPath extracts the value of the leaf Address from its parent oc.BgpPeer_UnicastIpv6Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertBgpPeer_UnicastIpv6Prefix_AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv6Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv6Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address with a ONCE subscription.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv6Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv4-address to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath extracts the value of the leaf NextHopIpv4Address from its parent oc.BgpPeer_UnicastIpv6Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertBgpPeer_UnicastIpv6Prefix_NextHopIpv4AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv6Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.NextHopIpv4Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv6Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address with a ONCE subscription.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv6Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/next-hop-ipv6-address to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath extracts the value of the leaf NextHopIpv6Address from its parent oc.BgpPeer_UnicastIpv6Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertBgpPeer_UnicastIpv6Prefix_NextHopIpv6AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv6Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.NextHopIpv6Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPath) Lookup(t testing.TB) *oc.QualifiedE_UnicastIpv6Prefix_Origin {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv6Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv6Prefix_OriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPath) Get(t testing.TB) oc.E_UnicastIpv6Prefix_Origin {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_UnicastIpv6Prefix_Origin {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_UnicastIpv6Prefix_Origin
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv6Prefix_OriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin with a ONCE subscription.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPathAny) Get(t testing.TB) []oc.E_UnicastIpv6Prefix_Origin {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_UnicastIpv6Prefix_Origin
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_UnicastIpv6Prefix_Origin {
	t.Helper()
	c := &oc.CollectionE_UnicastIpv6Prefix_Origin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_UnicastIpv6Prefix_Origin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_OriginPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_UnicastIpv6Prefix_Origin) bool) *oc.E_UnicastIpv6Prefix_OriginWatcher {
	t.Helper()
	w := &oc.E_UnicastIpv6Prefix_OriginWatcher{}
	gs := &oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv6Prefix_OriginPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_UnicastIpv6Prefix_Origin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_UnicastIpv6Prefix_Origin) bool) *oc.E_UnicastIpv6Prefix_OriginWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_OriginPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPath) Await(t testing.TB, timeout time.Duration, val oc.E_UnicastIpv6Prefix_Origin) *oc.QualifiedE_UnicastIpv6Prefix_Origin {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_UnicastIpv6Prefix_Origin) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_UnicastIpv6Prefix_Origin {
	t.Helper()
	c := &oc.CollectionE_UnicastIpv6Prefix_Origin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_UnicastIpv6Prefix_Origin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_OriginPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_UnicastIpv6Prefix_Origin) bool) *oc.E_UnicastIpv6Prefix_OriginWatcher {
	t.Helper()
	w := &oc.E_UnicastIpv6Prefix_OriginWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv6Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv6Prefix_OriginPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_UnicastIpv6Prefix_Origin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_UnicastIpv6Prefix_Origin) bool) *oc.E_UnicastIpv6Prefix_OriginWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_OriginPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/origin to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_OriginPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv6Prefix_OriginPath extracts the value of the leaf Origin from its parent oc.BgpPeer_UnicastIpv6Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedE_UnicastIpv6Prefix_Origin.
func convertBgpPeer_UnicastIpv6Prefix_OriginPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv6Prefix) *oc.QualifiedE_UnicastIpv6Prefix_Origin {
	t.Helper()
	qv := &oc.QualifiedE_UnicastIpv6Prefix_Origin{
		Metadata: md,
	}
	val := parent.Origin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv6Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv6Prefix_PathIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv6Prefix_PathIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id with a ONCE subscription.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_PathIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv6Prefix_PathIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_PathIdPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_PathIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv6Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv6Prefix_PathIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_PathIdPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/path-id to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_PathIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv6Prefix_PathIdPath extracts the value of the leaf PathId from its parent oc.BgpPeer_UnicastIpv6Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertBgpPeer_UnicastIpv6Prefix_PathIdPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv6Prefix) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.PathId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv6Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv6Prefix_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv6Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv6Prefix_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length with a ONCE subscription.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_PrefixLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv6Prefix_PrefixLengthPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_PrefixLengthPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv6Prefix_PrefixLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv6Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv6Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv6Prefix", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv6Prefix_PrefixLengthPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv6Prefix_PrefixLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv6-prefixes/unicast-ipv6-prefix/state/prefix-length to the batch object.
func (n *BgpPeer_UnicastIpv6Prefix_PrefixLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv6Prefix_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.BgpPeer_UnicastIpv6Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertBgpPeer_UnicastIpv6Prefix_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv6Prefix) *oc.QualifiedUint32 {
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
