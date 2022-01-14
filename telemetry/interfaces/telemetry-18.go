package interfaces

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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_MatchPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_MatchPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_MatchPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_MatchPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_MatchPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_Match{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_Match)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_MatchPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match) bool) *oc.Interface_Subinterface_Vlan_MatchWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_MatchWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_Match)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_MatchPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match) bool) *oc.Interface_Subinterface_Vlan_MatchWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_MatchPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_MatchPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_Match) *oc.QualifiedInterface_Subinterface_Vlan_Match {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_Match) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match to the batch object.
func (n *Interface_Subinterface_Vlan_MatchPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_MatchPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_MatchPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match) bool) *oc.Interface_Subinterface_Vlan_MatchWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_MatchPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match to the batch object.
func (n *Interface_Subinterface_Vlan_MatchPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	w := &oc.Uint16SliceWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Await(t testing.TB, timeout time.Duration, val []uint16) *oc.QualifiedUint16Slice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16Slice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/inner-vlan-ids to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath extracts the value of the leaf InnerVlanIds from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_InnerVlanIdsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.InnerVlanIds
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-list/state/outer-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath extracts the value of the leaf OuterVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList_OuterVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.OuterVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangeWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangeWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangeWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangeWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-high-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath extracts the value of the leaf InnerHighVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerHighVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerHighVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/inner-low-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath extracts the value of the leaf InnerLowVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_InnerLowVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerLowVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-high-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath extracts the value of the leaf OuterHighVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterHighVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.OuterHighVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-outer-range/state/outer-low-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath extracts the value of the leaf OuterLowVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange_OuterLowVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.OuterLowVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangeWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangeWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangeWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangeWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-high-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath extracts the value of the leaf InnerHighVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerHighVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerHighVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/inner-low-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath extracts the value of the leaf InnerLowVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_InnerLowVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerLowVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	w := &oc.Uint16SliceWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Await(t testing.TB, timeout time.Duration, val []uint16) *oc.QualifiedUint16Slice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16Slice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-inner-range/state/outer-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath extracts the value of the leaf OuterVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange_OuterVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.OuterVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/inner-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath extracts the value of the leaf InnerVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_InnerVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	w := &oc.Uint16SliceWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Await(t testing.TB, timeout time.Duration, val []uint16) *oc.QualifiedUint16Slice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16Slice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-list/state/outer-vlan-ids to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath extracts the value of the leaf OuterVlanIds from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList_OuterVlanIdsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.OuterVlanIds
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangeWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangeWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangeWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangeWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/inner-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath extracts the value of the leaf InnerVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_InnerVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.InnerVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
