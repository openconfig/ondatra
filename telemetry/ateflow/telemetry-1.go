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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/state/dst-ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_DstIpv6Path) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Flow{}
	md, ok := oc.Lookup(t, n, "Flow", goStruct, true, false)
	if ok {
		return convertFlow_DstIpv6Path(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/state/dst-ipv6 with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_DstIpv6Path) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/state/dst-ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_DstIpv6PathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_DstIpv6Path(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/state/dst-ipv6 with a ONCE subscription.
func (n *Flow_DstIpv6PathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/state/dst-ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_DstIpv6Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_DstIpv6Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Flow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow", gs, queryPath, true, false)
		return convertFlow_DstIpv6Path(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/state/dst-ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_DstIpv6Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_DstIpv6Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/state/dst-ipv6 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_DstIpv6Path) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/state/dst-ipv6 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/state/dst-ipv6 to the batch object.
func (n *Flow_DstIpv6Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/state/dst-ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_DstIpv6PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/state/dst-ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_DstIpv6PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_DstIpv6Path(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/state/dst-ipv6 to the batch object.
func (n *Flow_DstIpv6PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_DstIpv6Path extracts the value of the leaf DstIpv6 from its parent oc.Flow
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertFlow_DstIpv6Path(t testing.TB, md *genutil.Metadata, parent *oc.Flow) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/state/dst-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_DstPortPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Flow{}
	md, ok := oc.Lookup(t, n, "Flow", goStruct, true, false)
	if ok {
		return convertFlow_DstPortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/state/dst-port with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_DstPortPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/state/dst-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_DstPortPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertFlow_DstPortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/state/dst-port with a ONCE subscription.
func (n *Flow_DstPortPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/state/dst-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_DstPortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_DstPortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Flow{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow", gs, queryPath, true, false)
		return convertFlow_DstPortPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/state/dst-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_DstPortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_DstPortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/state/dst-port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_DstPortPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/state/dst-port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/state/dst-port to the batch object.
func (n *Flow_DstPortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/state/dst-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_DstPortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/state/dst-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_DstPortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Flow_DstPortPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/state/dst-port to the batch object.
func (n *Flow_DstPortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_DstPortPath extracts the value of the leaf DstPort from its parent oc.Flow
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertFlow_DstPortPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTrackingPath) Lookup(t testing.TB) *oc.QualifiedFlow_EgressTracking {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking", goStruct, false, false)
	if ok {
		return (&oc.QualifiedFlow_EgressTracking{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTrackingPath) Get(t testing.TB) *oc.Flow_EgressTracking {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTrackingPathAny) Lookup(t testing.TB) []*oc.QualifiedFlow_EgressTracking {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFlow_EgressTracking
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_EgressTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_EgressTracking", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedFlow_EgressTracking{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking with a ONCE subscription.
func (n *Flow_EgressTrackingPathAny) Get(t testing.TB) []*oc.Flow_EgressTracking {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Flow_EgressTracking
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTrackingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_EgressTracking {
	t.Helper()
	c := &oc.CollectionFlow_EgressTracking{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_EgressTracking) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedFlow_EgressTracking{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Flow_EgressTracking)))
		return false
	})
	return c
}

func watch_Flow_EgressTrackingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFlow_EgressTracking) bool) *oc.Flow_EgressTrackingWatcher {
	t.Helper()
	w := &oc.Flow_EgressTrackingWatcher{}
	gs := &oc.Flow_EgressTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking", gs, queryPath, false, false)
		return (&oc.QualifiedFlow_EgressTracking{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFlow_EgressTracking)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTrackingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_EgressTracking) bool) *oc.Flow_EgressTrackingWatcher {
	t.Helper()
	return watch_Flow_EgressTrackingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTrackingPath) Await(t testing.TB, timeout time.Duration, val *oc.Flow_EgressTracking) *oc.QualifiedFlow_EgressTracking {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFlow_EgressTracking) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking to the batch object.
func (n *Flow_EgressTrackingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTrackingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_EgressTracking {
	t.Helper()
	c := &oc.CollectionFlow_EgressTracking{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_EgressTracking) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTrackingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_EgressTracking) bool) *oc.Flow_EgressTrackingWatcher {
	t.Helper()
	return watch_Flow_EgressTrackingPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking to the batch object.
func (n *Flow_EgressTrackingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_CountersPath) Lookup(t testing.TB) *oc.QualifiedFlow_EgressTracking_Counters {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking_Counters", goStruct, false, false)
	if ok {
		return (&oc.QualifiedFlow_EgressTracking_Counters{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_CountersPath) Get(t testing.TB) *oc.Flow_EgressTracking_Counters {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_CountersPathAny) Lookup(t testing.TB) []*oc.QualifiedFlow_EgressTracking_Counters {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFlow_EgressTracking_Counters
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Flow_EgressTracking_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Flow_EgressTracking_Counters", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedFlow_EgressTracking_Counters{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters with a ONCE subscription.
func (n *Flow_EgressTracking_CountersPathAny) Get(t testing.TB) []*oc.Flow_EgressTracking_Counters {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Flow_EgressTracking_Counters
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_CountersPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_EgressTracking_Counters {
	t.Helper()
	c := &oc.CollectionFlow_EgressTracking_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_EgressTracking_Counters) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedFlow_EgressTracking_Counters{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Flow_EgressTracking_Counters)))
		return false
	})
	return c
}

func watch_Flow_EgressTracking_CountersPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFlow_EgressTracking_Counters) bool) *oc.Flow_EgressTracking_CountersWatcher {
	t.Helper()
	w := &oc.Flow_EgressTracking_CountersWatcher{}
	gs := &oc.Flow_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking_Counters", gs, queryPath, false, false)
		return (&oc.QualifiedFlow_EgressTracking_Counters{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFlow_EgressTracking_Counters)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_CountersPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_EgressTracking_Counters) bool) *oc.Flow_EgressTracking_CountersWatcher {
	t.Helper()
	return watch_Flow_EgressTracking_CountersPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_CountersPath) Await(t testing.TB, timeout time.Duration, val *oc.Flow_EgressTracking_Counters) *oc.QualifiedFlow_EgressTracking_Counters {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFlow_EgressTracking_Counters) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters to the batch object.
func (n *Flow_EgressTracking_CountersPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_CountersPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFlow_EgressTracking_Counters {
	t.Helper()
	c := &oc.CollectionFlow_EgressTracking_Counters{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFlow_EgressTracking_Counters) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_CountersPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFlow_EgressTracking_Counters) bool) *oc.Flow_EgressTracking_CountersWatcher {
	t.Helper()
	return watch_Flow_EgressTracking_CountersPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters to the batch object.
func (n *Flow_EgressTracking_CountersPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_Counters_InOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_Counters_InOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_Counters_InOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_Counters_InOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertFlow_EgressTracking_Counters_InOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets with a ONCE subscription.
func (n *Flow_EgressTracking_Counters_InOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_Counters_InOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_Counters_InOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking_Counters", gs, queryPath, true, false)
		return convertFlow_EgressTracking_Counters_InOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_Counters_InOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_Counters_InOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets to the batch object.
func (n *Flow_EgressTracking_Counters_InOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_Counters_InOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_Counters_InOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-octets to the batch object.
func (n *Flow_EgressTracking_Counters_InOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_Counters_InOctetsPath extracts the value of the leaf InOctets from its parent oc.Flow_EgressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_EgressTracking_Counters_InOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_Counters_InPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_Counters_InPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_Counters_InPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_Counters_InPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertFlow_EgressTracking_Counters_InPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts with a ONCE subscription.
func (n *Flow_EgressTracking_Counters_InPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_Counters_InPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_Counters_InPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking_Counters", gs, queryPath, true, false)
		return convertFlow_EgressTracking_Counters_InPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_Counters_InPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_Counters_InPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_Counters_InPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts to the batch object.
func (n *Flow_EgressTracking_Counters_InPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_Counters_InPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_Counters_InPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_Counters_InPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/in-pkts to the batch object.
func (n *Flow_EgressTracking_Counters_InPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_Counters_InPktsPath extracts the value of the leaf InPkts from its parent oc.Flow_EgressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_EgressTracking_Counters_InPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking_Counters) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Flow_EgressTracking_Counters_OutOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Flow_EgressTracking_Counters{}
	md, ok := oc.Lookup(t, n, "Flow_EgressTracking_Counters", goStruct, true, false)
	if ok {
		return convertFlow_EgressTracking_Counters_OutOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Flow_EgressTracking_Counters_OutOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Flow_EgressTracking_Counters_OutOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertFlow_EgressTracking_Counters_OutOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets with a ONCE subscription.
func (n *Flow_EgressTracking_Counters_OutOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_Counters_OutOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Flow_EgressTracking_Counters_OutOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Flow_EgressTracking_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Flow_EgressTracking_Counters", gs, queryPath, true, false)
		return convertFlow_EgressTracking_Counters_OutOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_Counters_OutOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Flow_EgressTracking_Counters_OutOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets to the batch object.
func (n *Flow_EgressTracking_Counters_OutOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Flow_EgressTracking_Counters_OutOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Flow_EgressTracking_Counters_OutOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Flow_EgressTracking_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-ate-flow/flows/flow/egress-tracking/egress-tracking/state/counters/out-octets to the batch object.
func (n *Flow_EgressTracking_Counters_OutOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertFlow_EgressTracking_Counters_OutOctetsPath extracts the value of the leaf OutOctets from its parent oc.Flow_EgressTracking_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertFlow_EgressTracking_Counters_OutOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Flow_EgressTracking_Counters) *oc.QualifiedUint64 {
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
