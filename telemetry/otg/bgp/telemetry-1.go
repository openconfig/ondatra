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

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4PrefixPath) Lookup(t testing.TB) *oc.QualifiedBgpPeer_UnicastIpv4Prefix {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix", goStruct, false, false)
	if ok {
		return (&oc.QualifiedBgpPeer_UnicastIpv4Prefix{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4PrefixPath) Get(t testing.TB) *oc.BgpPeer_UnicastIpv4Prefix {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedBgpPeer_UnicastIpv4Prefix {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBgpPeer_UnicastIpv4Prefix
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedBgpPeer_UnicastIpv4Prefix{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4PrefixPathAny) Get(t testing.TB) []*oc.BgpPeer_UnicastIpv4Prefix {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.BgpPeer_UnicastIpv4Prefix
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer_UnicastIpv4Prefix {
	t.Helper()
	c := &oc.CollectionBgpPeer_UnicastIpv4Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer_UnicastIpv4Prefix) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedBgpPeer_UnicastIpv4Prefix{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.BgpPeer_UnicastIpv4Prefix)))
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix) bool) *oc.BgpPeer_UnicastIpv4PrefixWatcher {
	t.Helper()
	w := &oc.BgpPeer_UnicastIpv4PrefixWatcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", gs, queryPath, false, false)
		qv := (&oc.QualifiedBgpPeer_UnicastIpv4Prefix{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer_UnicastIpv4Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix) bool) *oc.BgpPeer_UnicastIpv4PrefixWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4PrefixPath) Await(t testing.TB, timeout time.Duration, val *oc.BgpPeer_UnicastIpv4Prefix) *oc.QualifiedBgpPeer_UnicastIpv4Prefix {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBgpPeer_UnicastIpv4Prefix) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix to the batch object.
func (n *BgpPeer_UnicastIpv4PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer_UnicastIpv4Prefix {
	t.Helper()
	c := &oc.CollectionBgpPeer_UnicastIpv4Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer_UnicastIpv4Prefix) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix) bool) *oc.BgpPeer_UnicastIpv4PrefixWatcher {
	t.Helper()
	w := &oc.BgpPeer_UnicastIpv4PrefixWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv4Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedBgpPeer_UnicastIpv4Prefix{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer_UnicastIpv4Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix) bool) *oc.BgpPeer_UnicastIpv4PrefixWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix to the batch object.
func (n *BgpPeer_UnicastIpv4PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_AddressPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertBgpPeer_UnicastIpv4Prefix_AddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/address to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_AddressPath extracts the value of the leaf Address from its parent oc.BgpPeer_UnicastIpv4Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertBgpPeer_UnicastIpv4Prefix_AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix) *oc.QualifiedString {
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

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPath) Lookup(t testing.TB) *oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix_AsPath", goStruct, false, false)
	if ok {
		return (&oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPath) Get(t testing.TB) *oc.BgpPeer_UnicastIpv4Prefix_AsPath {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPathAny) Lookup(t testing.TB) []*oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_AsPath", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPathAny) Get(t testing.TB) []*oc.BgpPeer_UnicastIpv4Prefix_AsPath {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.BgpPeer_UnicastIpv4Prefix_AsPath
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer_UnicastIpv4Prefix_AsPath {
	t.Helper()
	c := &oc.CollectionBgpPeer_UnicastIpv4Prefix_AsPath{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.BgpPeer_UnicastIpv4Prefix_AsPath)))
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_AsPathPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) bool) *oc.BgpPeer_UnicastIpv4Prefix_AsPathWatcher {
	t.Helper()
	w := &oc.BgpPeer_UnicastIpv4Prefix_AsPathWatcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_AsPath", gs, queryPath, false, false)
		qv := (&oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) bool) *oc.BgpPeer_UnicastIpv4Prefix_AsPathWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_AsPathPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPath) Await(t testing.TB, timeout time.Duration, val *oc.BgpPeer_UnicastIpv4Prefix_AsPath) *oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer_UnicastIpv4Prefix_AsPath {
	t.Helper()
	c := &oc.CollectionBgpPeer_UnicastIpv4Prefix_AsPath{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_AsPathPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) bool) *oc.BgpPeer_UnicastIpv4Prefix_AsPathWatcher {
	t.Helper()
	w := &oc.BgpPeer_UnicastIpv4Prefix_AsPathWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_AsPath", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) bool) *oc.BgpPeer_UnicastIpv4Prefix_AsPathWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_AsPathPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_AsPathPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath) Lookup(t testing.TB) *oc.QualifiedUint32Slice {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix_AsPath", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath) Get(t testing.TB) []uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_AsPath", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPathAny) Get(t testing.TB) [][]uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32Slice {
	t.Helper()
	c := &oc.CollectionUint32Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32Slice) bool) *oc.Uint32SliceWatcher {
	t.Helper()
	w := &oc.Uint32SliceWatcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_AsPath", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32Slice) bool) *oc.Uint32SliceWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath) Await(t testing.TB, timeout time.Duration, val []uint32) *oc.QualifiedUint32Slice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32Slice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32Slice {
	t.Helper()
	c := &oc.CollectionUint32Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32Slice) bool) *oc.Uint32SliceWatcher {
	t.Helper()
	w := &oc.Uint32SliceWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_AsPath", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32Slice) bool) *oc.Uint32SliceWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/as_numbers to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath extracts the value of the leaf AsNumbers from its parent oc.BgpPeer_UnicastIpv4Prefix_AsPath
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32Slice.
func convertBgpPeer_UnicastIpv4Prefix_AsPath_AsNumbersPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix_AsPath) *oc.QualifiedUint32Slice {
	t.Helper()
	qv := &oc.QualifiedUint32Slice{
		Metadata: md,
	}
	val := parent.AsNumbers
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath) Lookup(t testing.TB) *oc.QualifiedE_State_SegmentType {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix_AsPath", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath) Get(t testing.TB) oc.E_State_SegmentType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_State_SegmentType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_State_SegmentType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_AsPath", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePathAny) Get(t testing.TB) []oc.E_State_SegmentType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_State_SegmentType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_State_SegmentType {
	t.Helper()
	c := &oc.CollectionE_State_SegmentType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_State_SegmentType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_State_SegmentType) bool) *oc.E_State_SegmentTypeWatcher {
	t.Helper()
	w := &oc.E_State_SegmentTypeWatcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_AsPath", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_State_SegmentType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_State_SegmentType) bool) *oc.E_State_SegmentTypeWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_State_SegmentType) *oc.QualifiedE_State_SegmentType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_State_SegmentType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_State_SegmentType {
	t.Helper()
	c := &oc.CollectionE_State_SegmentType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_State_SegmentType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_State_SegmentType) bool) *oc.E_State_SegmentTypeWatcher {
	t.Helper()
	w := &oc.E_State_SegmentTypeWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv4Prefix_AsPath{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_AsPath", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_State_SegmentType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_State_SegmentType) bool) *oc.E_State_SegmentTypeWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/as-path/segment-type to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath extracts the value of the leaf SegmentType from its parent oc.BgpPeer_UnicastIpv4Prefix_AsPath
// and combines the update with an existing Metadata to return a *oc.QualifiedE_State_SegmentType.
func convertBgpPeer_UnicastIpv4Prefix_AsPath_SegmentTypePath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix_AsPath) *oc.QualifiedE_State_SegmentType {
	t.Helper()
	qv := &oc.QualifiedE_State_SegmentType{
		Metadata: md,
	}
	val := parent.SegmentType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPath) Lookup(t testing.TB) *oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix_Community", goStruct, false, false)
	if ok {
		return (&oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPath) Get(t testing.TB) *oc.BgpPeer_UnicastIpv4Prefix_Community {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPathAny) Lookup(t testing.TB) []*oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPathAny) Get(t testing.TB) []*oc.BgpPeer_UnicastIpv4Prefix_Community {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.BgpPeer_UnicastIpv4Prefix_Community
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer_UnicastIpv4Prefix_Community {
	t.Helper()
	c := &oc.CollectionBgpPeer_UnicastIpv4Prefix_Community{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.BgpPeer_UnicastIpv4Prefix_Community)))
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_CommunityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community) bool) *oc.BgpPeer_UnicastIpv4Prefix_CommunityWatcher {
	t.Helper()
	w := &oc.BgpPeer_UnicastIpv4Prefix_CommunityWatcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", gs, queryPath, false, false)
		qv := (&oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community) bool) *oc.BgpPeer_UnicastIpv4Prefix_CommunityWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_CommunityPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPath) Await(t testing.TB, timeout time.Duration, val *oc.BgpPeer_UnicastIpv4Prefix_Community) *oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBgpPeer_UnicastIpv4Prefix_Community {
	t.Helper()
	c := &oc.CollectionBgpPeer_UnicastIpv4Prefix_Community{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_CommunityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community) bool) *oc.BgpPeer_UnicastIpv4Prefix_CommunityWatcher {
	t.Helper()
	w := &oc.BgpPeer_UnicastIpv4Prefix_CommunityWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv4Prefix_Community{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv4Prefix_Community{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBgpPeer_UnicastIpv4Prefix_Community) bool) *oc.BgpPeer_UnicastIpv4Prefix_CommunityWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_CommunityPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_CommunityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath) Lookup(t testing.TB) *oc.QualifiedE_State_CommunityType {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix_Community", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath) Get(t testing.TB) oc.E_State_CommunityType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_State_CommunityType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_State_CommunityType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePathAny) Get(t testing.TB) []oc.E_State_CommunityType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_State_CommunityType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_State_CommunityType {
	t.Helper()
	c := &oc.CollectionE_State_CommunityType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_State_CommunityType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_State_CommunityType) bool) *oc.E_State_CommunityTypeWatcher {
	t.Helper()
	w := &oc.E_State_CommunityTypeWatcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_State_CommunityType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_State_CommunityType) bool) *oc.E_State_CommunityTypeWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_State_CommunityType) *oc.QualifiedE_State_CommunityType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_State_CommunityType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_State_CommunityType {
	t.Helper()
	c := &oc.CollectionE_State_CommunityType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_State_CommunityType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_State_CommunityType) bool) *oc.E_State_CommunityTypeWatcher {
	t.Helper()
	w := &oc.E_State_CommunityTypeWatcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv4Prefix_Community{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv4Prefix_Community{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_State_CommunityType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_State_CommunityType) bool) *oc.E_State_CommunityTypeWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/community-type to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CommunityTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath extracts the value of the leaf CommunityType from its parent oc.BgpPeer_UnicastIpv4Prefix_Community
// and combines the update with an existing Metadata to return a *oc.QualifiedE_State_CommunityType.
func convertBgpPeer_UnicastIpv4Prefix_Community_CommunityTypePath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix_Community) *oc.QualifiedE_State_CommunityType {
	t.Helper()
	qv := &oc.QualifiedE_State_CommunityType{
		Metadata: md,
	}
	val := parent.CommunityType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix_Community", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv4Prefix_Community{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv4Prefix_Community{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-number to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath extracts the value of the leaf CustomAsNumber from its parent oc.BgpPeer_UnicastIpv4Prefix_Community
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsNumberPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix_Community) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.CustomAsNumber
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix_Community", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix_Community{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.BgpPeer_UnicastIpv4Prefix_Community{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.BgpPeer_UnicastIpv4Prefix_Community{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix_Community", structs[pre], queryPath, true, false)
			qv := convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/community/custom-as-value to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath extracts the value of the leaf CustomAsValue from its parent oc.BgpPeer_UnicastIpv4Prefix_Community
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertBgpPeer_UnicastIpv4Prefix_Community_CustomAsValuePath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix_Community) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.CustomAsValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertBgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv4-address to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath extracts the value of the leaf NextHopIpv4Address from its parent oc.BgpPeer_UnicastIpv4Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertBgpPeer_UnicastIpv4Prefix_NextHopIpv4AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix) *oc.QualifiedString {
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

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertBgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/next-hop-ipv6-address to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath extracts the value of the leaf NextHopIpv6Address from its parent oc.BgpPeer_UnicastIpv4Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertBgpPeer_UnicastIpv4Prefix_NextHopIpv6AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix) *oc.QualifiedString {
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

// Lookup fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPath) Lookup(t testing.TB) *oc.QualifiedE_UnicastIpv4Prefix_Origin {
	t.Helper()
	goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
	md, ok := oc.Lookup(t, n, "BgpPeer_UnicastIpv4Prefix", goStruct, true, false)
	if ok {
		return convertBgpPeer_UnicastIpv4Prefix_OriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPath) Get(t testing.TB) oc.E_UnicastIpv4Prefix_Origin {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_UnicastIpv4Prefix_Origin {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_UnicastIpv4Prefix_Origin
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.BgpPeer_UnicastIpv4Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertBgpPeer_UnicastIpv4Prefix_OriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin with a ONCE subscription.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPathAny) Get(t testing.TB) []oc.E_UnicastIpv4Prefix_Origin {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_UnicastIpv4Prefix_Origin
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_UnicastIpv4Prefix_Origin {
	t.Helper()
	c := &oc.CollectionE_UnicastIpv4Prefix_Origin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_UnicastIpv4Prefix_Origin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_OriginPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_UnicastIpv4Prefix_Origin) bool) *oc.E_UnicastIpv4Prefix_OriginWatcher {
	t.Helper()
	w := &oc.E_UnicastIpv4Prefix_OriginWatcher{}
	gs := &oc.BgpPeer_UnicastIpv4Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "BgpPeer_UnicastIpv4Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertBgpPeer_UnicastIpv4Prefix_OriginPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_UnicastIpv4Prefix_Origin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_UnicastIpv4Prefix_Origin) bool) *oc.E_UnicastIpv4Prefix_OriginWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_OriginPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPath) Await(t testing.TB, timeout time.Duration, val oc.E_UnicastIpv4Prefix_Origin) *oc.QualifiedE_UnicastIpv4Prefix_Origin {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_UnicastIpv4Prefix_Origin) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_UnicastIpv4Prefix_Origin {
	t.Helper()
	c := &oc.CollectionE_UnicastIpv4Prefix_Origin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_UnicastIpv4Prefix_Origin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_BgpPeer_UnicastIpv4Prefix_OriginPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_UnicastIpv4Prefix_Origin) bool) *oc.E_UnicastIpv4Prefix_OriginWatcher {
	t.Helper()
	w := &oc.E_UnicastIpv4Prefix_OriginWatcher{}
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
			qv := convertBgpPeer_UnicastIpv4Prefix_OriginPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_UnicastIpv4Prefix_Origin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_UnicastIpv4Prefix_Origin) bool) *oc.E_UnicastIpv4Prefix_OriginWatcher {
	t.Helper()
	return watch_BgpPeer_UnicastIpv4Prefix_OriginPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-bgp/bgp-peers/bgp-peer/unicast-ipv4-prefixes/unicast-ipv4-prefix/state/origin to the batch object.
func (n *BgpPeer_UnicastIpv4Prefix_OriginPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertBgpPeer_UnicastIpv4Prefix_OriginPath extracts the value of the leaf Origin from its parent oc.BgpPeer_UnicastIpv4Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedE_UnicastIpv4Prefix_Origin.
func convertBgpPeer_UnicastIpv4Prefix_OriginPath(t testing.TB, md *genutil.Metadata, parent *oc.BgpPeer_UnicastIpv4Prefix) *oc.QualifiedE_UnicastIpv4Prefix_Origin {
	t.Helper()
	qv := &oc.QualifiedE_UnicastIpv4Prefix_Origin{
		Metadata: md,
	}
	val := parent.Origin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

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
