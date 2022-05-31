package isis

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

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_InPsnpPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_InPsnpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_InPsnpPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_InPsnpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_InPsnpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_InPsnpPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_InPsnpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_InPsnpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_InPsnpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_InPsnpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_InPsnpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_InPsnpPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp to the batch object.
func (n *IsisRouter_Counters_Level2_InPsnpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_InPsnpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_InPsnpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_InPsnpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_InPsnpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_InPsnpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/in-psnp to the batch object.
func (n *IsisRouter_Counters_Level2_InPsnpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_InPsnpPath extracts the value of the leaf InPsnp from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_InPsnpPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_OutBcastHellosPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_OutBcastHellosPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_OutBcastHellosPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_OutBcastHellosPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_OutBcastHellosPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos to the batch object.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_OutBcastHellosPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_OutBcastHellosPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_OutBcastHellosPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-bcast-hellos to the batch object.
func (n *IsisRouter_Counters_Level2_OutBcastHellosPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_OutBcastHellosPath extracts the value of the leaf OutBcastHellos from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_OutBcastHellosPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_OutCsnpPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_OutCsnpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_OutCsnpPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_OutCsnpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_OutCsnpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_OutCsnpPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_OutCsnpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_OutCsnpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_OutCsnpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_OutCsnpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_OutCsnpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_OutCsnpPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp to the batch object.
func (n *IsisRouter_Counters_Level2_OutCsnpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_OutCsnpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_OutCsnpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_OutCsnpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_OutCsnpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_OutCsnpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-csnp to the batch object.
func (n *IsisRouter_Counters_Level2_OutCsnpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_OutCsnpPath extracts the value of the leaf OutCsnp from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_OutCsnpPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_OutLspPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_OutLspPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_OutLspPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_OutLspPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_OutLspPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_OutLspPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_OutLspPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_OutLspPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_OutLspPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_OutLspPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_OutLspPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_OutLspPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp to the batch object.
func (n *IsisRouter_Counters_Level2_OutLspPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_OutLspPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_OutLspPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_OutLspPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_OutLspPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_OutLspPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-lsp to the batch object.
func (n *IsisRouter_Counters_Level2_OutLspPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_OutLspPath extracts the value of the leaf OutLsp from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_OutLspPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutLsp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_OutP2PHellosPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_OutP2PHellosPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_OutP2PHellosPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_OutP2PHellosPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_OutP2PHellosPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos to the batch object.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_OutP2PHellosPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_OutP2PHellosPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_OutP2PHellosPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-p2p-hellos to the batch object.
func (n *IsisRouter_Counters_Level2_OutP2PHellosPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_OutP2PHellosPath extracts the value of the leaf OutP2PHellos from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_OutP2PHellosPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutP2PHellos
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_OutPsnpPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_OutPsnpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_OutPsnpPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_OutPsnpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_OutPsnpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_OutPsnpPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_OutPsnpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_OutPsnpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_OutPsnpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_OutPsnpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_OutPsnpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_OutPsnpPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp to the batch object.
func (n *IsisRouter_Counters_Level2_OutPsnpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_OutPsnpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_OutPsnpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_OutPsnpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_OutPsnpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_OutPsnpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/out-psnp to the batch object.
func (n *IsisRouter_Counters_Level2_OutPsnpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_OutPsnpPath extracts the value of the leaf OutPsnp from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_OutPsnpPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutPsnp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_SessionsFlapPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_SessionsFlapPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_SessionsFlapPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_SessionsFlapPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_SessionsFlapPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_SessionsFlapPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_SessionsFlapPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_SessionsFlapPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_SessionsFlapPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_SessionsFlapPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_SessionsFlapPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_SessionsFlapPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap to the batch object.
func (n *IsisRouter_Counters_Level2_SessionsFlapPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_SessionsFlapPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_SessionsFlapPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_SessionsFlapPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_SessionsFlapPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_SessionsFlapPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-flap to the batch object.
func (n *IsisRouter_Counters_Level2_SessionsFlapPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_SessionsFlapPath extracts the value of the leaf SessionsFlap from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_SessionsFlapPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.SessionsFlap
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_Counters_Level2_SessionsUpPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.IsisRouter_Counters_Level2{}
	md, ok := oc.Lookup(t, n, "IsisRouter_Counters_Level2", goStruct, true, false)
	if ok {
		return convertIsisRouter_Counters_Level2_SessionsUpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_Counters_Level2_SessionsUpPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_Counters_Level2_SessionsUpPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_Counters_Level2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_Counters_Level2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_Counters_Level2_SessionsUpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up with a ONCE subscription.
func (n *IsisRouter_Counters_Level2_SessionsUpPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_SessionsUpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_SessionsUpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_Counters_Level2", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_Counters_Level2_SessionsUpPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_SessionsUpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_SessionsUpPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_Counters_Level2_SessionsUpPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up to the batch object.
func (n *IsisRouter_Counters_Level2_SessionsUpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_Counters_Level2_SessionsUpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_Counters_Level2_SessionsUpPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.IsisRouter_Counters_Level2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_Counters_Level2{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_Counters_Level2", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_Counters_Level2_SessionsUpPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_Counters_Level2_SessionsUpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_IsisRouter_Counters_Level2_SessionsUpPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/counters/level2/sessions-up to the batch object.
func (n *IsisRouter_Counters_Level2_SessionsUpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_Counters_Level2_SessionsUpPath extracts the value of the leaf SessionsUp from its parent oc.IsisRouter_Counters_Level2
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertIsisRouter_Counters_Level2_SessionsUpPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_Counters_Level2) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.SessionsUp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.IsisRouter{}
	md, ok := oc.Lookup(t, n, "IsisRouter", goStruct, true, false)
	if ok {
		return convertIsisRouter_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a ONCE subscription.
func (n *IsisRouter_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.IsisRouter{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_NamePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/name to the batch object.
func (n *IsisRouter_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.IsisRouter{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/name to the batch object.
func (n *IsisRouter_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_NamePath extracts the value of the leaf Name from its parent oc.IsisRouter
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertIsisRouter_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter) *oc.QualifiedString {
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
