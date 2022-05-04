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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_EgressTracking_CountersPath) Lookup(t testing.TB) *oc.QualifiedFlow_IngressTracking_EgressTracking_Counters {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_EgressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_EgressTracking_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedFlow_IngressTracking_EgressTracking_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_EgressTracking_CountersPath) Get(t testing.TB) *oc.Flow_IngressTracking_EgressTracking_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_EgressTracking_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedFlow_IngressTracking_EgressTracking_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFlow_IngressTracking_EgressTracking_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_EgressTracking_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedFlow_IngressTracking_EgressTracking_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters with a ONCE subscription.
func (n *Flow_IngressTracking_EgressTracking_CountersPathAny) Get(t testing.TB) []*oc.Flow_IngressTracking_EgressTracking_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Flow_IngressTracking_EgressTracking_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_IngressTracking_EgressTracking_Counters {
	t.Helper()
	c := &oc.CollectionFlow_IngressTracking_EgressTracking_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_IngressTracking_EgressTracking_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedFlow_IngressTracking_EgressTracking_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Flow_IngressTracking_EgressTracking_Counters)))
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFlow_IngressTracking_EgressTracking_Counters) bool) *oc.Flow_IngressTracking_EgressTracking_CountersWatcher {
	t.Helper()
	w := &oc.Flow_IngressTracking_EgressTracking_CountersWatcher{}
	gs := &oc.Flow_IngressTracking_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", gs, queryPath, false, false)
		qv := (&oc.QualifiedFlow_IngressTracking_EgressTracking_Counters{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFlow_IngressTracking_EgressTracking_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_IngressTracking_EgressTracking_Counters) bool) *oc.Flow_IngressTracking_EgressTracking_CountersWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_EgressTracking_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Flow_IngressTracking_EgressTracking_Counters) *oc.QualifiedFlow_IngressTracking_EgressTracking_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFlow_IngressTracking_EgressTracking_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters to the batch object.
func (n *Flow_IngressTracking_EgressTracking_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_IngressTracking_EgressTracking_Counters {
	t.Helper()
	c := &oc.CollectionFlow_IngressTracking_EgressTracking_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_IngressTracking_EgressTracking_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_CountersPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFlow_IngressTracking_EgressTracking_Counters) bool) *oc.Flow_IngressTracking_EgressTracking_CountersWatcher {
	t.Helper()
	w := &oc.Flow_IngressTracking_EgressTracking_CountersWatcher{}
	structs := map[string]*oc.Flow_IngressTracking_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking_EgressTracking_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedFlow_IngressTracking_EgressTracking_Counters{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFlow_IngressTracking_EgressTracking_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_IngressTracking_EgressTracking_Counters) bool) *oc.Flow_IngressTracking_EgressTracking_CountersWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_CountersPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters to the batch object.
func (n *Flow_IngressTracking_EgressTracking_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_EgressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_EgressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_EgressTracking_Counters_InOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_EgressTracking_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_EgressTracking_Counters_InOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets with a ONCE subscription.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_Counters_InOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_IngressTracking_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_EgressTracking_Counters_InOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets to the batch object.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_Counters_InOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Flow_IngressTracking_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking_EgressTracking_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_EgressTracking_Counters_InOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_Counters_InOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-octets to the batch object.
func (n *Flow_IngressTracking_EgressTracking_Counters_InOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_EgressTracking_Counters_InOctetsPath extracts the value of the leaf InOctets from its parent oc.Flow_IngressTracking_EgressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_IngressTracking_EgressTracking_Counters_InOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_EgressTracking_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_EgressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_EgressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_EgressTracking_Counters_InPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_EgressTracking_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_EgressTracking_Counters_InPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts with a ONCE subscription.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_Counters_InPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_IngressTracking_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_EgressTracking_Counters_InPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_Counters_InPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts to the batch object.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_Counters_InPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Flow_IngressTracking_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking_EgressTracking_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_EgressTracking_Counters_InPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_Counters_InPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/in-pkts to the batch object.
func (n *Flow_IngressTracking_EgressTracking_Counters_InPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_EgressTracking_Counters_InPktsPath extracts the value of the leaf InPkts from its parent oc.Flow_IngressTracking_EgressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_IngressTracking_EgressTracking_Counters_InPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_EgressTracking_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_EgressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_EgressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_EgressTracking_Counters_OutOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_EgressTracking_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_EgressTracking_Counters_OutOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets with a ONCE subscription.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_Counters_OutOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_IngressTracking_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_EgressTracking_Counters_OutOctetsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets to the batch object.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_Counters_OutOctetsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Flow_IngressTracking_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking_EgressTracking_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_EgressTracking_Counters_OutOctetsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_Counters_OutOctetsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-octets to the batch object.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_EgressTracking_Counters_OutOctetsPath extracts the value of the leaf OutOctets from its parent oc.Flow_IngressTracking_EgressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_IngressTracking_EgressTracking_Counters_OutOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_EgressTracking_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_EgressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_EgressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_EgressTracking_Counters_OutPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_EgressTracking_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_EgressTracking_Counters_OutPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts with a ONCE subscription.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_Counters_OutPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_IngressTracking_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_EgressTracking_Counters_OutPktsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts to the batch object.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_Counters_OutPktsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Flow_IngressTracking_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Flow_IngressTracking_EgressTracking_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Flow_IngressTracking_EgressTracking_Counters", structs[pre], queryPath, true, false)
			qv := convertFlow_IngressTracking_EgressTracking_Counters_OutPktsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_Counters_OutPktsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/counters/out-pkts to the batch object.
func (n *Flow_IngressTracking_EgressTracking_Counters_OutPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_EgressTracking_Counters_OutPktsPath extracts the value of the leaf OutPkts from its parent oc.Flow_IngressTracking_EgressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_IngressTracking_EgressTracking_Counters_OutPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_EgressTracking_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_EgressTracking_FilterPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_EgressTracking_FilterPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_EgressTracking_FilterPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_EgressTracking_FilterPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_EgressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_EgressTracking_FilterPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter with a ONCE subscription.
func (n *Flow_IngressTracking_EgressTracking_FilterPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_FilterPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_FilterPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Flow_IngressTracking_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_EgressTracking_FilterPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_FilterPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_FilterPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_EgressTracking_FilterPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter to the batch object.
func (n *Flow_IngressTracking_EgressTracking_FilterPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_FilterPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_FilterPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertFlow_IngressTracking_EgressTracking_FilterPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_FilterPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_FilterPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/filter to the batch object.
func (n *Flow_IngressTracking_EgressTracking_FilterPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_EgressTracking_FilterPath extracts the value of the leaf Filter from its parent oc.Flow_IngressTracking_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertFlow_IngressTracking_EgressTracking_FilterPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_EgressTracking) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_EgressTracking_FirstPacketLatencyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertFlow_IngressTracking_EgressTracking_FirstPacketLatencyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency with a ONCE subscription.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_FirstPacketLatencyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_IngressTracking_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_EgressTracking_FirstPacketLatencyPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_FirstPacketLatencyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency to the batch object.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_FirstPacketLatencyPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertFlow_IngressTracking_EgressTracking_FirstPacketLatencyPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_FirstPacketLatencyPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/first-packet-latency to the batch object.
func (n *Flow_IngressTracking_EgressTracking_FirstPacketLatencyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_EgressTracking_FirstPacketLatencyPath extracts the value of the leaf FirstPacketLatency from its parent oc.Flow_IngressTracking_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_IngressTracking_EgressTracking_FirstPacketLatencyPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_EgressTracking) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePath) Lookup(t testing.TB) *oc.QualifiedFloat32 {
	t.Helper()
	goStruct := &oc.Flow_IngressTracking_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_IngressTracking_EgressTracking", goStruct, true, false)
	if ok {
		return convertFlow_IngressTracking_EgressTracking_InFrameRatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePath) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePathAny) Lookup(t testing.TB) []*oc.QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_IngressTracking_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_IngressTracking_EgressTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_IngressTracking_EgressTracking_InFrameRatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate with a ONCE subscription.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePathAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_InFrameRatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
	gs := &oc.Flow_IngressTracking_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_IngressTracking_EgressTracking", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertFlow_IngressTracking_EgressTracking_InFrameRatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_InFrameRatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePath) Await(t testing.TB, timeout time.Duration, val float32) *oc.QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate to the batch object.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat32 {
	t.Helper()
	c := &oc.CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_IngressTracking_EgressTracking_InFrameRatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	w := &oc.Float32Watcher{}
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
			qv := convertFlow_IngressTracking_EgressTracking_InFrameRatePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat32) bool) *oc.Float32Watcher {
	t.Helper()
	return watch_Flow_IngressTracking_EgressTracking_InFrameRatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/ingress-tracking/ingress-tracking/egress-tracking/egress-tracking/state/in-frame-rate to the batch object.
func (n *Flow_IngressTracking_EgressTracking_InFrameRatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_IngressTracking_EgressTracking_InFrameRatePath extracts the value of the leaf InFrameRate from its parent oc.Flow_IngressTracking_EgressTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat32.
func convertFlow_IngressTracking_EgressTracking_InFrameRatePath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_IngressTracking_EgressTracking) *oc.QualifiedFloat32 {
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
