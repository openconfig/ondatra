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

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop_EnableBfd{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop_EnableBfd", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Get(t testing.TB) *oc.LocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_EnableBfdPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop_EnableBfd{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop_EnableBfd", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_EnableBfdPathAny) Get(t testing.TB) []*oc.LocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LocalRoutes_Static_NextHop_EnableBfd
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	c := &oc.CollectionLocalRoutes_Static_NextHop_EnableBfd{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.LocalRoutes_Static_NextHop_EnableBfd)))
		return false
	})
	return c
}

func watch_LocalRoutes_Static_NextHop_EnableBfdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd) bool) *oc.LocalRoutes_Static_NextHop_EnableBfdWatcher {
	t.Helper()
	w := &oc.LocalRoutes_Static_NextHop_EnableBfdWatcher{}
	gs := &oc.LocalRoutes_Static_NextHop_EnableBfd{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LocalRoutes_Static_NextHop_EnableBfd", gs, queryPath, false, false)
		return (&oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd) bool) *oc.LocalRoutes_Static_NextHop_EnableBfdWatcher {
	t.Helper()
	return watch_LocalRoutes_Static_NextHop_EnableBfdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Await(t testing.TB, timeout time.Duration, val *oc.LocalRoutes_Static_NextHop_EnableBfd) *oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd to the batch object.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_NextHop_EnableBfdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	c := &oc.CollectionLocalRoutes_Static_NextHop_EnableBfd{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_NextHop_EnableBfdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd) bool) *oc.LocalRoutes_Static_NextHop_EnableBfdWatcher {
	t.Helper()
	return watch_LocalRoutes_Static_NextHop_EnableBfdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd to the batch object.
func (n *LocalRoutes_Static_NextHop_EnableBfdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
