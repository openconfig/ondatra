package lacp

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

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_MemberPath) Lookup(t testing.TB) *oc.QualifiedLacp_Interface_Member {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLacp_Interface_Member{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_MemberPath) Get(t testing.TB) *oc.Lacp_Interface_Member {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_MemberPathAny) Lookup(t testing.TB) []*oc.QualifiedLacp_Interface_Member {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLacp_Interface_Member
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface_Member{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface_Member", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLacp_Interface_Member{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member with a ONCE subscription.
func (n *Lacp_Interface_MemberPathAny) Get(t testing.TB) []*oc.Lacp_Interface_Member {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lacp_Interface_Member
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_MemberPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp_Interface_Member {
	t.Helper()
	c := &oc.CollectionLacp_Interface_Member{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp_Interface_Member) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLacp_Interface_Member{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lacp_Interface_Member)))
		return false
	})
	return c
}

func watch_Lacp_Interface_MemberPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp_Interface_Member) bool) *oc.Lacp_Interface_MemberWatcher {
	t.Helper()
	w := &oc.Lacp_Interface_MemberWatcher{}
	gs := &oc.Lacp_Interface_Member{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member", gs, queryPath, false, false)
		return (&oc.QualifiedLacp_Interface_Member{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp_Interface_Member)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_MemberPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp_Interface_Member) bool) *oc.Lacp_Interface_MemberWatcher {
	t.Helper()
	return watch_Lacp_Interface_MemberPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_MemberPath) Await(t testing.TB, timeout time.Duration, val *oc.Lacp_Interface_Member) *oc.QualifiedLacp_Interface_Member {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLacp_Interface_Member) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member to the batch object.
func (n *Lacp_Interface_MemberPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_MemberPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp_Interface_Member {
	t.Helper()
	c := &oc.CollectionLacp_Interface_Member{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp_Interface_Member) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_MemberPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp_Interface_Member) bool) *oc.Lacp_Interface_MemberWatcher {
	t.Helper()
	return watch_Lacp_Interface_MemberPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member to the batch object.
func (n *Lacp_Interface_MemberPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_Member_ActivityPath) Lookup(t testing.TB) *oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	goStruct := &oc.Lacp_Interface_Member{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface_Member", goStruct, true, false)
	if ok {
		return convertLacp_Interface_Member_ActivityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_Member_ActivityPath) Get(t testing.TB) oc.E_Lacp_LacpActivityType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_Member_ActivityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Lacp_LacpActivityType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface_Member{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface_Member", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_Member_ActivityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity with a ONCE subscription.
func (n *Lacp_Interface_Member_ActivityPathAny) Get(t testing.TB) []oc.E_Lacp_LacpActivityType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Lacp_LacpActivityType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_ActivityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lacp_LacpActivityType {
	t.Helper()
	c := &oc.CollectionE_Lacp_LacpActivityType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lacp_LacpActivityType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_Interface_Member_ActivityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpActivityType) bool) *oc.E_Lacp_LacpActivityTypeWatcher {
	t.Helper()
	w := &oc.E_Lacp_LacpActivityTypeWatcher{}
	gs := &oc.Lacp_Interface_Member{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_Interface_Member", gs, queryPath, true, false)
		return convertLacp_Interface_Member_ActivityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lacp_LacpActivityType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_ActivityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpActivityType) bool) *oc.E_Lacp_LacpActivityTypeWatcher {
	t.Helper()
	return watch_Lacp_Interface_Member_ActivityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_Interface_Member_ActivityPath) Await(t testing.TB, timeout time.Duration, val oc.E_Lacp_LacpActivityType) *oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Lacp_LacpActivityType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity to the batch object.
func (n *Lacp_Interface_Member_ActivityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_Interface_Member_ActivityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lacp_LacpActivityType {
	t.Helper()
	c := &oc.CollectionE_Lacp_LacpActivityType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lacp_LacpActivityType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_Interface_Member_ActivityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lacp_LacpActivityType) bool) *oc.E_Lacp_LacpActivityTypeWatcher {
	t.Helper()
	return watch_Lacp_Interface_Member_ActivityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-lacp/lacp/interfaces/interface/members/member/state/activity to the batch object.
func (n *Lacp_Interface_Member_ActivityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_Interface_Member_ActivityPath extracts the value of the leaf Activity from its parent oc.Lacp_Interface_Member
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Lacp_LacpActivityType.
func convertLacp_Interface_Member_ActivityPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface_Member) *oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	qv := &oc.QualifiedE_Lacp_LacpActivityType{
		Metadata: md,
	}
	val := parent.Activity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
