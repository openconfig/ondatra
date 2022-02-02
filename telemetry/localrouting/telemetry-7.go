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

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, true, false)
	if ok {
		return convertLocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", gs, queryPath, true, false)
		return convertLocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface to the batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/subinterface to the batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.LocalRoutes_Static_NextHop_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertLocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static_NextHop_InterfaceRef) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Subinterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_MetricPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop", goStruct, true, false)
	if ok {
		return convertLocalRoutes_Static_NextHop_MetricPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_MetricPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_MetricPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_NextHop_MetricPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_MetricPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_NextHop_MetricPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutes_Static_NextHop_MetricPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.LocalRoutes_Static_NextHop{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LocalRoutes_Static_NextHop", gs, queryPath, true, false)
		return convertLocalRoutes_Static_NextHop_MetricPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_NextHop_MetricPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LocalRoutes_Static_NextHop_MetricPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LocalRoutes_Static_NextHop_MetricPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric to the batch object.
func (n *LocalRoutes_Static_NextHop_MetricPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_NextHop_MetricPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_NextHop_MetricPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LocalRoutes_Static_NextHop_MetricPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/state/metric to the batch object.
func (n *LocalRoutes_Static_NextHop_MetricPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLocalRoutes_Static_NextHop_MetricPath extracts the value of the leaf Metric from its parent oc.LocalRoutes_Static_NextHop
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertLocalRoutes_Static_NextHop_MetricPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static_NextHop) *oc.QualifiedUint32 {
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
