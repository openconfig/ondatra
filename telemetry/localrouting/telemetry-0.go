package localrouting

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

// Lookup fetches the value at /openconfig-local-routing/local-routes with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutesPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes {
	t.Helper()
	goStruct := &oc.LocalRoutes{}
	md, ok := oc.Lookup(t, n, "LocalRoutes", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLocalRoutes{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutesPath) Get(t testing.TB) *oc.LocalRoutes {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutesPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLocalRoutes{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes with a ONCE subscription.
func (n *LocalRoutesPathAny) Get(t testing.TB) []*oc.LocalRoutes {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LocalRoutes
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLocalRoutes {
	t.Helper()
	c := &oc.CollectionLocalRoutes{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLocalRoutes) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLocalRoutes{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.LocalRoutes)))
		return false
	})
	return c
}

func watch_LocalRoutesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLocalRoutes) bool) *oc.LocalRoutesWatcher {
	t.Helper()
	w := &oc.LocalRoutesWatcher{}
	gs := &oc.LocalRoutes{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LocalRoutes", gs, queryPath, false, false)
		qv := (&oc.QualifiedLocalRoutes{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLocalRoutes)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLocalRoutes) bool) *oc.LocalRoutesWatcher {
	t.Helper()
	return watch_LocalRoutesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-local-routing/local-routes with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LocalRoutesPath) Await(t testing.TB, timeout time.Duration, val *oc.LocalRoutes) *oc.QualifiedLocalRoutes {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLocalRoutes) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-local-routing/local-routes failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-local-routing/local-routes to the batch object.
func (n *LocalRoutesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLocalRoutes {
	t.Helper()
	c := &oc.CollectionLocalRoutes{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLocalRoutes) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLocalRoutes) bool) *oc.LocalRoutesWatcher {
	t.Helper()
	w := &oc.LocalRoutesWatcher{}
	structs := map[string]*oc.LocalRoutes{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LocalRoutes{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LocalRoutes", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLocalRoutes{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLocalRoutes)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLocalRoutes) bool) *oc.LocalRoutesWatcher {
	t.Helper()
	return watch_LocalRoutesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-local-routing/local-routes to the batch object.
func (n *LocalRoutesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_AggregatePath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Aggregate {
	t.Helper()
	goStruct := &oc.LocalRoutes_Aggregate{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Aggregate", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLocalRoutes_Aggregate{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_AggregatePath) Get(t testing.TB) *oc.LocalRoutes_Aggregate {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_AggregatePathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Aggregate {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Aggregate
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Aggregate{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Aggregate", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLocalRoutes_Aggregate{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a ONCE subscription.
func (n *LocalRoutes_AggregatePathAny) Get(t testing.TB) []*oc.LocalRoutes_Aggregate {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LocalRoutes_Aggregate
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_AggregatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLocalRoutes_Aggregate {
	t.Helper()
	c := &oc.CollectionLocalRoutes_Aggregate{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLocalRoutes_Aggregate) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLocalRoutes_Aggregate{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.LocalRoutes_Aggregate)))
		return false
	})
	return c
}

func watch_LocalRoutes_AggregatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Aggregate) bool) *oc.LocalRoutes_AggregateWatcher {
	t.Helper()
	w := &oc.LocalRoutes_AggregateWatcher{}
	gs := &oc.LocalRoutes_Aggregate{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LocalRoutes_Aggregate", gs, queryPath, false, false)
		qv := (&oc.QualifiedLocalRoutes_Aggregate{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLocalRoutes_Aggregate)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_AggregatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Aggregate) bool) *oc.LocalRoutes_AggregateWatcher {
	t.Helper()
	return watch_LocalRoutes_AggregatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LocalRoutes_AggregatePath) Await(t testing.TB, timeout time.Duration, val *oc.LocalRoutes_Aggregate) *oc.QualifiedLocalRoutes_Aggregate {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLocalRoutes_Aggregate) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-local-routing/local-routes/local-aggregates/aggregate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-local-routing/local-routes/local-aggregates/aggregate to the batch object.
func (n *LocalRoutes_AggregatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_AggregatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLocalRoutes_Aggregate {
	t.Helper()
	c := &oc.CollectionLocalRoutes_Aggregate{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLocalRoutes_Aggregate) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutes_AggregatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Aggregate) bool) *oc.LocalRoutes_AggregateWatcher {
	t.Helper()
	w := &oc.LocalRoutes_AggregateWatcher{}
	structs := map[string]*oc.LocalRoutes_Aggregate{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LocalRoutes_Aggregate{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LocalRoutes_Aggregate", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLocalRoutes_Aggregate{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLocalRoutes_Aggregate)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_AggregatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Aggregate) bool) *oc.LocalRoutes_AggregateWatcher {
	t.Helper()
	return watch_LocalRoutes_AggregatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-local-routing/local-routes/local-aggregates/aggregate to the batch object.
func (n *LocalRoutes_AggregatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
