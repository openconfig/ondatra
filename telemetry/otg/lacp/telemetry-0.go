package lacp

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

// Lookup fetches the value at /open-traffic-generator-lacp/lacp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LacpPath) Lookup(t testing.TB) *oc.QualifiedLacp {
	t.Helper()
	goStruct := &oc.Lacp{}
	md, ok := oc.Lookup(t, n, "Lacp", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLacp{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LacpPath) Get(t testing.TB) *oc.Lacp {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LacpPathAny) Lookup(t testing.TB) []*oc.QualifiedLacp {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLacp
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLacp{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp with a ONCE subscription.
func (n *LacpPathAny) Get(t testing.TB) []*oc.Lacp {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lacp
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LacpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp {
	t.Helper()
	c := &oc.CollectionLacp{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLacp{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lacp)))
		return false
	})
	return c
}

func watch_LacpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp) bool) *oc.LacpWatcher {
	t.Helper()
	w := &oc.LacpWatcher{}
	gs := &oc.Lacp{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp", gs, queryPath, false, false)
		qv := (&oc.QualifiedLacp{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LacpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp) bool) *oc.LacpWatcher {
	t.Helper()
	return watch_LacpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LacpPath) Await(t testing.TB, timeout time.Duration, val *oc.Lacp) *oc.QualifiedLacp {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLacp) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp to the batch object.
func (n *LacpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LacpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp {
	t.Helper()
	c := &oc.CollectionLacp{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LacpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp) bool) *oc.LacpWatcher {
	t.Helper()
	w := &oc.LacpWatcher{}
	structs := map[string]*oc.Lacp{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLacp{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LacpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp) bool) *oc.LacpWatcher {
	t.Helper()
	return watch_LacpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp to the batch object.
func (n *LacpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMemberPath) Lookup(t testing.TB) *oc.QualifiedLacp_LagMember {
	t.Helper()
	goStruct := &oc.Lacp_LagMember{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLacp_LagMember{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMemberPath) Get(t testing.TB) *oc.Lacp_LagMember {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMemberPathAny) Lookup(t testing.TB) []*oc.QualifiedLacp_LagMember {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLacp_LagMember
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLacp_LagMember{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member with a ONCE subscription.
func (n *Lacp_LagMemberPathAny) Get(t testing.TB) []*oc.Lacp_LagMember {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lacp_LagMember
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMemberPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp_LagMember {
	t.Helper()
	c := &oc.CollectionLacp_LagMember{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp_LagMember) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLacp_LagMember{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lacp_LagMember)))
		return false
	})
	return c
}

func watch_Lacp_LagMemberPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp_LagMember) bool) *oc.Lacp_LagMemberWatcher {
	t.Helper()
	w := &oc.Lacp_LagMemberWatcher{}
	gs := &oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember", gs, queryPath, false, false)
		qv := (&oc.QualifiedLacp_LagMember{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp_LagMember)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMemberPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp_LagMember) bool) *oc.Lacp_LagMemberWatcher {
	t.Helper()
	return watch_Lacp_LagMemberPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMemberPath) Await(t testing.TB, timeout time.Duration, val *oc.Lacp_LagMember) *oc.QualifiedLacp_LagMember {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLacp_LagMember) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member to the batch object.
func (n *Lacp_LagMemberPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMemberPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp_LagMember {
	t.Helper()
	c := &oc.CollectionLacp_LagMember{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp_LagMember) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMemberPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp_LagMember) bool) *oc.Lacp_LagMemberWatcher {
	t.Helper()
	w := &oc.Lacp_LagMemberWatcher{}
	structs := map[string]*oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLacp_LagMember{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp_LagMember)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMemberPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp_LagMember) bool) *oc.Lacp_LagMemberWatcher {
	t.Helper()
	return watch_Lacp_LagMemberPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member to the batch object.
func (n *Lacp_LagMemberPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMember_ActivityPath) Lookup(t testing.TB) *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	goStruct := &oc.Lacp_LagMember{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember", goStruct, true, false)
	if ok {
		return convertLacp_LagMember_ActivityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMember_ActivityPath) Get(t testing.TB) oc.E_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMember_ActivityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_LagMember_ActivityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity with a ONCE subscription.
func (n *Lacp_LagMember_ActivityPathAny) Get(t testing.TB) []oc.E_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_OpenTrafficGeneratorLacp_LacpActivityType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_ActivityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	c := &oc.CollectionE_OpenTrafficGeneratorLacp_LacpActivityType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_ActivityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher {
	t.Helper()
	w := &oc.E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher{}
	gs := &oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_LagMember_ActivityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_ActivityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher {
	t.Helper()
	return watch_Lacp_LagMember_ActivityPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMember_ActivityPath) Await(t testing.TB, timeout time.Duration, val oc.E_OpenTrafficGeneratorLacp_LacpActivityType) *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity to the batch object.
func (n *Lacp_LagMember_ActivityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_ActivityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	c := &oc.CollectionE_OpenTrafficGeneratorLacp_LacpActivityType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_ActivityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher {
	t.Helper()
	w := &oc.E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher{}
	structs := map[string]*oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember", structs[pre], queryPath, true, false)
			qv := convertLacp_LagMember_ActivityPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_ActivityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) bool) *oc.E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher {
	t.Helper()
	return watch_Lacp_LagMember_ActivityPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/activity to the batch object.
func (n *Lacp_LagMember_ActivityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_LagMember_ActivityPath extracts the value of the leaf Activity from its parent oc.Lacp_LagMember
// and combines the update with an existing Metadata to return a *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType.
func convertLacp_LagMember_ActivityPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_LagMember) *oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	qv := &oc.QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType{
		Metadata: md,
	}
	val := parent.Activity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMember_AggregatablePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Lacp_LagMember{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember", goStruct, true, false)
	if ok {
		return convertLacp_LagMember_AggregatablePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMember_AggregatablePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMember_AggregatablePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_LagMember_AggregatablePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable with a ONCE subscription.
func (n *Lacp_LagMember_AggregatablePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_AggregatablePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_AggregatablePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_LagMember_AggregatablePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_AggregatablePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Lacp_LagMember_AggregatablePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMember_AggregatablePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable to the batch object.
func (n *Lacp_LagMember_AggregatablePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_AggregatablePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_AggregatablePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember", structs[pre], queryPath, true, false)
			qv := convertLacp_LagMember_AggregatablePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_AggregatablePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Lacp_LagMember_AggregatablePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/aggregatable to the batch object.
func (n *Lacp_LagMember_AggregatablePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_LagMember_AggregatablePath extracts the value of the leaf Aggregatable from its parent oc.Lacp_LagMember
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertLacp_LagMember_AggregatablePath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_LagMember) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Aggregatable
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMember_CollectingPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Lacp_LagMember{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember", goStruct, true, false)
	if ok {
		return convertLacp_LagMember_CollectingPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMember_CollectingPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMember_CollectingPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_LagMember_CollectingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting with a ONCE subscription.
func (n *Lacp_LagMember_CollectingPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_CollectingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_CollectingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_LagMember_CollectingPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_CollectingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Lacp_LagMember_CollectingPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMember_CollectingPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting to the batch object.
func (n *Lacp_LagMember_CollectingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_CollectingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_CollectingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Lacp_LagMember{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember", structs[pre], queryPath, true, false)
			qv := convertLacp_LagMember_CollectingPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_CollectingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Lacp_LagMember_CollectingPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/collecting to the batch object.
func (n *Lacp_LagMember_CollectingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_LagMember_CollectingPath extracts the value of the leaf Collecting from its parent oc.Lacp_LagMember
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertLacp_LagMember_CollectingPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_LagMember) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Collecting
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMember_CountersPath) Lookup(t testing.TB) *oc.QualifiedLacp_LagMember_Counters {
	t.Helper()
	goStruct := &oc.Lacp_LagMember_Counters{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLacp_LagMember_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMember_CountersPath) Get(t testing.TB) *oc.Lacp_LagMember_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMember_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedLacp_LagMember_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLacp_LagMember_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLacp_LagMember_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters with a ONCE subscription.
func (n *Lacp_LagMember_CountersPathAny) Get(t testing.TB) []*oc.Lacp_LagMember_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lacp_LagMember_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp_LagMember_Counters {
	t.Helper()
	c := &oc.CollectionLacp_LagMember_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp_LagMember_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLacp_LagMember_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Lacp_LagMember_Counters)))
		return false
	})
	return c
}

func watch_Lacp_LagMember_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp_LagMember_Counters) bool) *oc.Lacp_LagMember_CountersWatcher {
	t.Helper()
	w := &oc.Lacp_LagMember_CountersWatcher{}
	gs := &oc.Lacp_LagMember_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedLacp_LagMember_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp_LagMember_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp_LagMember_Counters) bool) *oc.Lacp_LagMember_CountersWatcher {
	t.Helper()
	return watch_Lacp_LagMember_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMember_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Lacp_LagMember_Counters) *oc.QualifiedLacp_LagMember_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLacp_LagMember_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters to the batch object.
func (n *Lacp_LagMember_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLacp_LagMember_Counters {
	t.Helper()
	c := &oc.CollectionLacp_LagMember_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLacp_LagMember_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLacp_LagMember_Counters) bool) *oc.Lacp_LagMember_CountersWatcher {
	t.Helper()
	w := &oc.Lacp_LagMember_CountersWatcher{}
	structs := map[string]*oc.Lacp_LagMember_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLacp_LagMember_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLacp_LagMember_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLacp_LagMember_Counters) bool) *oc.Lacp_LagMember_CountersWatcher {
	t.Helper()
	return watch_Lacp_LagMember_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters to the batch object.
func (n *Lacp_LagMember_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_LagMember_Counters_LacpInPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lacp_LagMember_Counters{}
	md, ok := oc.Lookup(t, n, "Lacp_LagMember_Counters", goStruct, true, false)
	if ok {
		return convertLacp_LagMember_Counters_LacpInPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_LagMember_Counters_LacpInPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_LagMember_Counters_LacpInPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_LagMember_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_LagMember_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLacp_LagMember_Counters_LacpInPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts with a ONCE subscription.
func (n *Lacp_LagMember_Counters_LacpInPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_Counters_LacpInPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_Counters_LacpInPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Lacp_LagMember_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lacp_LagMember_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLacp_LagMember_Counters_LacpInPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_Counters_LacpInPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_LagMember_Counters_LacpInPktsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lacp_LagMember_Counters_LacpInPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts to the batch object.
func (n *Lacp_LagMember_Counters_LacpInPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lacp_LagMember_Counters_LacpInPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lacp_LagMember_Counters_LacpInPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Lacp_LagMember_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lacp_LagMember_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lacp_LagMember_Counters", structs[pre], queryPath, true, false)
			qv := convertLacp_LagMember_Counters_LacpInPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lacp_LagMember_Counters_LacpInPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Lacp_LagMember_Counters_LacpInPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lacp/lacp/lag-members/lag-member/state/counters/lacp-in-pkts to the batch object.
func (n *Lacp_LagMember_Counters_LacpInPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLacp_LagMember_Counters_LacpInPktsPath extracts the value of the leaf LacpInPkts from its parent oc.Lacp_LagMember_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLacp_LagMember_Counters_LacpInPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_LagMember_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LacpInPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
