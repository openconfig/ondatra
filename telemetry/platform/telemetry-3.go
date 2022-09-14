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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXqPath) Lookup(t testing.TB) *oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasXq", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasXqPath) Get(t testing.TB) *oc.Component_OpticalChannel_ModulatorBiasXq {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXqPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasXqPathAny) Get(t testing.TB) []*oc.Component_OpticalChannel_ModulatorBiasXq {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_OpticalChannel_ModulatorBiasXq
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXqPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_ModulatorBiasXq {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_ModulatorBiasXq{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_OpticalChannel_ModulatorBiasXq)))
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXqPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq) bool) *oc.Component_OpticalChannel_ModulatorBiasXqWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_ModulatorBiasXqWatcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXqPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq) bool) *oc.Component_OpticalChannel_ModulatorBiasXqWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXqPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasXqPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_OpticalChannel_ModulatorBiasXq) *oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXqPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXqPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_ModulatorBiasXq {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_ModulatorBiasXq{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXqPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq) bool) *oc.Component_OpticalChannel_ModulatorBiasXqWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_ModulatorBiasXqWatcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasXq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXqPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasXq) bool) *oc.Component_OpticalChannel_ModulatorBiasXqWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXqPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXqPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasXq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasXq_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasXq_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasXq_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasXq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasXq_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/avg to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasXq_AvgPath extracts the value of the leaf Avg from its parent oc.Component_OpticalChannel_ModulatorBiasXq
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasXq_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasXq) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasXq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasXq_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasXq_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasXq_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasXq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasXq_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/instant to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasXq_InstantPath extracts the value of the leaf Instant from its parent oc.Component_OpticalChannel_ModulatorBiasXq
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasXq_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasXq) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasXq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasXq_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasXq_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasXq_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasXq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasXq_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/interval to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasXq_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_OpticalChannel_ModulatorBiasXq
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasXq_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasXq) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasXq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasXq_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasXq_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasXq_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasXq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasXq_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasXq_MaxPath extracts the value of the leaf Max from its parent oc.Component_OpticalChannel_ModulatorBiasXq
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasXq_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasXq) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasXq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasXq_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasXq_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasXq_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasXq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasXq_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/max-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasXq_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_OpticalChannel_ModulatorBiasXq
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasXq_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasXq) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasXq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasXq_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasXq_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasXq_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasXq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasXq_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasXq_MinPath extracts the value of the leaf Min from its parent oc.Component_OpticalChannel_ModulatorBiasXq
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasXq_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasXq) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasXq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasXq_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasXq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasXq_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasXq_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasXq_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasXq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasXq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasXq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasXq_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasXq_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-xq/min-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasXq_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasXq_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_OpticalChannel_ModulatorBiasXq
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasXq_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasXq) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePath) Lookup(t testing.TB) *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePath) Get(t testing.TB) *oc.Component_OpticalChannel_ModulatorBiasYPhase {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePathAny) Get(t testing.TB) []*oc.Component_OpticalChannel_ModulatorBiasYPhase {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_OpticalChannel_ModulatorBiasYPhase
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_ModulatorBiasYPhase {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_ModulatorBiasYPhase{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_OpticalChannel_ModulatorBiasYPhase)))
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhasePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase) bool) *oc.Component_OpticalChannel_ModulatorBiasYPhaseWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_ModulatorBiasYPhaseWatcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase) bool) *oc.Component_OpticalChannel_ModulatorBiasYPhaseWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhasePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_OpticalChannel_ModulatorBiasYPhase) *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_ModulatorBiasYPhase {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_ModulatorBiasYPhase{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhasePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase) bool) *oc.Component_OpticalChannel_ModulatorBiasYPhaseWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_ModulatorBiasYPhaseWatcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYPhase) bool) *oc.Component_OpticalChannel_ModulatorBiasYPhaseWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhasePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhasePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYPhase_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYPhase_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/avg to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYPhase_AvgPath extracts the value of the leaf Avg from its parent oc.Component_OpticalChannel_ModulatorBiasYPhase
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYPhase_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYPhase) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYPhase_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYPhase_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/instant to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYPhase_InstantPath extracts the value of the leaf Instant from its parent oc.Component_OpticalChannel_ModulatorBiasYPhase
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYPhase_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYPhase) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYPhase_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYPhase_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/interval to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYPhase_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_OpticalChannel_ModulatorBiasYPhase
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasYPhase_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYPhase) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxPath extracts the value of the leaf Max from its parent oc.Component_OpticalChannel_ModulatorBiasYPhase
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYPhase) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/max-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_OpticalChannel_ModulatorBiasYPhase
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasYPhase_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYPhase) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYPhase_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYPhase_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYPhase_MinPath extracts the value of the leaf Min from its parent oc.Component_OpticalChannel_ModulatorBiasYPhase
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYPhase_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYPhase) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYPhase_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYPhase_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYPhase_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYPhase{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYPhase{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYPhase", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYPhase_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYPhase_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-y-phase/min-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYPhase_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYPhase_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_OpticalChannel_ModulatorBiasYPhase
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasYPhase_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYPhase) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYiPath) Lookup(t testing.TB) *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYi", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYiPath) Get(t testing.TB) *oc.Component_OpticalChannel_ModulatorBiasYi {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYiPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYiPathAny) Get(t testing.TB) []*oc.Component_OpticalChannel_ModulatorBiasYi {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_OpticalChannel_ModulatorBiasYi
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYiPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_ModulatorBiasYi {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_ModulatorBiasYi{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_OpticalChannel_ModulatorBiasYi)))
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYiPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi) bool) *oc.Component_OpticalChannel_ModulatorBiasYiWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_ModulatorBiasYiWatcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYiPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi) bool) *oc.Component_OpticalChannel_ModulatorBiasYiWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYiPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYiPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_OpticalChannel_ModulatorBiasYi) *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYiPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYiPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_ModulatorBiasYi {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_ModulatorBiasYi{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYiPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi) bool) *oc.Component_OpticalChannel_ModulatorBiasYiWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_ModulatorBiasYiWatcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYi{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYiPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYi) bool) *oc.Component_OpticalChannel_ModulatorBiasYiWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYiPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYiPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYi", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYi_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYi_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYi_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYi{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYi_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/avg to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYi_AvgPath extracts the value of the leaf Avg from its parent oc.Component_OpticalChannel_ModulatorBiasYi
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYi_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYi) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYi", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYi_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYi_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYi_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYi{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYi_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/instant to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYi_InstantPath extracts the value of the leaf Instant from its parent oc.Component_OpticalChannel_ModulatorBiasYi
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYi_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYi) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYi", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYi_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYi_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYi_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYi{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYi_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/interval to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYi_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_OpticalChannel_ModulatorBiasYi
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasYi_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYi) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYi", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYi_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYi_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYi_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYi{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYi_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYi_MaxPath extracts the value of the leaf Max from its parent oc.Component_OpticalChannel_ModulatorBiasYi
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYi_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYi) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYi", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYi_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYi_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYi_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYi{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYi_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/max-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYi_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_OpticalChannel_ModulatorBiasYi
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasYi_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYi) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYi", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYi_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYi_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYi_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYi{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYi_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYi_MinPath extracts the value of the leaf Min from its parent oc.Component_OpticalChannel_ModulatorBiasYi
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYi_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYi) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYi", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYi_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYi{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYi_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYi_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYi_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYi{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYi{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYi", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYi_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYi_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yi/min-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYi_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYi_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_OpticalChannel_ModulatorBiasYi
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasYi_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYi) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYqPath) Lookup(t testing.TB) *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYq", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYqPath) Get(t testing.TB) *oc.Component_OpticalChannel_ModulatorBiasYq {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYqPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYqPathAny) Get(t testing.TB) []*oc.Component_OpticalChannel_ModulatorBiasYq {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_OpticalChannel_ModulatorBiasYq
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYqPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_ModulatorBiasYq {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_ModulatorBiasYq{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_OpticalChannel_ModulatorBiasYq)))
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYqPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq) bool) *oc.Component_OpticalChannel_ModulatorBiasYqWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_ModulatorBiasYqWatcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYqPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq) bool) *oc.Component_OpticalChannel_ModulatorBiasYqWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYqPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYqPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_OpticalChannel_ModulatorBiasYq) *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYqPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYqPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_ModulatorBiasYq {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_ModulatorBiasYq{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYqPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq) bool) *oc.Component_OpticalChannel_ModulatorBiasYqWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_ModulatorBiasYqWatcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYqPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ModulatorBiasYq) bool) *oc.Component_OpticalChannel_ModulatorBiasYqWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYqPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYqPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYq_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYq_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYq_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYq_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/avg to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYq_AvgPath extracts the value of the leaf Avg from its parent oc.Component_OpticalChannel_ModulatorBiasYq
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYq_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYq) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYq_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYq_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYq_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYq_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/instant to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYq_InstantPath extracts the value of the leaf Instant from its parent oc.Component_OpticalChannel_ModulatorBiasYq
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYq_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYq) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYq_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYq_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYq_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYq_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/interval to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYq_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_OpticalChannel_ModulatorBiasYq
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasYq_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYq) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYq_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYq_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYq_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYq_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYq_MaxPath extracts the value of the leaf Max from its parent oc.Component_OpticalChannel_ModulatorBiasYq
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYq_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYq) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYq_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYq_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYq_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYq_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/max-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYq_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_OpticalChannel_ModulatorBiasYq
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasYq_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYq) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYq_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYq_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYq_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYq_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYq_MinPath extracts the value of the leaf Min from its parent oc.Component_OpticalChannel_ModulatorBiasYq
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ModulatorBiasYq_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYq) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ModulatorBiasYq", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ModulatorBiasYq_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ModulatorBiasYq{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ModulatorBiasYq_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time with a ONCE subscription.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ModulatorBiasYq_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ModulatorBiasYq_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ModulatorBiasYq{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ModulatorBiasYq{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ModulatorBiasYq", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ModulatorBiasYq_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ModulatorBiasYq_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/modulator-bias-yq/min-time to the batch object.
func (n *Component_OpticalChannel_ModulatorBiasYq_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ModulatorBiasYq_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_OpticalChannel_ModulatorBiasYq
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ModulatorBiasYq_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ModulatorBiasYq) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/operational-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_OperationalModePath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_OperationalModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/operational-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_OperationalModePath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/operational-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_OperationalModePathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_OperationalModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/operational-mode with a ONCE subscription.
func (n *Component_OpticalChannel_OperationalModePathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/operational-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OperationalModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OperationalModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Component_OpticalChannel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_OperationalModePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/operational-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OperationalModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_OperationalModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/operational-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_OperationalModePath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/operational-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/operational-mode to the batch object.
func (n *Component_OpticalChannel_OperationalModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/operational-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OperationalModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OperationalModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.Component_OpticalChannel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_OperationalModePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/operational-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OperationalModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_OperationalModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/operational-mode to the batch object.
func (n *Component_OpticalChannel_OperationalModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_OperationalModePath extracts the value of the leaf OperationalMode from its parent oc.Component_OpticalChannel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertComponent_OpticalChannel_OperationalModePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.OperationalMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_OsnrPath) Lookup(t testing.TB) *oc.QualifiedComponent_OpticalChannel_Osnr {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Osnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Osnr", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_OpticalChannel_Osnr{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_OsnrPath) Get(t testing.TB) *oc.Component_OpticalChannel_Osnr {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_OsnrPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_OpticalChannel_Osnr {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_OpticalChannel_Osnr
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Osnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Osnr", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_OpticalChannel_Osnr{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr with a ONCE subscription.
func (n *Component_OpticalChannel_OsnrPathAny) Get(t testing.TB) []*oc.Component_OpticalChannel_Osnr {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_OpticalChannel_Osnr
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OsnrPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_Osnr {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_Osnr{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_Osnr) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_OpticalChannel_Osnr{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_OpticalChannel_Osnr)))
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OsnrPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_Osnr) bool) *oc.Component_OpticalChannel_OsnrWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_OsnrWatcher{}
	gs := &oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Osnr", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_OpticalChannel_Osnr{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_Osnr)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OsnrPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_Osnr) bool) *oc.Component_OpticalChannel_OsnrWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_OsnrPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/osnr with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_OsnrPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_OpticalChannel_Osnr) *oc.QualifiedComponent_OpticalChannel_Osnr {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_OpticalChannel_Osnr) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/osnr failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr to the batch object.
func (n *Component_OpticalChannel_OsnrPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OsnrPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_Osnr {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_Osnr{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_Osnr) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OsnrPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_Osnr) bool) *oc.Component_OpticalChannel_OsnrWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_OsnrWatcher{}
	structs := map[string]*oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Osnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Osnr", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_OpticalChannel_Osnr{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_Osnr)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OsnrPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_Osnr) bool) *oc.Component_OpticalChannel_OsnrWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_OsnrPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr to the batch object.
func (n *Component_OpticalChannel_OsnrPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_Osnr_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Osnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Osnr", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_Osnr_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_Osnr_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_Osnr_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Osnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Osnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_Osnr_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/avg with a ONCE subscription.
func (n *Component_OpticalChannel_Osnr_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Osnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_Osnr_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/osnr/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_Osnr_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/osnr/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/avg to the batch object.
func (n *Component_OpticalChannel_Osnr_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Osnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Osnr", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_Osnr_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/avg to the batch object.
func (n *Component_OpticalChannel_Osnr_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_Osnr_AvgPath extracts the value of the leaf Avg from its parent oc.Component_OpticalChannel_Osnr
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_Osnr_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_Osnr) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_Osnr_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Osnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Osnr", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_Osnr_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_Osnr_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_Osnr_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Osnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Osnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_Osnr_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/instant with a ONCE subscription.
func (n *Component_OpticalChannel_Osnr_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Osnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_Osnr_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/osnr/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_Osnr_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/osnr/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/instant to the batch object.
func (n *Component_OpticalChannel_Osnr_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Osnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Osnr", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_Osnr_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/instant to the batch object.
func (n *Component_OpticalChannel_Osnr_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_Osnr_InstantPath extracts the value of the leaf Instant from its parent oc.Component_OpticalChannel_Osnr
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_Osnr_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_Osnr) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_Osnr_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Osnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Osnr", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_Osnr_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_Osnr_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_Osnr_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Osnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Osnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_Osnr_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/interval with a ONCE subscription.
func (n *Component_OpticalChannel_Osnr_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Osnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_Osnr_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/osnr/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_Osnr_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/osnr/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/interval to the batch object.
func (n *Component_OpticalChannel_Osnr_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Osnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Osnr", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_Osnr_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/interval to the batch object.
func (n *Component_OpticalChannel_Osnr_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_Osnr_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_OpticalChannel_Osnr
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_Osnr_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_Osnr) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_Osnr_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Osnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Osnr", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_Osnr_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_Osnr_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_Osnr_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Osnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Osnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_Osnr_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/max with a ONCE subscription.
func (n *Component_OpticalChannel_Osnr_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Osnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_Osnr_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/osnr/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_Osnr_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/osnr/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/max to the batch object.
func (n *Component_OpticalChannel_Osnr_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Osnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Osnr", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_Osnr_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/max to the batch object.
func (n *Component_OpticalChannel_Osnr_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_Osnr_MaxPath extracts the value of the leaf Max from its parent oc.Component_OpticalChannel_Osnr
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_Osnr_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_Osnr) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_Osnr_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Osnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Osnr", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_Osnr_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_Osnr_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_Osnr_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Osnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Osnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_Osnr_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/max-time with a ONCE subscription.
func (n *Component_OpticalChannel_Osnr_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Osnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_Osnr_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/osnr/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_Osnr_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/osnr/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/max-time to the batch object.
func (n *Component_OpticalChannel_Osnr_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Osnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Osnr", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_Osnr_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/max-time to the batch object.
func (n *Component_OpticalChannel_Osnr_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_Osnr_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_OpticalChannel_Osnr
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_Osnr_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_Osnr) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_Osnr_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Osnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Osnr", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_Osnr_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_Osnr_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_Osnr_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Osnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Osnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_Osnr_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/min with a ONCE subscription.
func (n *Component_OpticalChannel_Osnr_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Osnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_Osnr_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/osnr/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_Osnr_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/osnr/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/min to the batch object.
func (n *Component_OpticalChannel_Osnr_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Osnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Osnr", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_Osnr_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/min to the batch object.
func (n *Component_OpticalChannel_Osnr_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_Osnr_MinPath extracts the value of the leaf Min from its parent oc.Component_OpticalChannel_Osnr
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_Osnr_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_Osnr) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_Osnr_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Osnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Osnr", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_Osnr_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/osnr/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_Osnr_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_Osnr_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Osnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Osnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_Osnr_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/osnr/min-time with a ONCE subscription.
func (n *Component_OpticalChannel_Osnr_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Osnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_Osnr_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/osnr/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_Osnr_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/osnr/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/min-time to the batch object.
func (n *Component_OpticalChannel_Osnr_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/osnr/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Osnr_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Osnr_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_Osnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Osnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Osnr", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_Osnr_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/osnr/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Osnr_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Osnr_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/osnr/min-time to the batch object.
func (n *Component_OpticalChannel_Osnr_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_Osnr_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_OpticalChannel_Osnr
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_Osnr_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_Osnr) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/output-power with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_OutputPowerPath) Lookup(t testing.TB) *oc.QualifiedComponent_OpticalChannel_OutputPower {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_OutputPower", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_OpticalChannel_OutputPower{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/output-power with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_OutputPowerPath) Get(t testing.TB) *oc.Component_OpticalChannel_OutputPower {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/output-power with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_OutputPowerPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_OpticalChannel_OutputPower {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_OpticalChannel_OutputPower
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_OutputPower", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_OpticalChannel_OutputPower{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/output-power with a ONCE subscription.
func (n *Component_OpticalChannel_OutputPowerPathAny) Get(t testing.TB) []*oc.Component_OpticalChannel_OutputPower {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_OpticalChannel_OutputPower
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/output-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OutputPowerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_OutputPower {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_OutputPower{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_OutputPower) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_OpticalChannel_OutputPower{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_OpticalChannel_OutputPower)))
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OutputPowerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_OutputPower) bool) *oc.Component_OpticalChannel_OutputPowerWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_OutputPowerWatcher{}
	gs := &oc.Component_OpticalChannel_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_OutputPower", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_OpticalChannel_OutputPower{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_OutputPower)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/output-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OutputPowerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_OutputPower) bool) *oc.Component_OpticalChannel_OutputPowerWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_OutputPowerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/output-power with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_OutputPowerPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_OpticalChannel_OutputPower) *oc.QualifiedComponent_OpticalChannel_OutputPower {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_OpticalChannel_OutputPower) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/output-power failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/output-power to the batch object.
func (n *Component_OpticalChannel_OutputPowerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/output-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OutputPowerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_OutputPower {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_OutputPower{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_OutputPower) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OutputPowerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_OutputPower) bool) *oc.Component_OpticalChannel_OutputPowerWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_OutputPowerWatcher{}
	structs := map[string]*oc.Component_OpticalChannel_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_OutputPower", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_OpticalChannel_OutputPower{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_OutputPower)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/output-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OutputPowerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_OutputPower) bool) *oc.Component_OpticalChannel_OutputPowerWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_OutputPowerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/output-power to the batch object.
func (n *Component_OpticalChannel_OutputPowerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/output-power/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_OutputPower_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_OutputPower_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/output-power/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_OutputPower_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/output-power/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_OutputPower_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_OutputPower_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/output-power/avg with a ONCE subscription.
func (n *Component_OpticalChannel_OutputPower_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/output-power/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OutputPower_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OutputPower_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_OutputPower_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/output-power/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OutputPower_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_OutputPower_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/output-power/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_OutputPower_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/output-power/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/output-power/avg to the batch object.
func (n *Component_OpticalChannel_OutputPower_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/output-power/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OutputPower_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OutputPower_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_OutputPower_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/output-power/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OutputPower_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_OutputPower_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/output-power/avg to the batch object.
func (n *Component_OpticalChannel_OutputPower_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_OutputPower_AvgPath extracts the value of the leaf Avg from its parent oc.Component_OpticalChannel_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_OutputPower_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_OutputPower) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/output-power/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_OutputPower_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_OutputPower_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/output-power/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_OutputPower_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/output-power/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_OutputPower_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_OutputPower_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/output-power/instant with a ONCE subscription.
func (n *Component_OpticalChannel_OutputPower_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/output-power/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OutputPower_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OutputPower_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_OutputPower_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/output-power/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OutputPower_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_OutputPower_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/output-power/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_OutputPower_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/output-power/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/output-power/instant to the batch object.
func (n *Component_OpticalChannel_OutputPower_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/output-power/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OutputPower_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OutputPower_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_OutputPower_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/output-power/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OutputPower_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_OutputPower_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/output-power/instant to the batch object.
func (n *Component_OpticalChannel_OutputPower_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_OutputPower_InstantPath extracts the value of the leaf Instant from its parent oc.Component_OpticalChannel_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_OutputPower_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_OutputPower) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/output-power/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_OutputPower_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_OutputPower_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/output-power/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_OutputPower_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/output-power/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_OutputPower_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_OutputPower_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/output-power/interval with a ONCE subscription.
func (n *Component_OpticalChannel_OutputPower_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/output-power/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OutputPower_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OutputPower_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_OutputPower_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/output-power/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OutputPower_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_OutputPower_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/output-power/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_OutputPower_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/output-power/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/output-power/interval to the batch object.
func (n *Component_OpticalChannel_OutputPower_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/output-power/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OutputPower_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OutputPower_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_OutputPower_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/output-power/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OutputPower_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_OutputPower_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/output-power/interval to the batch object.
func (n *Component_OpticalChannel_OutputPower_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_OutputPower_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_OpticalChannel_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_OutputPower_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_OutputPower) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/output-power/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_OutputPower_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_OutputPower_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/output-power/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_OutputPower_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/output-power/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_OutputPower_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_OutputPower_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/output-power/max with a ONCE subscription.
func (n *Component_OpticalChannel_OutputPower_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/output-power/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OutputPower_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OutputPower_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_OutputPower_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/output-power/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OutputPower_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_OutputPower_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/output-power/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_OutputPower_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/output-power/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/output-power/max to the batch object.
func (n *Component_OpticalChannel_OutputPower_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/output-power/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_OutputPower_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_OutputPower_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_OutputPower_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/output-power/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_OutputPower_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_OutputPower_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/output-power/max to the batch object.
func (n *Component_OpticalChannel_OutputPower_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_OutputPower_MaxPath extracts the value of the leaf Max from its parent oc.Component_OpticalChannel_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_OutputPower_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_OutputPower) *oc.QualifiedFloat64 {
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
