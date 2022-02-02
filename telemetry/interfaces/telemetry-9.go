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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
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
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-high-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath extracts the value of the leaf OuterHighVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterHighVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
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
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged-outer-range/state/outer-low-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath extracts the value of the leaf OuterLowVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange_OuterLowVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_DoubleTagged
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTagged{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_Match_DoubleTagged)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTaggedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_Match_DoubleTaggedWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTagged", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_Match_DoubleTagged) *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_DoubleTagged{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged) bool) *oc.Interface_Subinterface_Vlan_Match_DoubleTaggedWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTaggedPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTaggedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTagged", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/inner-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath extracts the value of the leaf InnerVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTagged
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTagged_InnerVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTagged) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTagged", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_DoubleTagged{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_DoubleTagged", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/double-tagged/state/outer-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath extracts the value of the leaf OuterVlanId from its parent oc.Interface_Subinterface_Vlan_Match_DoubleTagged
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_DoubleTagged_OuterVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_DoubleTagged) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTaggedList", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedList", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_SingleTaggedList
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_SingleTaggedList{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_Match_SingleTaggedList)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_SingleTaggedListPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList) bool) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedListWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedListWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedList{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedList", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList) bool) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedListWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedListPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_Match_SingleTaggedList) *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_SingleTaggedList{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList) bool) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedListWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedListPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedListPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedList{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTaggedList", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedList{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedList", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	w := &oc.Uint16SliceWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedList{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedList", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Await(t testing.TB, timeout time.Duration, val []uint16) *oc.QualifiedUint16Slice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16Slice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-list/state/vlan-ids to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath extracts the value of the leaf VlanIds from its parent oc.Interface_Subinterface_Vlan_Match_SingleTaggedList
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertInterface_Subinterface_Vlan_Match_SingleTaggedList_VlanIdsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_SingleTaggedList) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.VlanIds
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTagged{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTagged", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTagged{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTagged", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_SingleTagged
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_SingleTagged{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_Match_SingleTagged)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_SingleTaggedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged) bool) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_SingleTagged{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTagged", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged) bool) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_Match_SingleTagged) *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_SingleTagged{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTagged) bool) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_SingleTaggedRange{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange)))
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_SingleTaggedRangePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange) bool) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRangeWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRangeWatcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedRange", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange) bool) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRangeWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedRangePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange) *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_Match_SingleTaggedRange{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange) bool) *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRangeWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedRangePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRangePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/high-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath extracts the value of the leaf HighVlanId from its parent oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_HighVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.HighVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedRange", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTaggedRange", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged-range/state/low-vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath extracts the value of the leaf LowVlanId from its parent oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_SingleTaggedRange_LowVlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_SingleTaggedRange) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.LowVlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTagged{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_Match_SingleTagged", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_Match_SingleTagged{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTagged", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_Subinterface_Vlan_Match_SingleTagged{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan_Match_SingleTagged", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/match/single-tagged/state/vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_Match_SingleTagged_VlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath extracts the value of the leaf VlanId from its parent oc.Interface_Subinterface_Vlan_Match_SingleTagged
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_Match_SingleTagged_VlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_Match_SingleTagged) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.VlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan", goStruct, true, false)
	if ok {
		return convertInterface_Subinterface_Vlan_VlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Get(t testing.TB) oc.Interface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_VlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_VlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_VlanIdPathAny) Get(t testing.TB) []oc.Interface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Interface_Subinterface_Vlan_VlanId_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_VlanId_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Subinterface_Vlan_VlanIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union) bool) *oc.Interface_Subinterface_Vlan_VlanId_UnionWatcher {
	t.Helper()
	w := &oc.Interface_Subinterface_Vlan_VlanId_UnionWatcher{}
	gs := &oc.Interface_Subinterface_Vlan{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Subinterface_Vlan", gs, queryPath, true, false)
		return convertInterface_Subinterface_Vlan_VlanIdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union) bool) *oc.Interface_Subinterface_Vlan_VlanId_UnionWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_VlanIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Await(t testing.TB, timeout time.Duration, val oc.Interface_Subinterface_Vlan_VlanId_Union) *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_VlanIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Subinterface_Vlan_VlanIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	c := &oc.CollectionInterface_Subinterface_Vlan_VlanId_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Subinterface_Vlan_VlanIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union) bool) *oc.Interface_Subinterface_Vlan_VlanId_UnionWatcher {
	t.Helper()
	return watch_Interface_Subinterface_Vlan_VlanIdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/state/vlan-id to the batch object.
func (n *Interface_Subinterface_Vlan_VlanIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Subinterface_Vlan_VlanIdPath extracts the value of the leaf VlanId from its parent oc.Interface_Subinterface_Vlan
// and combines the update with an existing Metadata to return a *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union.
func convertInterface_Subinterface_Vlan_VlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan) *oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	qv := &oc.QualifiedInterface_Subinterface_Vlan_VlanId_Union{
		Metadata: md,
	}
	val := parent.VlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/tpid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_TpidPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_TpidPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_VlanTypes_TPID_TYPES{
		Metadata: md,
	}).SetVal(goStruct.GetTpid())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/tpid with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_TpidPath) Get(t testing.TB) oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/tpid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_TpidPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_TPID_TYPES
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_TpidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/tpid with a ONCE subscription.
func (n *Interface_TpidPathAny) Get(t testing.TB) []oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_TPID_TYPES
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/tpid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_TpidPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_VlanTypes_TPID_TYPES {
	t.Helper()
	c := &oc.CollectionE_VlanTypes_TPID_TYPES{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_VlanTypes_TPID_TYPES) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_TpidPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_TPID_TYPES) bool) *oc.E_VlanTypes_TPID_TYPESWatcher {
	t.Helper()
	w := &oc.E_VlanTypes_TPID_TYPESWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return convertInterface_TpidPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_VlanTypes_TPID_TYPES)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/tpid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_TpidPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_TPID_TYPES) bool) *oc.E_VlanTypes_TPID_TYPESWatcher {
	t.Helper()
	return watch_Interface_TpidPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/tpid with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_TpidPath) Await(t testing.TB, timeout time.Duration, val oc.E_VlanTypes_TPID_TYPES) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_VlanTypes_TPID_TYPES) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/tpid failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/tpid to the batch object.
func (n *Interface_TpidPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/tpid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_TpidPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_VlanTypes_TPID_TYPES {
	t.Helper()
	c := &oc.CollectionE_VlanTypes_TPID_TYPES{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_VlanTypes_TPID_TYPES) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/tpid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_TpidPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_VlanTypes_TPID_TYPES) bool) *oc.E_VlanTypes_TPID_TYPESWatcher {
	t.Helper()
	return watch_Interface_TpidPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/tpid to the batch object.
func (n *Interface_TpidPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_TpidPath extracts the value of the leaf Tpid from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_TPID_TYPES.
func convertInterface_TpidPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_TPID_TYPES{
		Metadata: md,
	}
	val := parent.Tpid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/transceiver with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_TransceiverPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_TransceiverPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/transceiver with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_TransceiverPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/transceiver with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_TransceiverPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_TransceiverPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/transceiver with a ONCE subscription.
func (n *Interface_TransceiverPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/transceiver with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_TransceiverPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_TransceiverPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return convertInterface_TransceiverPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/transceiver with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_TransceiverPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_TransceiverPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/transceiver with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_TransceiverPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/transceiver failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/transceiver to the batch object.
func (n *Interface_TransceiverPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/transceiver with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_TransceiverPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/transceiver with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_TransceiverPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_TransceiverPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/transceiver to the batch object.
func (n *Interface_TransceiverPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_TransceiverPath extracts the value of the leaf Transceiver from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_TransceiverPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Transceiver
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/state/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_TypePath) Lookup(t testing.TB) *oc.QualifiedE_IETFInterfaces_InterfaceType {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, true, false)
	if ok {
		return convertInterface_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/state/type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_TypePath) Get(t testing.TB) oc.E_IETFInterfaces_InterfaceType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/state/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_IETFInterfaces_InterfaceType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IETFInterfaces_InterfaceType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/state/type with a ONCE subscription.
func (n *Interface_TypePathAny) Get(t testing.TB) []oc.E_IETFInterfaces_InterfaceType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IETFInterfaces_InterfaceType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_TypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IETFInterfaces_InterfaceType {
	t.Helper()
	c := &oc.CollectionE_IETFInterfaces_InterfaceType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IETFInterfaces_InterfaceType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_TypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IETFInterfaces_InterfaceType) bool) *oc.E_IETFInterfaces_InterfaceTypeWatcher {
	t.Helper()
	w := &oc.E_IETFInterfaces_InterfaceTypeWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, true, false)
		return convertInterface_TypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IETFInterfaces_InterfaceType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_TypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IETFInterfaces_InterfaceType) bool) *oc.E_IETFInterfaces_InterfaceTypeWatcher {
	t.Helper()
	return watch_Interface_TypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/state/type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_TypePath) Await(t testing.TB, timeout time.Duration, val oc.E_IETFInterfaces_InterfaceType) *oc.QualifiedE_IETFInterfaces_InterfaceType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IETFInterfaces_InterfaceType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/state/type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/type to the batch object.
func (n *Interface_TypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_TypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IETFInterfaces_InterfaceType {
	t.Helper()
	c := &oc.CollectionE_IETFInterfaces_InterfaceType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IETFInterfaces_InterfaceType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IETFInterfaces_InterfaceType) bool) *oc.E_IETFInterfaces_InterfaceTypeWatcher {
	t.Helper()
	return watch_Interface_TypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/state/type to the batch object.
func (n *Interface_TypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_TypePath extracts the value of the leaf Type from its parent oc.Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IETFInterfaces_InterfaceType.
func convertInterface_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface) *oc.QualifiedE_IETFInterfaces_InterfaceType {
	t.Helper()
	qv := &oc.QualifiedE_IETFInterfaces_InterfaceType{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
