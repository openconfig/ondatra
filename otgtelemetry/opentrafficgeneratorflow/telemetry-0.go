package opentrafficgeneratorflow

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

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *FlowPath) Lookup(t testing.TB) *oc.QualifiedFlow {
	t.Helper()
	goStruct := &oc.Flow{}
	md, ok := oc.Lookup(t, n, "Flow", goStruct, false, false)
	if ok {
		return (&oc.QualifiedFlow{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *FlowPath) Get(t testing.TB) *oc.Flow {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *FlowPathAny) Lookup(t testing.TB) []*oc.QualifiedFlow {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFlow
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedFlow{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow with a ONCE subscription.
func (n *FlowPathAny) Get(t testing.TB) []*oc.Flow {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Flow
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *FlowPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow {
	t.Helper()
	c := &oc.CollectionFlow{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedFlow{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Flow)))
		return false
	})
	return c
}

func watch_FlowPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFlow) bool) *oc.FlowWatcher {
	t.Helper()
	w := &oc.FlowWatcher{}
	gs := &oc.Flow{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow", gs, queryPath, false, false)
		return (&oc.QualifiedFlow{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFlow)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *FlowPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow) bool) *oc.FlowWatcher {
	t.Helper()
	return watch_FlowPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *FlowPath) Await(t testing.TB, timeout time.Duration, val *oc.Flow) *oc.QualifiedFlow {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFlow) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow to the batch object.
func (n *FlowPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *FlowPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow {
	t.Helper()
	c := &oc.CollectionFlow{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *FlowPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow) bool) *oc.FlowWatcher {
	t.Helper()
	return watch_FlowPath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow to the batch object.
func (n *FlowPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_CountersPath) Lookup(t testing.TB) *oc.QualifiedFlow_Counters {
	t.Helper()
	goStruct := &oc.Flow_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedFlow_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/state/counters with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_CountersPath) Get(t testing.TB) *oc.Flow_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedFlow_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFlow_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedFlow_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/state/counters with a ONCE subscription.
func (n *Flow_CountersPathAny) Get(t testing.TB) []*oc.Flow_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Flow_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_Counters {
	t.Helper()
	c := &oc.CollectionFlow_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedFlow_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Flow_Counters)))
		return false
	})
	return c
}

func watch_Flow_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFlow_Counters) bool) *oc.Flow_CountersWatcher {
	t.Helper()
	w := &oc.Flow_CountersWatcher{}
	gs := &oc.Flow_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_Counters", gs, queryPath, false, false)
		return (&oc.QualifiedFlow_Counters{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFlow_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_Counters) bool) *oc.Flow_CountersWatcher {
	t.Helper()
	return watch_Flow_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Flow_Counters) *oc.QualifiedFlow_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFlow_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/counters to the batch object.
func (n *Flow_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_Counters {
	t.Helper()
	c := &oc.CollectionFlow_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_Counters) bool) *oc.Flow_CountersWatcher {
	t.Helper()
	return watch_Flow_CountersPath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/counters to the batch object.
func (n *Flow_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/state/counters/in-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_Counters_InOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_Counters", goStruct, true, false)
	if ok {
		return convertFlow_Counters_InOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/state/counters/in-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_Counters_InOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/state/counters/in-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_Counters_InOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_Counters_InOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/state/counters/in-octets with a ONCE subscription.
func (n *Flow_Counters_InOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Counters_InOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_Counters_InOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_Counters", gs, queryPath, true, false)
		return convertFlow_Counters_InOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Counters_InOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/state/counters/in-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_Counters_InOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/state/counters/in-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/counters/in-octets to the batch object.
func (n *Flow_Counters_InOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Counters_InOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Counters_InOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/counters/in-octets to the batch object.
func (n *Flow_Counters_InOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_Counters_InOctetsPath extracts the value of the leaf InOctets from its parent oc.Flow_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_Counters_InOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/state/counters/in-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_Counters_InPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_Counters", goStruct, true, false)
	if ok {
		return convertFlow_Counters_InPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/state/counters/in-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_Counters_InPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/state/counters/in-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_Counters_InPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_Counters_InPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/state/counters/in-pkts with a ONCE subscription.
func (n *Flow_Counters_InPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Counters_InPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_Counters_InPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_Counters", gs, queryPath, true, false)
		return convertFlow_Counters_InPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Counters_InPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_Counters_InPktsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/state/counters/in-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_Counters_InPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/state/counters/in-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/counters/in-pkts to the batch object.
func (n *Flow_Counters_InPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Counters_InPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Counters_InPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_Counters_InPktsPath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/counters/in-pkts to the batch object.
func (n *Flow_Counters_InPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_Counters_InPktsPath extracts the value of the leaf InPkts from its parent oc.Flow_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_Counters_InPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/state/counters/out-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_Counters_OutOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_Counters", goStruct, true, false)
	if ok {
		return convertFlow_Counters_OutOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/state/counters/out-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_Counters_OutOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/state/counters/out-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_Counters_OutOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_Counters_OutOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/state/counters/out-octets with a ONCE subscription.
func (n *Flow_Counters_OutOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Counters_OutOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_Counters_OutOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_Counters", gs, queryPath, true, false)
		return convertFlow_Counters_OutOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Counters_OutOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/state/counters/out-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_Counters_OutOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/state/counters/out-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/counters/out-octets to the batch object.
func (n *Flow_Counters_OutOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Counters_OutOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Counters_OutOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/counters/out-octets to the batch object.
func (n *Flow_Counters_OutOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_Counters_OutOctetsPath extracts the value of the leaf OutOctets from its parent oc.Flow_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_Counters_OutOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/state/counters/out-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_Counters_OutPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_Counters", goStruct, true, false)
	if ok {
		return convertFlow_Counters_OutPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/state/counters/out-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_Counters_OutPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/state/counters/out-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_Counters_OutPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_Counters_OutPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/state/counters/out-pkts with a ONCE subscription.
func (n *Flow_Counters_OutPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Counters_OutPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_Counters_OutPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_Counters", gs, queryPath, true, false)
		return convertFlow_Counters_OutPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Counters_OutPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/state/counters/out-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_Counters_OutPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/state/counters/out-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/counters/out-pkts to the batch object.
func (n *Flow_Counters_OutPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Counters_OutPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Counters_OutPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/state/counters/out-pkts to the batch object.
func (n *Flow_Counters_OutPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_Counters_OutPktsPath extracts the value of the leaf OutPkts from its parent oc.Flow_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_Counters_OutPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/enumerations/enumeration with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EnumerationPath) Lookup(t testing.TB) *oc.QualifiedFlow_Enumeration {
	t.Helper()
	goStruct := &oc.Flow_Enumeration{}
	md, ok := oc.Lookup(t, n, "Flow_Enumeration", goStruct, false, false)
	if ok {
		return (&oc.QualifiedFlow_Enumeration{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/enumerations/enumeration with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EnumerationPath) Get(t testing.TB) *oc.Flow_Enumeration {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EnumerationPathAny) Lookup(t testing.TB) []*oc.QualifiedFlow_Enumeration {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFlow_Enumeration
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_Enumeration{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_Enumeration", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedFlow_Enumeration{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration with a ONCE subscription.
func (n *Flow_EnumerationPathAny) Get(t testing.TB) []*oc.Flow_Enumeration {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Flow_Enumeration
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EnumerationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_Enumeration {
	t.Helper()
	c := &oc.CollectionFlow_Enumeration{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_Enumeration) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedFlow_Enumeration{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Flow_Enumeration)))
		return false
	})
	return c
}

func watch_Flow_EnumerationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFlow_Enumeration) bool) *oc.Flow_EnumerationWatcher {
	t.Helper()
	w := &oc.Flow_EnumerationWatcher{}
	gs := &oc.Flow_Enumeration{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_Enumeration", gs, queryPath, false, false)
		return (&oc.QualifiedFlow_Enumeration{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFlow_Enumeration)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EnumerationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_Enumeration) bool) *oc.Flow_EnumerationWatcher {
	t.Helper()
	return watch_Flow_EnumerationPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EnumerationPath) Await(t testing.TB, timeout time.Duration, val *oc.Flow_Enumeration) *oc.QualifiedFlow_Enumeration {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFlow_Enumeration) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/enumerations/enumeration failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/enumerations/enumeration to the batch object.
func (n *Flow_EnumerationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EnumerationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_Enumeration {
	t.Helper()
	c := &oc.CollectionFlow_Enumeration{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_Enumeration) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EnumerationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_Enumeration) bool) *oc.Flow_EnumerationWatcher {
	t.Helper()
	return watch_Flow_EnumerationPath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/enumerations/enumeration to the batch object.
func (n *Flow_EnumerationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_Enumeration_CountersPath) Lookup(t testing.TB) *oc.QualifiedFlow_Enumeration_Counters {
	t.Helper()
	goStruct := &oc.Flow_Enumeration_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_Enumeration_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedFlow_Enumeration_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_Enumeration_CountersPath) Get(t testing.TB) *oc.Flow_Enumeration_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_Enumeration_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedFlow_Enumeration_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFlow_Enumeration_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_Enumeration_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_Enumeration_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedFlow_Enumeration_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters with a ONCE subscription.
func (n *Flow_Enumeration_CountersPathAny) Get(t testing.TB) []*oc.Flow_Enumeration_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Flow_Enumeration_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Enumeration_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_Enumeration_Counters {
	t.Helper()
	c := &oc.CollectionFlow_Enumeration_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_Enumeration_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedFlow_Enumeration_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Flow_Enumeration_Counters)))
		return false
	})
	return c
}

func watch_Flow_Enumeration_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFlow_Enumeration_Counters) bool) *oc.Flow_Enumeration_CountersWatcher {
	t.Helper()
	w := &oc.Flow_Enumeration_CountersWatcher{}
	gs := &oc.Flow_Enumeration_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_Enumeration_Counters", gs, queryPath, false, false)
		return (&oc.QualifiedFlow_Enumeration_Counters{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFlow_Enumeration_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Enumeration_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_Enumeration_Counters) bool) *oc.Flow_Enumeration_CountersWatcher {
	t.Helper()
	return watch_Flow_Enumeration_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_Enumeration_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Flow_Enumeration_Counters) *oc.QualifiedFlow_Enumeration_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFlow_Enumeration_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters to the batch object.
func (n *Flow_Enumeration_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Enumeration_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_Enumeration_Counters {
	t.Helper()
	c := &oc.CollectionFlow_Enumeration_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_Enumeration_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Enumeration_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_Enumeration_Counters) bool) *oc.Flow_Enumeration_CountersWatcher {
	t.Helper()
	return watch_Flow_Enumeration_CountersPath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters to the batch object.
func (n *Flow_Enumeration_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_Enumeration_Counters_InOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_Enumeration_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_Enumeration_Counters", goStruct, true, false)
	if ok {
		return convertFlow_Enumeration_Counters_InOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_Enumeration_Counters_InOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_Enumeration_Counters_InOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_Enumeration_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_Enumeration_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_Enumeration_Counters_InOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets with a ONCE subscription.
func (n *Flow_Enumeration_Counters_InOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Enumeration_Counters_InOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_Enumeration_Counters_InOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_Enumeration_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_Enumeration_Counters", gs, queryPath, true, false)
		return convertFlow_Enumeration_Counters_InOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Enumeration_Counters_InOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_Enumeration_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_Enumeration_Counters_InOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets to the batch object.
func (n *Flow_Enumeration_Counters_InOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_Enumeration_Counters_InOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_Enumeration_Counters_InOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_Enumeration_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-flow/flows/flow/enumerations/enumeration/state/counters/in-octets to the batch object.
func (n *Flow_Enumeration_Counters_InOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_Enumeration_Counters_InOctetsPath extracts the value of the leaf InOctets from its parent oc.Flow_Enumeration_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_Enumeration_Counters_InOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_Enumeration_Counters) *oc.QualifiedUint64 {
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
