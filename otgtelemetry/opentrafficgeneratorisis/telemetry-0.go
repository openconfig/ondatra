package opentrafficgeneratorisis

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

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouterPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter {
	t.Helper()
	goStruct := &oc.IsisRouter{}
	md, ok := oc.Lookup(t, n, "IsisRouter", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouterPath) Get(t testing.TB) *oc.IsisRouter {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouterPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router with a ONCE subscription.
func (n *IsisRouterPathAny) Get(t testing.TB) []*oc.IsisRouter {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouterPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter {
	t.Helper()
	c := &oc.CollectionIsisRouter{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter)))
		return false
	})
	return c
}

func watch_IsisRouterPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter) bool) *oc.IsisRouterWatcher {
	t.Helper()
	w := &oc.IsisRouterWatcher{}
	gs := &oc.IsisRouter{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouterPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter) bool) *oc.IsisRouterWatcher {
	t.Helper()
	return watch_IsisRouterPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouterPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter) *oc.QualifiedIsisRouter {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router to the batch object.
func (n *IsisRouterPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouterPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter {
	t.Helper()
	c := &oc.CollectionIsisRouter{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouterPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter) bool) *oc.IsisRouterWatcher {
	t.Helper()
	w := &oc.IsisRouterWatcher{}
	structs := map[string]*oc.IsisRouter{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouterPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter) bool) *oc.IsisRouterWatcher {
	t.Helper()
	return watch_IsisRouterPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router to the batch object.
func (n *IsisRouterPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_CountersPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_Counters {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_CountersPath) Get(t testing.TB) *oc.IsisRouter_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters with a ONCE subscription.
func (n *IsisRouter_CountersPathAny) Get(t testing.TB) []*oc.IsisRouter_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_Counters {
	t.Helper()
	c := &oc.CollectionIsisRouter_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_Counters)))
		return false
	})
	return c
}

func watch_IsisRouter_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters) bool) *oc.IsisRouter_CountersWatcher {
	t.Helper()
	w := &oc.IsisRouter_CountersWatcher{}
	gs := &oc.IsisRouter_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters) bool) *oc.IsisRouter_CountersWatcher {
	t.Helper()
	return watch_IsisRouter_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_Counters) *oc.QualifiedIsisRouter_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters to the batch object.
func (n *IsisRouter_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_Counters {
	t.Helper()
	c := &oc.CollectionIsisRouter_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters) bool) *oc.IsisRouter_CountersWatcher {
	t.Helper()
	w := &oc.IsisRouter_CountersWatcher{}
	structs := map[string]*oc.IsisRouter_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters) bool) *oc.IsisRouter_CountersWatcher {
	t.Helper()
	return watch_IsisRouter_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters to the batch object.
func (n *IsisRouter_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1Path) Lookup(t testing.TB) *oc.QualifiedIsisRouter_Counters_Level1 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_Counters_Level1{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1Path) Get(t testing.TB) *oc.IsisRouter_Counters_Level1 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1PathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_Counters_Level1 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_Counters_Level1
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_Counters_Level1{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 with a ONCE subscription.
func (n *IsisRouter_Counters_Level1PathAny) Get(t testing.TB) []*oc.IsisRouter_Counters_Level1 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_Counters_Level1
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_Counters_Level1 {
	t.Helper()
	c := &oc.CollectionIsisRouter_Counters_Level1{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_Counters_Level1) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_Counters_Level1{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_Counters_Level1)))
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters_Level1) bool) *oc.IsisRouter_Counters_Level1Watcher {
	t.Helper()
	w := &oc.IsisRouter_Counters_Level1Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_Counters_Level1{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_Counters_Level1)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters_Level1) bool) *oc.IsisRouter_Counters_Level1Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1Path(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1Path) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_Counters_Level1) *oc.QualifiedIsisRouter_Counters_Level1 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_Counters_Level1) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 to the batch object.
func (n *IsisRouter_Counters_Level1Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_Counters_Level1 {
	t.Helper()
	c := &oc.CollectionIsisRouter_Counters_Level1{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_Counters_Level1) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1PathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters_Level1) bool) *oc.IsisRouter_Counters_Level1Watcher {
	t.Helper()
	w := &oc.IsisRouter_Counters_Level1Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_Counters_Level1{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_Counters_Level1)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_Counters_Level1) bool) *oc.IsisRouter_Counters_Level1Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1PathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1 to the batch object.
func (n *IsisRouter_Counters_Level1PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_DatabaseSizePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_DatabaseSizePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_DatabaseSizePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_DatabaseSizePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_DatabaseSizePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_DatabaseSizePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_DatabaseSizePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_DatabaseSizePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_DatabaseSizePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_DatabaseSizePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_DatabaseSizePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_DatabaseSizePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size to the batch object.
func (n *IsisRouter_Counters_Level1_DatabaseSizePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_DatabaseSizePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_DatabaseSizePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_DatabaseSizePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_DatabaseSizePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_DatabaseSizePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/database-size to the batch object.
func (n *IsisRouter_Counters_Level1_DatabaseSizePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_DatabaseSizePath extracts the value of the leaf DatabaseSize from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_DatabaseSizePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.DatabaseSize
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_InBcastHellosPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_InBcastHellosPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_InBcastHellosPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_InBcastHellosPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_InBcastHellosPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_InBcastHellosPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_InBcastHellosPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_InBcastHellosPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_InBcastHellosPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_InBcastHellosPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_InBcastHellosPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_InBcastHellosPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos to the batch object.
func (n *IsisRouter_Counters_Level1_InBcastHellosPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_InBcastHellosPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_InBcastHellosPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_InBcastHellosPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_InBcastHellosPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_InBcastHellosPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-bcast-hellos to the batch object.
func (n *IsisRouter_Counters_Level1_InBcastHellosPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_InBcastHellosPath extracts the value of the leaf InBcastHellos from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_InBcastHellosPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InBcastHellos
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_InCsnpPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_InCsnpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_InCsnpPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_InCsnpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_InCsnpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_InCsnpPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_InCsnpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_InCsnpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_InCsnpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_InCsnpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_InCsnpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_InCsnpPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp to the batch object.
func (n *IsisRouter_Counters_Level1_InCsnpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_InCsnpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_InCsnpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_InCsnpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_InCsnpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_InCsnpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-csnp to the batch object.
func (n *IsisRouter_Counters_Level1_InCsnpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_InCsnpPath extracts the value of the leaf InCsnp from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_InCsnpPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InCsnp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_InLspPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_InLspPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_InLspPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_InLspPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_InLspPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_InLspPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_InLspPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_InLspPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_InLspPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_InLspPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_InLspPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_InLspPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp to the batch object.
func (n *IsisRouter_Counters_Level1_InLspPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_InLspPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_InLspPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_InLspPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_InLspPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_InLspPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-lsp to the batch object.
func (n *IsisRouter_Counters_Level1_InLspPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_InLspPath extracts the value of the leaf InLsp from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_InLspPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InLsp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_InP2PHellosPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_InP2PHellosPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_InP2PHellosPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_InP2PHellosPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_InP2PHellosPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_InP2PHellosPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_InP2PHellosPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_InP2PHellosPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_InP2PHellosPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_InP2PHellosPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_InP2PHellosPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_InP2PHellosPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos to the batch object.
func (n *IsisRouter_Counters_Level1_InP2PHellosPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_InP2PHellosPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_InP2PHellosPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_InP2PHellosPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_InP2PHellosPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_InP2PHellosPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-p2p-hellos to the batch object.
func (n *IsisRouter_Counters_Level1_InP2PHellosPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_InP2PHellosPath extracts the value of the leaf InP2PHellos from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_InP2PHellosPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InP2PHellos
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_InPsnpPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_InPsnpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_InPsnpPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_InPsnpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_InPsnpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_InPsnpPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_InPsnpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_InPsnpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_InPsnpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_InPsnpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_InPsnpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_InPsnpPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp to the batch object.
func (n *IsisRouter_Counters_Level1_InPsnpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_InPsnpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_InPsnpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_InPsnpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_InPsnpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_InPsnpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/in-psnp to the batch object.
func (n *IsisRouter_Counters_Level1_InPsnpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_InPsnpPath extracts the value of the leaf InPsnp from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_InPsnpPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InPsnp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_OutBcastHellosPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_OutBcastHellosPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_OutBcastHellosPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_OutBcastHellosPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_OutBcastHellosPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos to the batch object.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_OutBcastHellosPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_OutBcastHellosPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_OutBcastHellosPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-bcast-hellos to the batch object.
func (n *IsisRouter_Counters_Level1_OutBcastHellosPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_OutBcastHellosPath extracts the value of the leaf OutBcastHellos from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_OutBcastHellosPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutBcastHellos
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level1_OutCsnpPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level1{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level1", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level1_OutCsnpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level1_OutCsnpPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level1_OutCsnpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level1{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level1", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level1_OutCsnpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp with a ONCE subscription.
func (n *IsisRouter_Counters_Level1_OutCsnpPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_OutCsnpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_OutCsnpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level1", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level1_OutCsnpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_OutCsnpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_OutCsnpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level1_OutCsnpPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp to the batch object.
func (n *IsisRouter_Counters_Level1_OutCsnpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level1_OutCsnpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level1_OutCsnpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level1{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level1{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level1", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level1_OutCsnpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level1_OutCsnpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level1_OutCsnpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level1/out-csnp to the batch object.
func (n *IsisRouter_Counters_Level1_OutCsnpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level1_OutCsnpPath extracts the value of the leaf OutCsnp from its parent oc.IsisRouter_Counters_Level1
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level1_OutCsnpPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level1) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutCsnp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
