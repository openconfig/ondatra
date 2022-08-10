package platform

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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PostFecBer_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PostFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PostFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PostFecBer_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PostFecBer_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PostFecBer_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PostFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PostFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PostFecBer_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time with a ONCE subscription.
func (n *Component_Transceiver_PostFecBer_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PostFecBer_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PostFecBer_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time to the batch object.
func (n *Component_Transceiver_PostFecBer_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PostFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PostFecBer", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PostFecBer_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/max-time to the batch object.
func (n *Component_Transceiver_PostFecBer_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PostFecBer_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_Transceiver_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_PostFecBer_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PostFecBer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PostFecBer_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PostFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PostFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PostFecBer_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PostFecBer_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PostFecBer_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PostFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PostFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PostFecBer_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min with a ONCE subscription.
func (n *Component_Transceiver_PostFecBer_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PostFecBer_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PostFecBer_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/min to the batch object.
func (n *Component_Transceiver_PostFecBer_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PostFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PostFecBer", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PostFecBer_MinPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/min to the batch object.
func (n *Component_Transceiver_PostFecBer_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PostFecBer_MinPath extracts the value of the leaf Min from its parent oc.Component_Transceiver_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_PostFecBer_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PostFecBer) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Min
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PostFecBer_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PostFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PostFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PostFecBer_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PostFecBer_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PostFecBer_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PostFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PostFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PostFecBer_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time with a ONCE subscription.
func (n *Component_Transceiver_PostFecBer_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PostFecBer_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PostFecBer_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time to the batch object.
func (n *Component_Transceiver_PostFecBer_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PostFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PostFecBer", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PostFecBer_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/min-time to the batch object.
func (n *Component_Transceiver_PostFecBer_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PostFecBer_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_Transceiver_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_PostFecBer_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PostFecBer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MinTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PreFecBerPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver_PreFecBer {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PreFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PreFecBer", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Transceiver_PreFecBer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PreFecBerPath) Get(t testing.TB) *oc.Component_Transceiver_PreFecBer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PreFecBerPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver_PreFecBer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver_PreFecBer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PreFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PreFecBer", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver_PreFecBer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber with a ONCE subscription.
func (n *Component_Transceiver_PreFecBerPathAny) Get(t testing.TB) []*oc.Component_Transceiver_PreFecBer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver_PreFecBer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_PreFecBer {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_PreFecBer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_PreFecBer) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Transceiver_PreFecBer{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Transceiver_PreFecBer)))
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_PreFecBer) bool) *oc.Component_Transceiver_PreFecBerWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_PreFecBerWatcher{}
	gs := &oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PreFecBer", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_Transceiver_PreFecBer{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_PreFecBer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_PreFecBer) bool) *oc.Component_Transceiver_PreFecBerWatcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PreFecBerPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Transceiver_PreFecBer) *oc.QualifiedComponent_Transceiver_PreFecBer {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Transceiver_PreFecBer) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/pre-fec-ber failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber to the batch object.
func (n *Component_Transceiver_PreFecBerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_PreFecBer {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_PreFecBer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_PreFecBer) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_PreFecBer) bool) *oc.Component_Transceiver_PreFecBerWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_PreFecBerWatcher{}
	structs := map[string]*oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PreFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PreFecBer", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_Transceiver_PreFecBer{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_PreFecBer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_PreFecBer) bool) *oc.Component_Transceiver_PreFecBerWatcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber to the batch object.
func (n *Component_Transceiver_PreFecBerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PreFecBer_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PreFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PreFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PreFecBer_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PreFecBer_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PreFecBer_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PreFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PreFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PreFecBer_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg with a ONCE subscription.
func (n *Component_Transceiver_PreFecBer_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PreFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PreFecBer_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PreFecBer_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg to the batch object.
func (n *Component_Transceiver_PreFecBer_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PreFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PreFecBer", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PreFecBer_AvgPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/avg to the batch object.
func (n *Component_Transceiver_PreFecBer_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PreFecBer_AvgPath extracts the value of the leaf Avg from its parent oc.Component_Transceiver_PreFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_PreFecBer_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PreFecBer) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PreFecBer_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PreFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PreFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PreFecBer_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PreFecBer_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PreFecBer_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PreFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PreFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PreFecBer_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant with a ONCE subscription.
func (n *Component_Transceiver_PreFecBer_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PreFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PreFecBer_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PreFecBer_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant to the batch object.
func (n *Component_Transceiver_PreFecBer_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PreFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PreFecBer", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PreFecBer_InstantPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/instant to the batch object.
func (n *Component_Transceiver_PreFecBer_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PreFecBer_InstantPath extracts the value of the leaf Instant from its parent oc.Component_Transceiver_PreFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_PreFecBer_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PreFecBer) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Instant
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PreFecBer_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PreFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PreFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PreFecBer_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PreFecBer_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PreFecBer_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PreFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PreFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PreFecBer_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval with a ONCE subscription.
func (n *Component_Transceiver_PreFecBer_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PreFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PreFecBer_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PreFecBer_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval to the batch object.
func (n *Component_Transceiver_PreFecBer_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PreFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PreFecBer", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PreFecBer_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/interval to the batch object.
func (n *Component_Transceiver_PreFecBer_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PreFecBer_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_Transceiver_PreFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_PreFecBer_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PreFecBer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PreFecBer_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PreFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PreFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PreFecBer_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PreFecBer_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PreFecBer_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PreFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PreFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PreFecBer_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max with a ONCE subscription.
func (n *Component_Transceiver_PreFecBer_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PreFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PreFecBer_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PreFecBer_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max to the batch object.
func (n *Component_Transceiver_PreFecBer_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PreFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PreFecBer", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PreFecBer_MaxPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max to the batch object.
func (n *Component_Transceiver_PreFecBer_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PreFecBer_MaxPath extracts the value of the leaf Max from its parent oc.Component_Transceiver_PreFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_PreFecBer_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PreFecBer) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Max
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PreFecBer_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PreFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PreFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PreFecBer_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PreFecBer_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PreFecBer_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PreFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PreFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PreFecBer_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time with a ONCE subscription.
func (n *Component_Transceiver_PreFecBer_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PreFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PreFecBer_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PreFecBer_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time to the batch object.
func (n *Component_Transceiver_PreFecBer_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PreFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PreFecBer", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PreFecBer_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/max-time to the batch object.
func (n *Component_Transceiver_PreFecBer_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PreFecBer_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_Transceiver_PreFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_PreFecBer_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PreFecBer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PreFecBer_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PreFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PreFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PreFecBer_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PreFecBer_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PreFecBer_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PreFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PreFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PreFecBer_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min with a ONCE subscription.
func (n *Component_Transceiver_PreFecBer_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PreFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PreFecBer_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PreFecBer_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min to the batch object.
func (n *Component_Transceiver_PreFecBer_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PreFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PreFecBer", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PreFecBer_MinPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min to the batch object.
func (n *Component_Transceiver_PreFecBer_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PreFecBer_MinPath extracts the value of the leaf Min from its parent oc.Component_Transceiver_PreFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_PreFecBer_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PreFecBer) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Min
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PreFecBer_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PreFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PreFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PreFecBer_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PreFecBer_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PreFecBer_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PreFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PreFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PreFecBer_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time with a ONCE subscription.
func (n *Component_Transceiver_PreFecBer_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PreFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PreFecBer_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PreFecBer_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time to the batch object.
func (n *Component_Transceiver_PreFecBer_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PreFecBer_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PreFecBer_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_PreFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_PreFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PreFecBer", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PreFecBer_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PreFecBer_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PreFecBer_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/pre-fec-ber/min-time to the batch object.
func (n *Component_Transceiver_PreFecBer_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PreFecBer_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_Transceiver_PreFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_PreFecBer_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PreFecBer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MinTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/present with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PresentPath) Lookup(t testing.TB) *oc.QualifiedE_Transceiver_Present {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PresentPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/present with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PresentPath) Get(t testing.TB) oc.E_Transceiver_Present {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/present with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PresentPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Transceiver_Present {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Transceiver_Present
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_PresentPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/present with a ONCE subscription.
func (n *Component_Transceiver_PresentPathAny) Get(t testing.TB) []oc.E_Transceiver_Present {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Transceiver_Present
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/present with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PresentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Transceiver_Present {
	t.Helper()
	c := &oc.CollectionE_Transceiver_Present{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Transceiver_Present) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PresentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Transceiver_Present) bool) *oc.E_Transceiver_PresentWatcher {
	t.Helper()
	w := &oc.E_Transceiver_PresentWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PresentPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Transceiver_Present)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/present with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PresentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Transceiver_Present) bool) *oc.E_Transceiver_PresentWatcher {
	t.Helper()
	return watch_Component_Transceiver_PresentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/present with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PresentPath) Await(t testing.TB, timeout time.Duration, val oc.E_Transceiver_Present) *oc.QualifiedE_Transceiver_Present {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Transceiver_Present) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/present failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/present to the batch object.
func (n *Component_Transceiver_PresentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/present with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PresentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Transceiver_Present {
	t.Helper()
	c := &oc.CollectionE_Transceiver_Present{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Transceiver_Present) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PresentPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Transceiver_Present) bool) *oc.E_Transceiver_PresentWatcher {
	t.Helper()
	w := &oc.E_Transceiver_PresentWatcher{}
	structs := map[string]*oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_PresentPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Transceiver_Present)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/present with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PresentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Transceiver_Present) bool) *oc.E_Transceiver_PresentWatcher {
	t.Helper()
	return watch_Component_Transceiver_PresentPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/present to the batch object.
func (n *Component_Transceiver_PresentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PresentPath extracts the value of the leaf Present from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Transceiver_Present.
func convertComponent_Transceiver_PresentPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_Transceiver_Present {
	t.Helper()
	qv := &oc.QualifiedE_Transceiver_Present{
		Metadata: md,
	}
	val := parent.Present
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/serial-no with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_SerialNoPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_SerialNoPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/serial-no with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_SerialNoPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/serial-no with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_SerialNoPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_SerialNoPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/serial-no with a ONCE subscription.
func (n *Component_Transceiver_SerialNoPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/serial-no with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SerialNoPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SerialNoPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_SerialNoPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/serial-no with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SerialNoPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_SerialNoPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/serial-no with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_SerialNoPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/serial-no failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/serial-no to the batch object.
func (n *Component_Transceiver_SerialNoPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/serial-no with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SerialNoPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SerialNoPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_SerialNoPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/serial-no with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SerialNoPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_SerialNoPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/serial-no to the batch object.
func (n *Component_Transceiver_SerialNoPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_SerialNoPath extracts the value of the leaf SerialNo from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Transceiver_SerialNoPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SerialNo
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_SonetSdhComplianceCodePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_SonetSdhComplianceCodePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_SonetSdhComplianceCodePath) Get(t testing.TB) oc.E_TransportTypes_SONET_APPLICATION_CODE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_SonetSdhComplianceCodePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_SonetSdhComplianceCodePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code with a ONCE subscription.
func (n *Component_Transceiver_SonetSdhComplianceCodePathAny) Get(t testing.TB) []oc.E_TransportTypes_SONET_APPLICATION_CODE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_SONET_APPLICATION_CODE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SonetSdhComplianceCodePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_SONET_APPLICATION_CODE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_SONET_APPLICATION_CODE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SonetSdhComplianceCodePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE) bool) *oc.E_TransportTypes_SONET_APPLICATION_CODEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_SONET_APPLICATION_CODEWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_SonetSdhComplianceCodePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SonetSdhComplianceCodePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE) bool) *oc.E_TransportTypes_SONET_APPLICATION_CODEWatcher {
	t.Helper()
	return watch_Component_Transceiver_SonetSdhComplianceCodePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_SonetSdhComplianceCodePath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_SONET_APPLICATION_CODE) *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code to the batch object.
func (n *Component_Transceiver_SonetSdhComplianceCodePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SonetSdhComplianceCodePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_SONET_APPLICATION_CODE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_SONET_APPLICATION_CODE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SonetSdhComplianceCodePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE) bool) *oc.E_TransportTypes_SONET_APPLICATION_CODEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_SONET_APPLICATION_CODEWatcher{}
	structs := map[string]*oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_SonetSdhComplianceCodePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SonetSdhComplianceCodePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE) bool) *oc.E_TransportTypes_SONET_APPLICATION_CODEWatcher {
	t.Helper()
	return watch_Component_Transceiver_SonetSdhComplianceCodePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/sonet-sdh-compliance-code to the batch object.
func (n *Component_Transceiver_SonetSdhComplianceCodePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_SonetSdhComplianceCodePath extracts the value of the leaf SonetSdhComplianceCode from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE.
func convertComponent_Transceiver_SonetSdhComplianceCodePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_SONET_APPLICATION_CODE{
		Metadata: md,
	}
	val := parent.SonetSdhComplianceCode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_SupplyVoltagePath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver_SupplyVoltage {
	t.Helper()
	goStruct := &oc.Component_Transceiver_SupplyVoltage{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_SupplyVoltage", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Transceiver_SupplyVoltage{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_SupplyVoltagePath) Get(t testing.TB) *oc.Component_Transceiver_SupplyVoltage {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_SupplyVoltagePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver_SupplyVoltage {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver_SupplyVoltage
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_SupplyVoltage{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver_SupplyVoltage{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage with a ONCE subscription.
func (n *Component_Transceiver_SupplyVoltagePathAny) Get(t testing.TB) []*oc.Component_Transceiver_SupplyVoltage {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver_SupplyVoltage
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltagePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_SupplyVoltage {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_SupplyVoltage{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_SupplyVoltage) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Transceiver_SupplyVoltage{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Transceiver_SupplyVoltage)))
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltagePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_SupplyVoltage) bool) *oc.Component_Transceiver_SupplyVoltageWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_SupplyVoltageWatcher{}
	gs := &oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_SupplyVoltage", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_Transceiver_SupplyVoltage{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_SupplyVoltage)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltagePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_SupplyVoltage) bool) *oc.Component_Transceiver_SupplyVoltageWatcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltagePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/supply-voltage with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_SupplyVoltagePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Transceiver_SupplyVoltage) *oc.QualifiedComponent_Transceiver_SupplyVoltage {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Transceiver_SupplyVoltage) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/supply-voltage failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage to the batch object.
func (n *Component_Transceiver_SupplyVoltagePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltagePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_SupplyVoltage {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_SupplyVoltage{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_SupplyVoltage) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltagePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_SupplyVoltage) bool) *oc.Component_Transceiver_SupplyVoltageWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_SupplyVoltageWatcher{}
	structs := map[string]*oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_SupplyVoltage{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_Transceiver_SupplyVoltage{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_SupplyVoltage)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltagePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_SupplyVoltage) bool) *oc.Component_Transceiver_SupplyVoltageWatcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltagePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage to the batch object.
func (n *Component_Transceiver_SupplyVoltagePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_SupplyVoltage_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_SupplyVoltage{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_SupplyVoltage", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_SupplyVoltage_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_SupplyVoltage_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_SupplyVoltage_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_SupplyVoltage{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_SupplyVoltage_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/avg with a ONCE subscription.
func (n *Component_Transceiver_SupplyVoltage_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_SupplyVoltage", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_SupplyVoltage_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/supply-voltage/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_SupplyVoltage_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/supply-voltage/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/avg to the batch object.
func (n *Component_Transceiver_SupplyVoltage_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_SupplyVoltage{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_SupplyVoltage_AvgPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/avg to the batch object.
func (n *Component_Transceiver_SupplyVoltage_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_SupplyVoltage_AvgPath extracts the value of the leaf Avg from its parent oc.Component_Transceiver_SupplyVoltage
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_SupplyVoltage_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_SupplyVoltage) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_SupplyVoltage_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_SupplyVoltage{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_SupplyVoltage", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_SupplyVoltage_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_SupplyVoltage_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_SupplyVoltage_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_SupplyVoltage{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_SupplyVoltage_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/instant with a ONCE subscription.
func (n *Component_Transceiver_SupplyVoltage_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_SupplyVoltage", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_SupplyVoltage_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/supply-voltage/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_SupplyVoltage_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/supply-voltage/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/instant to the batch object.
func (n *Component_Transceiver_SupplyVoltage_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_SupplyVoltage{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_SupplyVoltage_InstantPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/instant to the batch object.
func (n *Component_Transceiver_SupplyVoltage_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_SupplyVoltage_InstantPath extracts the value of the leaf Instant from its parent oc.Component_Transceiver_SupplyVoltage
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_SupplyVoltage_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_SupplyVoltage) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Instant
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_SupplyVoltage_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_SupplyVoltage{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_SupplyVoltage", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_SupplyVoltage_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_SupplyVoltage_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_SupplyVoltage_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_SupplyVoltage{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_SupplyVoltage_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/interval with a ONCE subscription.
func (n *Component_Transceiver_SupplyVoltage_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_SupplyVoltage", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_SupplyVoltage_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/supply-voltage/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_SupplyVoltage_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/supply-voltage/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/interval to the batch object.
func (n *Component_Transceiver_SupplyVoltage_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_SupplyVoltage{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_SupplyVoltage_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/interval to the batch object.
func (n *Component_Transceiver_SupplyVoltage_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_SupplyVoltage_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_Transceiver_SupplyVoltage
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_SupplyVoltage_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_SupplyVoltage) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_SupplyVoltage_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_SupplyVoltage{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_SupplyVoltage", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_SupplyVoltage_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_SupplyVoltage_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_SupplyVoltage_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_SupplyVoltage{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_SupplyVoltage_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max with a ONCE subscription.
func (n *Component_Transceiver_SupplyVoltage_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_SupplyVoltage", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_SupplyVoltage_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_SupplyVoltage_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/supply-voltage/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/max to the batch object.
func (n *Component_Transceiver_SupplyVoltage_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_SupplyVoltage{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_SupplyVoltage_MaxPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/max to the batch object.
func (n *Component_Transceiver_SupplyVoltage_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_SupplyVoltage_MaxPath extracts the value of the leaf Max from its parent oc.Component_Transceiver_SupplyVoltage
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_SupplyVoltage_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_SupplyVoltage) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Max
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_SupplyVoltage{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_SupplyVoltage", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_SupplyVoltage_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_SupplyVoltage{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_SupplyVoltage_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time with a ONCE subscription.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_SupplyVoltage", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_SupplyVoltage_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time to the batch object.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_SupplyVoltage{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_SupplyVoltage_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/max-time to the batch object.
func (n *Component_Transceiver_SupplyVoltage_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_SupplyVoltage_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_Transceiver_SupplyVoltage
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_SupplyVoltage_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_SupplyVoltage) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_SupplyVoltage_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_SupplyVoltage{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_SupplyVoltage", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_SupplyVoltage_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_SupplyVoltage_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_SupplyVoltage_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_SupplyVoltage{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_SupplyVoltage_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min with a ONCE subscription.
func (n *Component_Transceiver_SupplyVoltage_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_SupplyVoltage", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_SupplyVoltage_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_SupplyVoltage_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/supply-voltage/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/min to the batch object.
func (n *Component_Transceiver_SupplyVoltage_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_SupplyVoltage{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_SupplyVoltage_MinPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/min to the batch object.
func (n *Component_Transceiver_SupplyVoltage_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_SupplyVoltage_MinPath extracts the value of the leaf Min from its parent oc.Component_Transceiver_SupplyVoltage
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_SupplyVoltage_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_SupplyVoltage) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Min
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_SupplyVoltage_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_SupplyVoltage{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_SupplyVoltage", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_SupplyVoltage_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_SupplyVoltage_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_SupplyVoltage_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_SupplyVoltage{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_SupplyVoltage_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time with a ONCE subscription.
func (n *Component_Transceiver_SupplyVoltage_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_SupplyVoltage", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_SupplyVoltage_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_SupplyVoltage_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time to the batch object.
func (n *Component_Transceiver_SupplyVoltage_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_SupplyVoltage_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_SupplyVoltage_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_SupplyVoltage{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_SupplyVoltage{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_SupplyVoltage", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_SupplyVoltage_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_SupplyVoltage_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_SupplyVoltage_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/supply-voltage/min-time to the batch object.
func (n *Component_Transceiver_SupplyVoltage_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_SupplyVoltage_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_Transceiver_SupplyVoltage
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_SupplyVoltage_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_SupplyVoltage) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MinTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/vendor-part with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_VendorPartPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_VendorPartPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/vendor-part with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_VendorPartPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/vendor-part with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_VendorPartPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_VendorPartPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/vendor-part with a ONCE subscription.
func (n *Component_Transceiver_VendorPartPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/vendor-part with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_VendorPartPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_VendorPartPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_VendorPartPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/vendor-part with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_VendorPartPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_VendorPartPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/vendor-part with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_VendorPartPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/vendor-part failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/vendor-part to the batch object.
func (n *Component_Transceiver_VendorPartPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/vendor-part with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_VendorPartPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_VendorPartPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_VendorPartPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/vendor-part with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_VendorPartPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_VendorPartPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/vendor-part to the batch object.
func (n *Component_Transceiver_VendorPartPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_VendorPartPath extracts the value of the leaf VendorPart from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Transceiver_VendorPartPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.VendorPart
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/vendor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_VendorPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_VendorPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/vendor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_VendorPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/vendor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_VendorPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_VendorPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/vendor with a ONCE subscription.
func (n *Component_Transceiver_VendorPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/vendor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_VendorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_VendorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_VendorPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/vendor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_VendorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_VendorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/vendor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_VendorPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/vendor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/vendor to the batch object.
func (n *Component_Transceiver_VendorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/vendor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_VendorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_VendorPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_VendorPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/vendor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_VendorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_VendorPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/vendor to the batch object.
func (n *Component_Transceiver_VendorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_VendorPath extracts the value of the leaf Vendor from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Transceiver_VendorPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Vendor
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/vendor-rev with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_VendorRevPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_VendorRevPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/vendor-rev with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_VendorRevPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/vendor-rev with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_VendorRevPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_VendorRevPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/vendor-rev with a ONCE subscription.
func (n *Component_Transceiver_VendorRevPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/vendor-rev with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_VendorRevPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_VendorRevPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_VendorRevPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/vendor-rev with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_VendorRevPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_VendorRevPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/vendor-rev with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_VendorRevPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/vendor-rev failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/vendor-rev to the batch object.
func (n *Component_Transceiver_VendorRevPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/vendor-rev with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_VendorRevPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_VendorRevPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_VendorRevPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/vendor-rev with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_VendorRevPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Transceiver_VendorRevPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/vendor-rev to the batch object.
func (n *Component_Transceiver_VendorRevPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_VendorRevPath extracts the value of the leaf VendorRev from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Transceiver_VendorRevPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.VendorRev
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_TypePath) Lookup(t testing.TB) *oc.QualifiedComponent_Type_Union {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_TypePath) Get(t testing.TB) oc.Component_Type_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Type_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Type_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/type with a ONCE subscription.
func (n *Component_TypePathAny) Get(t testing.TB) []oc.Component_Type_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Component_Type_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_TypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Type_Union {
	t.Helper()
	c := &oc.CollectionComponent_Type_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Type_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_TypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Type_Union) bool) *oc.Component_Type_UnionWatcher {
	t.Helper()
	w := &oc.Component_Type_UnionWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_TypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Type_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_TypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Type_Union) bool) *oc.Component_Type_UnionWatcher {
	t.Helper()
	return watch_Component_TypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_TypePath) Await(t testing.TB, timeout time.Duration, val oc.Component_Type_Union) *oc.QualifiedComponent_Type_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Type_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/type to the batch object.
func (n *Component_TypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_TypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Type_Union {
	t.Helper()
	c := &oc.CollectionComponent_Type_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Type_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_TypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Type_Union) bool) *oc.Component_Type_UnionWatcher {
	t.Helper()
	w := &oc.Component_Type_UnionWatcher{}
	structs := map[string]*oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component", structs[pre], queryPath, true, false)
			qv := convertComponent_TypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Type_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Type_Union) bool) *oc.Component_Type_UnionWatcher {
	t.Helper()
	return watch_Component_TypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/type to the batch object.
func (n *Component_TypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_TypePath extracts the value of the leaf Type from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedComponent_Type_Union.
func convertComponent_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedComponent_Type_Union {
	t.Helper()
	qv := &oc.QualifiedComponent_Type_Union{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/used-power with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_UsedPowerPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_UsedPowerPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/used-power with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_UsedPowerPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/used-power with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_UsedPowerPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_UsedPowerPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/used-power with a ONCE subscription.
func (n *Component_UsedPowerPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/used-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_UsedPowerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_UsedPowerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_UsedPowerPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/used-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_UsedPowerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Component_UsedPowerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/used-power with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_UsedPowerPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/used-power failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/used-power to the batch object.
func (n *Component_UsedPowerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/used-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_UsedPowerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_UsedPowerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component", structs[pre], queryPath, true, false)
			qv := convertComponent_UsedPowerPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/used-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_UsedPowerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Component_UsedPowerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/used-power to the batch object.
func (n *Component_UsedPowerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_UsedPowerPath extracts the value of the leaf UsedPower from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertComponent_UsedPowerPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.UsedPower
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
