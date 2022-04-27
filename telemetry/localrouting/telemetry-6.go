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

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Get(t testing.TB) *oc.LocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPathAny) Get(t testing.TB) []*oc.LocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LocalRoutes_Static_NextHop_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	c := &oc.CollectionLocalRoutes_Static_NextHop_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.LocalRoutes_Static_NextHop_InterfaceRef)))
		return false
	})
	return c
}

func watch_LocalRoutes_Static_NextHop_InterfaceRefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef) bool) *oc.LocalRoutes_Static_NextHop_InterfaceRefWatcher {
	t.Helper()
	w := &oc.LocalRoutes_Static_NextHop_InterfaceRefWatcher{}
	gs := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", gs, queryPath, false, false)
		qv := (&oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef) bool) *oc.LocalRoutes_Static_NextHop_InterfaceRefWatcher {
	t.Helper()
	return watch_LocalRoutes_Static_NextHop_InterfaceRefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Await(t testing.TB, timeout time.Duration, val *oc.LocalRoutes_Static_NextHop_InterfaceRef) *oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref to the batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	c := &oc.CollectionLocalRoutes_Static_NextHop_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutes_Static_NextHop_InterfaceRefPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef) bool) *oc.LocalRoutes_Static_NextHop_InterfaceRefWatcher {
	t.Helper()
	w := &oc.LocalRoutes_Static_NextHop_InterfaceRefWatcher{}
	structs := map[string]*oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef) bool) *oc.LocalRoutes_Static_NextHop_InterfaceRefWatcher {
	t.Helper()
	return watch_LocalRoutes_Static_NextHop_InterfaceRefPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref to the batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, true, false)
	if ok {
		return convertLocalRoutes_Static_NextHop_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_NextHop_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLocalRoutes_Static_NextHop_InterfaceRef_InterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface to the batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutes_Static_NextHop_InterfaceRef_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", structs[pre], queryPath, true, false)
			qv := convertLocalRoutes_Static_NextHop_InterfaceRef_InterfacePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LocalRoutes_Static_NextHop_InterfaceRef_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/state/interface to the batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLocalRoutes_Static_NextHop_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.LocalRoutes_Static_NextHop_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLocalRoutes_Static_NextHop_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static_NextHop_InterfaceRef) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Interface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
