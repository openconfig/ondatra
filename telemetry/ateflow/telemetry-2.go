package ateflow

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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_Counters_OutPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_Counters_OutPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_Counters_OutPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_Counters_OutPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_EgressTracking_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_EgressTracking_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_EgressTracking_Counters_OutPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts with a ONCE subscription.
func (n *Flow_EgressTracking_Counters_OutPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_Counters_OutPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_Counters_OutPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_EgressTracking_Counters_OutPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_Counters_OutPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_Counters_OutPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts to the batch object.
func (n *Flow_EgressTracking_Counters_OutPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_Counters_OutPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_Counters_OutPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Flow_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_EgressTracking_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_EgressTracking_Counters", structs[pre], queryPath, true, false)
			qv := convertFlow_EgressTracking_Counters_OutPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_Counters_OutPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_Counters_OutPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-pkts to the batch object.
func (n *Flow_EgressTracking_Counters_OutPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_Counters_OutPktsPath extracts the value of the leaf OutPkts from its parent oc.Flow_EgressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_EgressTracking_Counters_OutPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_FilterPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_FilterPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_FilterPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_FilterPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_EgressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_EgressTracking_FilterPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter with a ONCE subscription.
func (n *Flow_EgressTracking_FilterPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_FilterPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_FilterPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_EgressTracking_FilterPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_FilterPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_EgressTracking_FilterPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_FilterPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter to the batch object.
func (n *Flow_EgressTracking_FilterPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_FilterPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_FilterPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_EgressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_EgressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_EgressTracking_FilterPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_FilterPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_EgressTracking_FilterPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/filter to the batch object.
func (n *Flow_EgressTracking_FilterPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_FilterPath extracts the value of the leaf Filter from its parent oc.Flow_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertFlow_EgressTracking_FilterPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Filter
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_FirstPacketLatencyPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_FirstPacketLatencyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_FirstPacketLatencyPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_FirstPacketLatencyPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_EgressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_EgressTracking_FirstPacketLatencyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency with a ONCE subscription.
func (n *Flow_EgressTracking_FirstPacketLatencyPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_FirstPacketLatencyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_FirstPacketLatencyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_EgressTracking_FirstPacketLatencyPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_FirstPacketLatencyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_FirstPacketLatencyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_FirstPacketLatencyPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency to the batch object.
func (n *Flow_EgressTracking_FirstPacketLatencyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_FirstPacketLatencyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_FirstPacketLatencyPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_EgressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_EgressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_EgressTracking_FirstPacketLatencyPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_FirstPacketLatencyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_FirstPacketLatencyPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/first-packet-latency to the batch object.
func (n *Flow_EgressTracking_FirstPacketLatencyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_FirstPacketLatencyPath extracts the value of the leaf FirstPacketLatency from its parent oc.Flow_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_EgressTracking_FirstPacketLatencyPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FirstPacketLatency
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_InFrameRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_InFrameRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_InFrameRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_InFrameRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_EgressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_EgressTracking_InFrameRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate with a ONCE subscription.
func (n *Flow_EgressTracking_InFrameRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_InFrameRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_InFrameRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_EgressTracking_InFrameRatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_InFrameRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_InFrameRatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_InFrameRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate to the batch object.
func (n *Flow_EgressTracking_InFrameRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_InFrameRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_InFrameRatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	structs := map[string]*oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_EgressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_EgressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_EgressTracking_InFrameRatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_InFrameRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_InFrameRatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-frame-rate to the batch object.
func (n *Flow_EgressTracking_InFrameRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_InFrameRatePath extracts the value of the leaf InFrameRate from its parent oc.Flow_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertFlow_EgressTracking_InFrameRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking) *oc.QualifiedFloat32 {
	t.Helper()
	qv := &oc.QualifiedFloat32{
		Metadata: md,
	}
	val := parent.InFrameRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_InRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_InRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_InRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_InRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_EgressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_EgressTracking_InRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate with a ONCE subscription.
func (n *Flow_EgressTracking_InRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_InRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_InRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_EgressTracking_InRatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_InRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_InRatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_InRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate to the batch object.
func (n *Flow_EgressTracking_InRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_InRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_InRatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	structs := map[string]*oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_EgressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_EgressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_EgressTracking_InRatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_InRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_InRatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/in-rate to the batch object.
func (n *Flow_EgressTracking_InRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_InRatePath extracts the value of the leaf InRate from its parent oc.Flow_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertFlow_EgressTracking_InRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking) *oc.QualifiedFloat32 {
	t.Helper()
	qv := &oc.QualifiedFloat32{
		Metadata: md,
	}
	val := parent.InRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_LossPctPath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_LossPctPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_LossPctPath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_LossPctPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_EgressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_EgressTracking_LossPctPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct with a ONCE subscription.
func (n *Flow_EgressTracking_LossPctPathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_LossPctPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_LossPctPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_EgressTracking_LossPctPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_LossPctPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_LossPctPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_LossPctPath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct to the batch object.
func (n *Flow_EgressTracking_LossPctPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_LossPctPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_LossPctPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	structs := map[string]*oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_EgressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_EgressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_EgressTracking_LossPctPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_LossPctPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_LossPctPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/loss-pct to the batch object.
func (n *Flow_EgressTracking_LossPctPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_LossPctPath extracts the value of the leaf LossPct from its parent oc.Flow_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertFlow_EgressTracking_LossPctPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking) *oc.QualifiedFloat32 {
	t.Helper()
	qv := &oc.QualifiedFloat32{
		Metadata: md,
	}
	val := parent.LossPct
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_OutFrameRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_OutFrameRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_OutFrameRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_OutFrameRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_EgressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_EgressTracking_OutFrameRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate with a ONCE subscription.
func (n *Flow_EgressTracking_OutFrameRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_OutFrameRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_OutFrameRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_EgressTracking_OutFrameRatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_OutFrameRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_OutFrameRatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_OutFrameRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate to the batch object.
func (n *Flow_EgressTracking_OutFrameRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_OutFrameRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_OutFrameRatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	structs := map[string]*oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_EgressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_EgressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_EgressTracking_OutFrameRatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_OutFrameRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_OutFrameRatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-frame-rate to the batch object.
func (n *Flow_EgressTracking_OutFrameRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_OutFrameRatePath extracts the value of the leaf OutFrameRate from its parent oc.Flow_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertFlow_EgressTracking_OutFrameRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking) *oc.QualifiedFloat32 {
	t.Helper()
	qv := &oc.QualifiedFloat32{
		Metadata: md,
	}
	val := parent.OutFrameRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_OutRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_OutRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_OutRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_OutRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_EgressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_EgressTracking_OutRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate with a ONCE subscription.
func (n *Flow_EgressTracking_OutRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_OutRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_OutRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_EgressTracking_OutRatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_OutRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_OutRatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_OutRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate to the batch object.
func (n *Flow_EgressTracking_OutRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_OutRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_OutRatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	structs := map[string]*oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_EgressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_EgressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_EgressTracking_OutRatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_OutRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_OutRatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/out-rate to the batch object.
func (n *Flow_EgressTracking_OutRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_OutRatePath extracts the value of the leaf OutRate from its parent oc.Flow_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertFlow_EgressTracking_OutRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking) *oc.QualifiedFloat32 {
	t.Helper()
	qv := &oc.QualifiedFloat32{
		Metadata: md,
	}
	val := parent.OutRate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
	}
	return qv
}
