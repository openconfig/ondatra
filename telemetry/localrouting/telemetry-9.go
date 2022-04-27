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

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/state/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_PrefixPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static", goStruct, true, false)
	if ok {
		return convertLocalRoutes_Static_PrefixPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/state/prefix with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_PrefixPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/state/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_PrefixPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/state/prefix with a ONCE subscription.
func (n *LocalRoutes_Static_PrefixPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/state/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_PrefixPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutes_Static_PrefixPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LocalRoutes_Static{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LocalRoutes_Static", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLocalRoutes_Static_PrefixPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/state/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_PrefixPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LocalRoutes_Static_PrefixPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-local-routing/local-routes/static-routes/static/state/prefix with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LocalRoutes_Static_PrefixPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-local-routing/local-routes/static-routes/static/state/prefix failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/state/prefix to the batch object.
func (n *LocalRoutes_Static_PrefixPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/state/prefix with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_PrefixPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutes_Static_PrefixPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.LocalRoutes_Static{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LocalRoutes_Static{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LocalRoutes_Static", structs[pre], queryPath, true, false)
			qv := convertLocalRoutes_Static_PrefixPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/state/prefix with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_PrefixPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LocalRoutes_Static_PrefixPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/state/prefix to the batch object.
func (n *LocalRoutes_Static_PrefixPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLocalRoutes_Static_PrefixPath extracts the value of the leaf Prefix from its parent oc.LocalRoutes_Static
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLocalRoutes_Static_PrefixPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/state/set-tag with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_SetTagPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Static_SetTag_Union {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static", goStruct, true, false)
	if ok {
		return convertLocalRoutes_Static_SetTagPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/state/set-tag with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_SetTagPath) Get(t testing.TB) oc.LocalRoutes_Static_SetTag_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/state/set-tag with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_SetTagPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Static_SetTag_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Static_SetTag_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_SetTagPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/state/set-tag with a ONCE subscription.
func (n *LocalRoutes_Static_SetTagPathAny) Get(t testing.TB) []oc.LocalRoutes_Static_SetTag_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.LocalRoutes_Static_SetTag_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/state/set-tag with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_SetTagPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLocalRoutes_Static_SetTag_Union {
	t.Helper()
	c := &oc.CollectionLocalRoutes_Static_SetTag_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLocalRoutes_Static_SetTag_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutes_Static_SetTagPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_SetTag_Union) bool) *oc.LocalRoutes_Static_SetTag_UnionWatcher {
	t.Helper()
	w := &oc.LocalRoutes_Static_SetTag_UnionWatcher{}
	gs := &oc.LocalRoutes_Static{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LocalRoutes_Static", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLocalRoutes_Static_SetTagPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLocalRoutes_Static_SetTag_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/state/set-tag with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_SetTagPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_SetTag_Union) bool) *oc.LocalRoutes_Static_SetTag_UnionWatcher {
	t.Helper()
	return watch_LocalRoutes_Static_SetTagPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-local-routing/local-routes/static-routes/static/state/set-tag with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LocalRoutes_Static_SetTagPath) Await(t testing.TB, timeout time.Duration, val oc.LocalRoutes_Static_SetTag_Union) *oc.QualifiedLocalRoutes_Static_SetTag_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLocalRoutes_Static_SetTag_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-local-routing/local-routes/static-routes/static/state/set-tag failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/state/set-tag to the batch object.
func (n *LocalRoutes_Static_SetTagPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-local-routing/local-routes/static-routes/static/state/set-tag with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LocalRoutes_Static_SetTagPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLocalRoutes_Static_SetTag_Union {
	t.Helper()
	c := &oc.CollectionLocalRoutes_Static_SetTag_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLocalRoutes_Static_SetTag_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LocalRoutes_Static_SetTagPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_SetTag_Union) bool) *oc.LocalRoutes_Static_SetTag_UnionWatcher {
	t.Helper()
	w := &oc.LocalRoutes_Static_SetTag_UnionWatcher{}
	structs := map[string]*oc.LocalRoutes_Static{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LocalRoutes_Static{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LocalRoutes_Static", structs[pre], queryPath, true, false)
			qv := convertLocalRoutes_Static_SetTagPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLocalRoutes_Static_SetTag_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-local-routing/local-routes/static-routes/static/state/set-tag with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LocalRoutes_Static_SetTagPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLocalRoutes_Static_SetTag_Union) bool) *oc.LocalRoutes_Static_SetTag_UnionWatcher {
	t.Helper()
	return watch_LocalRoutes_Static_SetTagPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-local-routing/local-routes/static-routes/static/state/set-tag to the batch object.
func (n *LocalRoutes_Static_SetTagPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLocalRoutes_Static_SetTagPath extracts the value of the leaf SetTag from its parent oc.LocalRoutes_Static
// and combines the update with an existing Metadata to return a *oc.QualifiedLocalRoutes_Static_SetTag_Union.
func convertLocalRoutes_Static_SetTagPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static) *oc.QualifiedLocalRoutes_Static_SetTag_Union {
	t.Helper()
	qv := &oc.QualifiedLocalRoutes_Static_SetTag_Union{
		Metadata: md,
	}
	val := parent.SetTag
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
