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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_Counters_InPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_Counters_InPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_Counters_InPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_Counters_InPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_Counters_InPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts with a ONCE subscription.
func (n *Flow_IngressTracking_Counters_InPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_Counters_InPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_Counters_InPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_IngressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_Counters_InPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_Counters_InPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_Counters_InPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_Counters_InPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts to the batch object.
func (n *Flow_IngressTracking_Counters_InPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_Counters_InPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_Counters_InPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Flow_IngressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking_Counters", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_Counters_InPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_Counters_InPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_Counters_InPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/in-pkts to the batch object.
func (n *Flow_IngressTracking_Counters_InPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_Counters_InPktsPath extracts the value of the leaf InPkts from its parent oc.Flow_IngressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_IngressTracking_Counters_InPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_Counters_OutOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_Counters_OutOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_Counters_OutOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_Counters_OutOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_Counters_OutOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets with a ONCE subscription.
func (n *Flow_IngressTracking_Counters_OutOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_Counters_OutOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_Counters_OutOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_IngressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_Counters_OutOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_Counters_OutOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_Counters_OutOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets to the batch object.
func (n *Flow_IngressTracking_Counters_OutOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_Counters_OutOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_Counters_OutOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Flow_IngressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking_Counters", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_Counters_OutOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_Counters_OutOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_Counters_OutOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-octets to the batch object.
func (n *Flow_IngressTracking_Counters_OutOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_Counters_OutOctetsPath extracts the value of the leaf OutOctets from its parent oc.Flow_IngressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_IngressTracking_Counters_OutOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_Counters_OutPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_Counters_OutPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_Counters_OutPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_Counters_OutPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_Counters_OutPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts with a ONCE subscription.
func (n *Flow_IngressTracking_Counters_OutPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_Counters_OutPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_Counters_OutPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_IngressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_Counters_OutPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_Counters_OutPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_Counters_OutPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts to the batch object.
func (n *Flow_IngressTracking_Counters_OutPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_Counters_OutPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_Counters_OutPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Flow_IngressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking_Counters", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_Counters_OutPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_Counters_OutPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_Counters_OutPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/counters/out-pkts to the batch object.
func (n *Flow_IngressTracking_Counters_OutPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_Counters_OutPktsPath extracts the value of the leaf OutPkts from its parent oc.Flow_IngressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_IngressTracking_Counters_OutPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_DstIpv4Path) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_DstIpv4Path(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_DstIpv4Path) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_DstIpv4PathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_DstIpv4Path(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 with a ONCE subscription.
func (n *Flow_IngressTracking_DstIpv4PathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_DstIpv4Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_DstIpv4Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Flow_IngressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_DstIpv4Path(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_DstIpv4Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_DstIpv4Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_DstIpv4Path) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 to the batch object.
func (n *Flow_IngressTracking_DstIpv4Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_DstIpv4PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_DstIpv4PathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Flow_IngressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_DstIpv4Path(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_DstIpv4PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_DstIpv4PathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv4 to the batch object.
func (n *Flow_IngressTracking_DstIpv4PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_DstIpv4Path extracts the value of the leaf DstIpv4 from its parent oc.Flow_IngressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertFlow_IngressTracking_DstIpv4Path(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DstIpv4
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_DstIpv6Path) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_DstIpv6Path(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_DstIpv6Path) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_DstIpv6PathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_DstIpv6Path(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 with a ONCE subscription.
func (n *Flow_IngressTracking_DstIpv6PathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_DstIpv6Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_DstIpv6Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Flow_IngressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_DstIpv6Path(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_DstIpv6Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_DstIpv6Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_DstIpv6Path) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 to the batch object.
func (n *Flow_IngressTracking_DstIpv6Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_DstIpv6PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_DstIpv6PathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Flow_IngressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_DstIpv6Path(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_DstIpv6PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_DstIpv6PathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-ipv6 to the batch object.
func (n *Flow_IngressTracking_DstIpv6PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_DstIpv6Path extracts the value of the leaf DstIpv6 from its parent oc.Flow_IngressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertFlow_IngressTracking_DstIpv6Path(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DstIpv6
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_DstPortPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_DstPortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_DstPortPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_DstPortPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_DstPortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port with a ONCE subscription.
func (n *Flow_IngressTracking_DstPortPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_DstPortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_DstPortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Flow_IngressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_DstPortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_DstPortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_DstPortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_DstPortPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port to the batch object.
func (n *Flow_IngressTracking_DstPortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_DstPortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_DstPortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Flow_IngressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_DstPortPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_DstPortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_DstPortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/state/dst-port to the batch object.
func (n *Flow_IngressTracking_DstPortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_DstPortPath extracts the value of the leaf DstPort from its parent oc.Flow_IngressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertFlow_IngressTracking_DstPortPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DstPort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_EgressTrackingPath) Lookup(t testing.TB) *oc.QualifiedFlow_IngressTracking_EgressTracking {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_EgressTracking", goStruct, false, false)
	if ok {
		return (&oc.QualifiedFlow_IngressTracking_EgressTracking{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_EgressTrackingPath) Get(t testing.TB) *oc.Flow_IngressTracking_EgressTracking {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_EgressTrackingPathAny) Lookup(t testing.TB) []*oc.QualifiedFlow_IngressTracking_EgressTracking {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFlow_IngressTracking_EgressTracking
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_EgressTracking", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedFlow_IngressTracking_EgressTracking{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking with a ONCE subscription.
func (n *Flow_IngressTracking_EgressTrackingPathAny) Get(t testing.TB) []*oc.Flow_IngressTracking_EgressTracking {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Flow_IngressTracking_EgressTracking
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTrackingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_IngressTracking_EgressTracking {
	t.Helper()
	c := &oc.CollectionFlow_IngressTracking_EgressTracking{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_IngressTracking_EgressTracking) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedFlow_IngressTracking_EgressTracking{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Flow_IngressTracking_EgressTracking)))
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTrackingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFlow_IngressTracking_EgressTracking) bool) *oc.Flow_IngressTracking_EgressTrackingWatcher {
	t.Helper()
	w := &oc.Flow_IngressTracking_EgressTrackingWatcher{}
	gs := &oc.Flow_IngressTracking_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_EgressTracking", gs, queryPath, false, false)
		qv := (&oc.QualifiedFlow_IngressTracking_EgressTracking{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFlow_IngressTracking_EgressTracking)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTrackingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_IngressTracking_EgressTracking) bool) *oc.Flow_IngressTracking_EgressTrackingWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTrackingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_EgressTrackingPath) Await(t testing.TB, timeout time.Duration, val *oc.Flow_IngressTracking_EgressTracking) *oc.QualifiedFlow_IngressTracking_EgressTracking {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFlow_IngressTracking_EgressTracking) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking to the batch object.
func (n *Flow_IngressTracking_EgressTrackingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTrackingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_IngressTracking_EgressTracking {
	t.Helper()
	c := &oc.CollectionFlow_IngressTracking_EgressTracking{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_IngressTracking_EgressTracking) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTrackingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFlow_IngressTracking_EgressTracking) bool) *oc.Flow_IngressTracking_EgressTrackingWatcher {
	t.Helper()
	w := &oc.Flow_IngressTracking_EgressTrackingWatcher{}
	structs := map[string]*oc.Flow_IngressTracking_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking_EgressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking_EgressTracking", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedFlow_IngressTracking_EgressTracking{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFlow_IngressTracking_EgressTracking)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTrackingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_IngressTracking_EgressTracking) bool) *oc.Flow_IngressTracking_EgressTrackingWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTrackingPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking to the batch object.
func (n *Flow_IngressTracking_EgressTrackingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_EgressTracking_ConvergenceTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_EgressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_EgressTracking_ConvergenceTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time with a ONCE subscription.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_ConvergenceTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_IngressTracking_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_EgressTracking_ConvergenceTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_ConvergenceTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time to the batch object.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_ConvergenceTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Flow_IngressTracking_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking_EgressTracking{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking_EgressTracking", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_EgressTracking_ConvergenceTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_ConvergenceTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/convergence-time to the batch object.
func (n *Flow_IngressTracking_EgressTracking_ConvergenceTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_EgressTracking_ConvergenceTimePath extracts the value of the leaf ConvergenceTime from its parent oc.Flow_IngressTracking_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_IngressTracking_EgressTracking_ConvergenceTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_EgressTracking) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ConvergenceTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
