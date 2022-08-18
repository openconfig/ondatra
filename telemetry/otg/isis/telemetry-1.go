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

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabasePath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabasePath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabasePathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabasePathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabasePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabasePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase) bool) *oc.IsisRouter_LinkStateDatabaseWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabaseWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabasePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase) bool) *oc.IsisRouter_LinkStateDatabaseWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabasePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabasePath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase) *oc.QualifiedIsisRouter_LinkStateDatabase {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database to the batch object.
func (n *IsisRouter_LinkStateDatabasePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabasePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabasePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase) bool) *oc.IsisRouter_LinkStateDatabaseWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabaseWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabasePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase) bool) *oc.IsisRouter_LinkStateDatabaseWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabasePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database to the batch object.
func (n *IsisRouter_LinkStateDatabasePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_LspsPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_LspsPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_LspsPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_LspsPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_LspsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_LspsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps) bool) *oc.IsisRouter_LinkStateDatabase_LspsWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_LspsWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_LspsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps) bool) *oc.IsisRouter_LinkStateDatabase_LspsWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_LspsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_LspsPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps to the batch object.
func (n *IsisRouter_LinkStateDatabase_LspsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_LspsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_LspsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps) bool) *oc.IsisRouter_LinkStateDatabase_LspsWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_LspsWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_LspsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps) bool) *oc.IsisRouter_LinkStateDatabase_LspsWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_LspsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps to the batch object.
func (n *IsisRouter_LinkStateDatabase_LspsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPath) Lookup(t testing.TB) *oc.QualifiedE_Lsps_FlagsSlice {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_FlagsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPath) Get(t testing.TB) []oc.E_Lsps_Flags {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Lsps_FlagsSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Lsps_FlagsSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_FlagsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPathAny) Get(t testing.TB) [][]oc.E_Lsps_Flags {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_Lsps_Flags
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lsps_FlagsSlice {
	t.Helper()
	c := &oc.CollectionE_Lsps_FlagsSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lsps_FlagsSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_FlagsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lsps_FlagsSlice) bool) *oc.E_Lsps_FlagsSliceWatcher {
	t.Helper()
	w := &oc.E_Lsps_FlagsSliceWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_FlagsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lsps_FlagsSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lsps_FlagsSlice) bool) *oc.E_Lsps_FlagsSliceWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_FlagsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPath) Await(t testing.TB, timeout time.Duration, val []oc.E_Lsps_Flags) *oc.QualifiedE_Lsps_FlagsSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Lsps_FlagsSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lsps_FlagsSlice {
	t.Helper()
	c := &oc.CollectionE_Lsps_FlagsSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lsps_FlagsSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_FlagsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lsps_FlagsSlice) bool) *oc.E_Lsps_FlagsSliceWatcher {
	t.Helper()
	w := &oc.E_Lsps_FlagsSliceWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_FlagsPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lsps_FlagsSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lsps_FlagsSlice) bool) *oc.E_Lsps_FlagsSliceWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_FlagsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/flags to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_FlagsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_FlagsPath extracts the value of the leaf Flags from its parent oc.IsisRouter_LinkStateDatabase_Lsps
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Lsps_FlagsSlice.
func convertIsisRouter_LinkStateDatabase_Lsps_FlagsPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps) *oc.QualifiedE_Lsps_FlagsSlice {
	t.Helper()
	qv := &oc.QualifiedE_Lsps_FlagsSlice{
		Metadata: md,
	}
	val := parent.Flags
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_IsTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_IsTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_IsTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_IsTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_IsTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_IsTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_IsTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_IsTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/is-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_IsTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_IsTypePath extracts the value of the leaf IsType from its parent oc.IsisRouter_LinkStateDatabase_Lsps
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertIsisRouter_LinkStateDatabase_Lsps_IsTypePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.IsType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_LspIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_LspIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_LspIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_LspIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_LspIdPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_LspIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_LspIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_LspIdPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/lsp-id to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_LspIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_LspIdPath extracts the value of the leaf LspId from its parent oc.IsisRouter_LinkStateDatabase_Lsps
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertIsisRouter_LinkStateDatabase_Lsps_LspIdPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LspId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_PduLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_PduLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_PduLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_PduLengthPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_PduLengthPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_PduLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_PduLengthPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_PduLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-length to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_PduLengthPath extracts the value of the leaf PduLength from its parent oc.IsisRouter_LinkStateDatabase_Lsps
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertIsisRouter_LinkStateDatabase_Lsps_PduLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.PduLength
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePath) Lookup(t testing.TB) *oc.QualifiedE_Lsps_PduType {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_PduTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePath) Get(t testing.TB) oc.E_Lsps_PduType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Lsps_PduType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Lsps_PduType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_PduTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePathAny) Get(t testing.TB) []oc.E_Lsps_PduType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Lsps_PduType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lsps_PduType {
	t.Helper()
	c := &oc.CollectionE_Lsps_PduType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lsps_PduType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_PduTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lsps_PduType) bool) *oc.E_Lsps_PduTypeWatcher {
	t.Helper()
	w := &oc.E_Lsps_PduTypeWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_PduTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lsps_PduType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lsps_PduType) bool) *oc.E_Lsps_PduTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_PduTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Lsps_PduType) *oc.QualifiedE_Lsps_PduType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Lsps_PduType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Lsps_PduType {
	t.Helper()
	c := &oc.CollectionE_Lsps_PduType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Lsps_PduType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_PduTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Lsps_PduType) bool) *oc.E_Lsps_PduTypeWatcher {
	t.Helper()
	w := &oc.E_Lsps_PduTypeWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_PduTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Lsps_PduType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Lsps_PduType) bool) *oc.E_Lsps_PduTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_PduTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/pdu-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_PduTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_PduTypePath extracts the value of the leaf PduType from its parent oc.IsisRouter_LinkStateDatabase_Lsps
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Lsps_PduType.
func convertIsisRouter_LinkStateDatabase_Lsps_PduTypePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps) *oc.QualifiedE_Lsps_PduType {
	t.Helper()
	qv := &oc.QualifiedE_Lsps_PduType{
		Metadata: md,
	}
	val := parent.PduType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/remaining-lifetime to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath extracts the value of the leaf RemainingLifetime from its parent oc.IsisRouter_LinkStateDatabase_Lsps
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertIsisRouter_LinkStateDatabase_Lsps_RemainingLifetimePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.RemainingLifetime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/state/sequence-number to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_SequenceNumberPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath extracts the value of the leaf SequenceNumber from its parent oc.IsisRouter_LinkStateDatabase_Lsps
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertIsisRouter_LinkStateDatabase_Lsps_SequenceNumberPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.SequenceNumber
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_TlvsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_TlvsWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_TlvsWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_TlvsWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_TlvsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_TlvsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_TlvsWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_TlvsWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_TlvsWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_TlvsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_TlvsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/metric to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath extracts the value of the leaf Metric from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_MetricPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Metric
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath) Lookup(t testing.TB) *oc.QualifiedE_State_FlagsSlice {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath) Get(t testing.TB) []oc.E_State_Flags {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_State_FlagsSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_State_FlagsSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPathAny) Get(t testing.TB) [][]oc.E_State_Flags {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_State_Flags
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_State_FlagsSlice {
	t.Helper()
	c := &oc.CollectionE_State_FlagsSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_State_FlagsSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_State_FlagsSlice) bool) *oc.E_State_FlagsSliceWatcher {
	t.Helper()
	w := &oc.E_State_FlagsSliceWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_State_FlagsSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_State_FlagsSlice) bool) *oc.E_State_FlagsSliceWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath) Await(t testing.TB, timeout time.Duration, val []oc.E_State_Flags) *oc.QualifiedE_State_FlagsSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_State_FlagsSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_State_FlagsSlice {
	t.Helper()
	c := &oc.CollectionE_State_FlagsSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_State_FlagsSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_State_FlagsSlice) bool) *oc.E_State_FlagsSliceWatcher {
	t.Helper()
	w := &oc.E_State_FlagsSliceWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_State_FlagsSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_State_FlagsSlice) bool) *oc.E_State_FlagsSliceWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-attributes/flags to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath extracts the value of the leaf Flags from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes
// and combines the update with an existing Metadata to return a *oc.QualifiedE_State_FlagsSlice.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes_FlagsPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) *oc.QualifiedE_State_FlagsSlice {
	t.Helper()
	qv := &oc.QualifiedE_State_FlagsSlice{
		Metadata: md,
	}
	val := parent.Flags
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix-length to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/prefix to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath extracts the value of the leaf Prefix from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Prefix
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath) Lookup(t testing.TB) *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath) Get(t testing.TB) oc.E_ExtendedIpv4Reachability_Prefix_RedistributionType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePathAny) Get(t testing.TB) []oc.E_ExtendedIpv4Reachability_Prefix_RedistributionType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_ExtendedIpv4Reachability_Prefix_RedistributionType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_ExtendedIpv4Reachability_Prefix_RedistributionType {
	t.Helper()
	c := &oc.CollectionE_ExtendedIpv4Reachability_Prefix_RedistributionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) bool) *oc.E_ExtendedIpv4Reachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	w := &oc.E_ExtendedIpv4Reachability_Prefix_RedistributionTypeWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) bool) *oc.E_ExtendedIpv4Reachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_ExtendedIpv4Reachability_Prefix_RedistributionType) *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_ExtendedIpv4Reachability_Prefix_RedistributionType {
	t.Helper()
	c := &oc.CollectionE_ExtendedIpv4Reachability_Prefix_RedistributionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) bool) *oc.E_ExtendedIpv4Reachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	w := &oc.E_ExtendedIpv4Reachability_Prefix_RedistributionTypeWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) bool) *oc.E_ExtendedIpv4Reachability_Prefix_RedistributionTypeWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-ipv4-reachability/prefixes/prefix/state/redistribution-type to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath extracts the value of the leaf RedistributionType from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_RedistributionTypePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) *oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType {
	t.Helper()
	qv := &oc.QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType{
		Metadata: md,
	}
	val := parent.RedistributionType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/extended-is-reachability/neighbors/neighbor/state/system-id to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath extracts the value of the leaf SystemId from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor_SystemIdPath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SystemId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames", goStruct, true, false)
	if ok {
		return convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames", structs[pre], queryPath, true, false)
			qv := convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/hostnames/state/hostname to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath extracts the value of the leaf Hostname from its parent oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames_HostnamePath(t testing.TB, md *genutil.Metadata, parent *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Hostname
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPath) Lookup(t testing.TB) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability {
	t.Helper()
	goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{}
	md, ok := oc.Lookup(t, n, "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability", goStruct, false, false)
	if ok {
		return (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPath) Get(t testing.TB) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPathAny) Lookup(t testing.TB) []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability with a ONCE subscription.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPathAny) Get(t testing.TB) []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability)))
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityWatcher{}
	gs := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability", gs, queryPath, false, false)
		qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPath) Await(t testing.TB, timeout time.Duration, val *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability {
	t.Helper()
	c := &oc.CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityWatcher {
	t.Helper()
	w := &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityWatcher{}
	structs := map[string]*oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) bool) *oc.IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityWatcher {
	t.Helper()
	return watch_IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-isis/isis-routers/isis-router/state/link-state-database/lsp-states/lsps/tlvs/ipv4-external-reachability to the batch object.
func (n *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
