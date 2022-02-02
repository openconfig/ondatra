package gnmicollectormetadata

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

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetLeavesAdded with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetLeavesAddedPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, false)
	if ok {
		return convertMeta_TargetLeavesAddedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetLeavesAdded with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetLeavesAddedPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetLeavesAdded with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetLeavesAddedPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertMeta_TargetLeavesAddedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetLeavesAdded with a ONCE subscription.
func (n *Meta_TargetLeavesAddedPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/targetLeavesAdded with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_TargetLeavesAddedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_TargetLeavesAddedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	gs := &oc.Meta{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Meta", gs, queryPath, true, false)
		return convertMeta_TargetLeavesAddedPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/targetLeavesAdded with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_TargetLeavesAddedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_TargetLeavesAddedPath(t, n, timeout, predicate)
}

// Await observes values at /gnmi-collector-metadata/meta/targetLeavesAdded with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Meta_TargetLeavesAddedPath) Await(t testing.TB, timeout time.Duration, val int64) *oc.QualifiedInt64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /gnmi-collector-metadata/meta/targetLeavesAdded failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /gnmi-collector-metadata/meta/targetLeavesAdded to the batch object.
func (n *Meta_TargetLeavesAddedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/targetLeavesAdded with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_TargetLeavesAddedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/targetLeavesAdded with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_TargetLeavesAddedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_TargetLeavesAddedPath(t, n, timeout, predicate)
}

// Batch adds /gnmi-collector-metadata/meta/targetLeavesAdded to the batch object.
func (n *Meta_TargetLeavesAddedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertMeta_TargetLeavesAddedPath extracts the value of the leaf TargetLeavesAdded from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetLeavesAddedPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetLeavesAdded
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetLeavesDeleted with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetLeavesDeletedPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, false)
	if ok {
		return convertMeta_TargetLeavesDeletedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetLeavesDeleted with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetLeavesDeletedPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetLeavesDeleted with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetLeavesDeletedPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertMeta_TargetLeavesDeletedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetLeavesDeleted with a ONCE subscription.
func (n *Meta_TargetLeavesDeletedPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/targetLeavesDeleted with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_TargetLeavesDeletedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_TargetLeavesDeletedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	gs := &oc.Meta{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Meta", gs, queryPath, true, false)
		return convertMeta_TargetLeavesDeletedPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/targetLeavesDeleted with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_TargetLeavesDeletedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_TargetLeavesDeletedPath(t, n, timeout, predicate)
}

// Await observes values at /gnmi-collector-metadata/meta/targetLeavesDeleted with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Meta_TargetLeavesDeletedPath) Await(t testing.TB, timeout time.Duration, val int64) *oc.QualifiedInt64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /gnmi-collector-metadata/meta/targetLeavesDeleted failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /gnmi-collector-metadata/meta/targetLeavesDeleted to the batch object.
func (n *Meta_TargetLeavesDeletedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/targetLeavesDeleted with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_TargetLeavesDeletedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/targetLeavesDeleted with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_TargetLeavesDeletedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_TargetLeavesDeletedPath(t, n, timeout, predicate)
}

// Batch adds /gnmi-collector-metadata/meta/targetLeavesDeleted to the batch object.
func (n *Meta_TargetLeavesDeletedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertMeta_TargetLeavesDeletedPath extracts the value of the leaf TargetLeavesDeleted from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetLeavesDeletedPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetLeavesDeleted
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetLeavesEmpty with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetLeavesEmptyPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, false)
	if ok {
		return convertMeta_TargetLeavesEmptyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetLeavesEmpty with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetLeavesEmptyPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetLeavesEmpty with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetLeavesEmptyPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertMeta_TargetLeavesEmptyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetLeavesEmpty with a ONCE subscription.
func (n *Meta_TargetLeavesEmptyPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/targetLeavesEmpty with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_TargetLeavesEmptyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_TargetLeavesEmptyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	gs := &oc.Meta{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Meta", gs, queryPath, true, false)
		return convertMeta_TargetLeavesEmptyPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/targetLeavesEmpty with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_TargetLeavesEmptyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_TargetLeavesEmptyPath(t, n, timeout, predicate)
}

// Await observes values at /gnmi-collector-metadata/meta/targetLeavesEmpty with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Meta_TargetLeavesEmptyPath) Await(t testing.TB, timeout time.Duration, val int64) *oc.QualifiedInt64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /gnmi-collector-metadata/meta/targetLeavesEmpty failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /gnmi-collector-metadata/meta/targetLeavesEmpty to the batch object.
func (n *Meta_TargetLeavesEmptyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/targetLeavesEmpty with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_TargetLeavesEmptyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/targetLeavesEmpty with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_TargetLeavesEmptyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_TargetLeavesEmptyPath(t, n, timeout, predicate)
}

// Batch adds /gnmi-collector-metadata/meta/targetLeavesEmpty to the batch object.
func (n *Meta_TargetLeavesEmptyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertMeta_TargetLeavesEmptyPath extracts the value of the leaf TargetLeavesEmpty from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetLeavesEmptyPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetLeavesEmpty
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
