package gnmicollectormetadata

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

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetSize with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetSizePath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, false)
	if ok {
		return convertMeta_TargetSizePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetSize with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetSizePath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetSize with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetSizePathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_TargetSizePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetSize with a ONCE subscription.
func (n *Meta_TargetSizePathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/targetSize with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_TargetSizePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_TargetSizePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	gs := &oc.Meta{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Meta", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertMeta_TargetSizePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/targetSize with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_TargetSizePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_TargetSizePath(t, n, timeout, predicate)
}

// Await observes values at /gnmi-collector-metadata/meta/targetSize with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Meta_TargetSizePath) Await(t testing.TB, timeout time.Duration, val int64) *oc.QualifiedInt64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /gnmi-collector-metadata/meta/targetSize failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /gnmi-collector-metadata/meta/targetSize to the batch object.
func (n *Meta_TargetSizePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/targetSize with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_TargetSizePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_TargetSizePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	structs := map[string]*oc.Meta{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Meta{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Meta", structs[pre], queryPath, true, false)
			qv := convertMeta_TargetSizePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/targetSize with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_TargetSizePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_TargetSizePathAny(t, n, timeout, predicate)
}

// Batch adds /gnmi-collector-metadata/meta/targetSize to the batch object.
func (n *Meta_TargetSizePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertMeta_TargetSizePath extracts the value of the leaf TargetSize from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetSizePath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetSize
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latency/window with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_WindowPath) Lookup(t testing.TB) *oc.QualifiedMeta_Window {
	t.Helper()
	goStruct := &oc.Meta_Window{}
	md, ok := oc.Lookup(t, n, "Meta_Window", goStruct, false, false)
	if ok {
		return (&oc.QualifiedMeta_Window{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latency/window with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_WindowPath) Get(t testing.TB) *oc.Meta_Window {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latency/window with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_WindowPathAny) Lookup(t testing.TB) []*oc.QualifiedMeta_Window {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedMeta_Window
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta_Window{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta_Window", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedMeta_Window{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latency/window with a ONCE subscription.
func (n *Meta_WindowPathAny) Get(t testing.TB) []*oc.Meta_Window {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Meta_Window
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/latency/window with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_WindowPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionMeta_Window {
	t.Helper()
	c := &oc.CollectionMeta_Window{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedMeta_Window) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedMeta_Window{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Meta_Window)))
		return false
	})
	return c
}

func watch_Meta_WindowPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedMeta_Window) bool) *oc.Meta_WindowWatcher {
	t.Helper()
	w := &oc.Meta_WindowWatcher{}
	gs := &oc.Meta_Window{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Meta_Window", gs, queryPath, false, false)
		qv := (&oc.QualifiedMeta_Window{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedMeta_Window)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/latency/window with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_WindowPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedMeta_Window) bool) *oc.Meta_WindowWatcher {
	t.Helper()
	return watch_Meta_WindowPath(t, n, timeout, predicate)
}

// Await observes values at /gnmi-collector-metadata/meta/latency/window with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Meta_WindowPath) Await(t testing.TB, timeout time.Duration, val *oc.Meta_Window) *oc.QualifiedMeta_Window {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedMeta_Window) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /gnmi-collector-metadata/meta/latency/window failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /gnmi-collector-metadata/meta/latency/window to the batch object.
func (n *Meta_WindowPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/latency/window with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_WindowPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionMeta_Window {
	t.Helper()
	c := &oc.CollectionMeta_Window{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedMeta_Window) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_WindowPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedMeta_Window) bool) *oc.Meta_WindowWatcher {
	t.Helper()
	w := &oc.Meta_WindowWatcher{}
	structs := map[string]*oc.Meta_Window{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Meta_Window{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Meta_Window", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedMeta_Window{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedMeta_Window)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/latency/window with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_WindowPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedMeta_Window) bool) *oc.Meta_WindowWatcher {
	t.Helper()
	return watch_Meta_WindowPathAny(t, n, timeout, predicate)
}

// Batch adds /gnmi-collector-metadata/meta/latency/window to the batch object.
func (n *Meta_WindowPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latency/window/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_Window_AvgPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta_Window{}
	md, ok := oc.Lookup(t, n, "Meta_Window", goStruct, true, false)
	if ok {
		return convertMeta_Window_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latency/window/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_Window_AvgPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latency/window/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_Window_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta_Window{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta_Window", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertMeta_Window_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latency/window/avg with a ONCE subscription.
func (n *Meta_Window_AvgPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/latency/window/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_Window_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_Window_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	gs := &oc.Meta_Window{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Meta_Window", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertMeta_Window_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/latency/window/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_Window_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_Window_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /gnmi-collector-metadata/meta/latency/window/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Meta_Window_AvgPath) Await(t testing.TB, timeout time.Duration, val int64) *oc.QualifiedInt64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /gnmi-collector-metadata/meta/latency/window/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /gnmi-collector-metadata/meta/latency/window/avg to the batch object.
func (n *Meta_Window_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/latency/window/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_Window_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_Window_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	structs := map[string]*oc.Meta_Window{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Meta_Window{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Meta_Window", structs[pre], queryPath, true, false)
			qv := convertMeta_Window_AvgPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/latency/window/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_Window_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_Window_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /gnmi-collector-metadata/meta/latency/window/avg to the batch object.
func (n *Meta_Window_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertMeta_Window_AvgPath extracts the value of the leaf Avg from its parent oc.Meta_Window
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_Window_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta_Window) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latency/window/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_Window_MaxPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta_Window{}
	md, ok := oc.Lookup(t, n, "Meta_Window", goStruct, true, false)
	if ok {
		return convertMeta_Window_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latency/window/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_Window_MaxPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latency/window/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_Window_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta_Window{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta_Window", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertMeta_Window_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latency/window/max with a ONCE subscription.
func (n *Meta_Window_MaxPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/latency/window/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_Window_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_Window_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	gs := &oc.Meta_Window{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Meta_Window", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertMeta_Window_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/latency/window/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_Window_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_Window_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /gnmi-collector-metadata/meta/latency/window/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Meta_Window_MaxPath) Await(t testing.TB, timeout time.Duration, val int64) *oc.QualifiedInt64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /gnmi-collector-metadata/meta/latency/window/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /gnmi-collector-metadata/meta/latency/window/max to the batch object.
func (n *Meta_Window_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/latency/window/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_Window_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_Window_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	structs := map[string]*oc.Meta_Window{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Meta_Window{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Meta_Window", structs[pre], queryPath, true, false)
			qv := convertMeta_Window_MaxPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/latency/window/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_Window_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_Window_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /gnmi-collector-metadata/meta/latency/window/max to the batch object.
func (n *Meta_Window_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertMeta_Window_MaxPath extracts the value of the leaf Max from its parent oc.Meta_Window
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_Window_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta_Window) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.Max
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latency/window/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_Window_MinPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta_Window{}
	md, ok := oc.Lookup(t, n, "Meta_Window", goStruct, true, false)
	if ok {
		return convertMeta_Window_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latency/window/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_Window_MinPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latency/window/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_Window_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta_Window{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta_Window", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertMeta_Window_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latency/window/min with a ONCE subscription.
func (n *Meta_Window_MinPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/latency/window/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_Window_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_Window_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	gs := &oc.Meta_Window{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Meta_Window", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertMeta_Window_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/latency/window/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_Window_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_Window_MinPath(t, n, timeout, predicate)
}

// Await observes values at /gnmi-collector-metadata/meta/latency/window/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Meta_Window_MinPath) Await(t testing.TB, timeout time.Duration, val int64) *oc.QualifiedInt64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /gnmi-collector-metadata/meta/latency/window/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /gnmi-collector-metadata/meta/latency/window/min to the batch object.
func (n *Meta_Window_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/latency/window/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_Window_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_Window_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	structs := map[string]*oc.Meta_Window{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Meta_Window{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Meta_Window", structs[pre], queryPath, true, false)
			qv := convertMeta_Window_MinPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/latency/window/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_Window_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Meta_Window_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /gnmi-collector-metadata/meta/latency/window/min to the batch object.
func (n *Meta_Window_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertMeta_Window_MinPath extracts the value of the leaf Min from its parent oc.Meta_Window
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_Window_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta_Window) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.Min
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latency/window/state/size with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_Window_SizePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Meta_Window{}
	md, ok := oc.Lookup(t, n, "Meta_Window", goStruct, true, false)
	if ok {
		return convertMeta_Window_SizePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latency/window/state/size with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_Window_SizePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latency/window/state/size with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_Window_SizePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta_Window{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta_Window", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertMeta_Window_SizePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latency/window/state/size with a ONCE subscription.
func (n *Meta_Window_SizePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/latency/window/state/size with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_Window_SizePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_Window_SizePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Meta_Window{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Meta_Window", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertMeta_Window_SizePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/latency/window/state/size with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_Window_SizePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Meta_Window_SizePath(t, n, timeout, predicate)
}

// Await observes values at /gnmi-collector-metadata/meta/latency/window/state/size with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Meta_Window_SizePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /gnmi-collector-metadata/meta/latency/window/state/size failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /gnmi-collector-metadata/meta/latency/window/state/size to the batch object.
func (n *Meta_Window_SizePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /gnmi-collector-metadata/meta/latency/window/state/size with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Meta_Window_SizePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Meta_Window_SizePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Meta_Window{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Meta_Window{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Meta_Window", structs[pre], queryPath, true, false)
			qv := convertMeta_Window_SizePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /gnmi-collector-metadata/meta/latency/window/state/size with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Meta_Window_SizePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Meta_Window_SizePathAny(t, n, timeout, predicate)
}

// Batch adds /gnmi-collector-metadata/meta/latency/window/state/size to the batch object.
func (n *Meta_Window_SizePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertMeta_Window_SizePath extracts the value of the leaf Size from its parent oc.Meta_Window
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertMeta_Window_SizePath(t testing.TB, md *genutil.Metadata, parent *oc.Meta_Window) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Size
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
