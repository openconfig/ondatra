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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_InputPower_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_InputPower_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_InputPower_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_InputPower_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_InputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_InputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_InputPower_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/min with a ONCE subscription.
func (n *Component_Transceiver_InputPower_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_InputPower_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_InputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_InputPower_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/input-power/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_InputPower_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/input-power/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/min to the batch object.
func (n *Component_Transceiver_InputPower_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_InputPower_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_InputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_InputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_InputPower_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/min to the batch object.
func (n *Component_Transceiver_InputPower_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_InputPower_MinPath extracts the value of the leaf Min from its parent oc.Component_Transceiver_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_InputPower_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_InputPower) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_InputPower_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_InputPower_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_InputPower_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_InputPower_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_InputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_InputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_InputPower_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/min-time with a ONCE subscription.
func (n *Component_Transceiver_InputPower_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_InputPower_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_InputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_InputPower_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/input-power/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_InputPower_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/input-power/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/min-time to the batch object.
func (n *Component_Transceiver_InputPower_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_InputPower_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_InputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_InputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_InputPower_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/min-time to the batch object.
func (n *Component_Transceiver_InputPower_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_InputPower_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_Transceiver_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_InputPower_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_InputPower) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_LaserBiasCurrentPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver_LaserBiasCurrent {
	t.Helper()
	goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_LaserBiasCurrent", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Transceiver_LaserBiasCurrent{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_LaserBiasCurrentPath) Get(t testing.TB) *oc.Component_Transceiver_LaserBiasCurrent {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_LaserBiasCurrentPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver_LaserBiasCurrent {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver_LaserBiasCurrent
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver_LaserBiasCurrent{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current with a ONCE subscription.
func (n *Component_Transceiver_LaserBiasCurrentPathAny) Get(t testing.TB) []*oc.Component_Transceiver_LaserBiasCurrent {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver_LaserBiasCurrent
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_LaserBiasCurrent {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_LaserBiasCurrent{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_LaserBiasCurrent) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Transceiver_LaserBiasCurrent{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Transceiver_LaserBiasCurrent)))
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_LaserBiasCurrent) bool) *oc.Component_Transceiver_LaserBiasCurrentWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_LaserBiasCurrentWatcher{}
	gs := &oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_Transceiver_LaserBiasCurrent{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_LaserBiasCurrent)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_LaserBiasCurrent) bool) *oc.Component_Transceiver_LaserBiasCurrentWatcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/laser-bias-current with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_LaserBiasCurrentPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Transceiver_LaserBiasCurrent) *oc.QualifiedComponent_Transceiver_LaserBiasCurrent {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Transceiver_LaserBiasCurrent) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/laser-bias-current failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current to the batch object.
func (n *Component_Transceiver_LaserBiasCurrentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_LaserBiasCurrent {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_LaserBiasCurrent{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_LaserBiasCurrent) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrentPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_LaserBiasCurrent) bool) *oc.Component_Transceiver_LaserBiasCurrentWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_LaserBiasCurrentWatcher{}
	structs := map[string]*oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_LaserBiasCurrent{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_Transceiver_LaserBiasCurrent{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_LaserBiasCurrent)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_LaserBiasCurrent) bool) *oc.Component_Transceiver_LaserBiasCurrentWatcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrentPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current to the batch object.
func (n *Component_Transceiver_LaserBiasCurrentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_LaserBiasCurrent", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_LaserBiasCurrent_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_LaserBiasCurrent_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg with a ONCE subscription.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_LaserBiasCurrent_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_LaserBiasCurrent{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_LaserBiasCurrent_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_LaserBiasCurrent_AvgPath extracts the value of the leaf Avg from its parent oc.Component_Transceiver_LaserBiasCurrent
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_LaserBiasCurrent_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_LaserBiasCurrent) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_LaserBiasCurrent", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_LaserBiasCurrent_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_LaserBiasCurrent_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant with a ONCE subscription.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_LaserBiasCurrent_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_LaserBiasCurrent{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_LaserBiasCurrent_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_LaserBiasCurrent_InstantPath extracts the value of the leaf Instant from its parent oc.Component_Transceiver_LaserBiasCurrent
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_LaserBiasCurrent_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_LaserBiasCurrent) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_LaserBiasCurrent", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_LaserBiasCurrent_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_LaserBiasCurrent_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval with a ONCE subscription.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_LaserBiasCurrent_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_LaserBiasCurrent{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_LaserBiasCurrent_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_LaserBiasCurrent_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_Transceiver_LaserBiasCurrent
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_LaserBiasCurrent_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_LaserBiasCurrent) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_LaserBiasCurrent", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_LaserBiasCurrent_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_LaserBiasCurrent_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max with a ONCE subscription.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_LaserBiasCurrent_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/max to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_LaserBiasCurrent{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_LaserBiasCurrent_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/max to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_LaserBiasCurrent_MaxPath extracts the value of the leaf Max from its parent oc.Component_Transceiver_LaserBiasCurrent
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_LaserBiasCurrent_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_LaserBiasCurrent) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_LaserBiasCurrent", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_LaserBiasCurrent_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_LaserBiasCurrent_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time with a ONCE subscription.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_LaserBiasCurrent_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_LaserBiasCurrent{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_LaserBiasCurrent_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_LaserBiasCurrent_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_Transceiver_LaserBiasCurrent
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_LaserBiasCurrent_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_LaserBiasCurrent) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_LaserBiasCurrent", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_LaserBiasCurrent_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_LaserBiasCurrent_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_LaserBiasCurrent_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min with a ONCE subscription.
func (n *Component_Transceiver_LaserBiasCurrent_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_LaserBiasCurrent_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_LaserBiasCurrent_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/min to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_LaserBiasCurrent{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_LaserBiasCurrent_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/min to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_LaserBiasCurrent_MinPath extracts the value of the leaf Min from its parent oc.Component_Transceiver_LaserBiasCurrent
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_LaserBiasCurrent_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_LaserBiasCurrent) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_LaserBiasCurrent", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_LaserBiasCurrent_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_LaserBiasCurrent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_LaserBiasCurrent_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time with a ONCE subscription.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_LaserBiasCurrent_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_LaserBiasCurrent_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_LaserBiasCurrent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_LaserBiasCurrent{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_LaserBiasCurrent_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/laser-bias-current/min-time to the batch object.
func (n *Component_Transceiver_LaserBiasCurrent_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_LaserBiasCurrent_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_Transceiver_LaserBiasCurrent
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_LaserBiasCurrent_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_LaserBiasCurrent) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/module-functional-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_ModuleFunctionalTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/module-functional-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Get(t testing.TB) oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/module-functional-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_ModuleFunctionalTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_ModuleFunctionalTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/module-functional-type with a ONCE subscription.
func (n *Component_Transceiver_ModuleFunctionalTypePathAny) Get(t testing.TB) []oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/module-functional-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_ModuleFunctionalTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) bool) *oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPEWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_ModuleFunctionalTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/module-functional-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) bool) *oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_ModuleFunctionalTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/module-functional-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/module-functional-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/module-functional-type to the batch object.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/module-functional-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_ModuleFunctionalTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_ModuleFunctionalTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) bool) *oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPEWatcher{}
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
			qv := convertComponent_Transceiver_ModuleFunctionalTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/module-functional-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_ModuleFunctionalTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) bool) *oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_ModuleFunctionalTypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/module-functional-type to the batch object.
func (n *Component_Transceiver_ModuleFunctionalTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_ModuleFunctionalTypePath extracts the value of the leaf ModuleFunctionalType from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE.
func convertComponent_Transceiver_ModuleFunctionalTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE{
		Metadata: md,
	}
	val := parent.ModuleFunctionalType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/otn-compliance-code with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_OtnComplianceCodePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_OtnComplianceCodePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/otn-compliance-code with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_OtnComplianceCodePath) Get(t testing.TB) oc.E_TransportTypes_OTN_APPLICATION_CODE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/otn-compliance-code with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_OtnComplianceCodePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_OtnComplianceCodePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/otn-compliance-code with a ONCE subscription.
func (n *Component_Transceiver_OtnComplianceCodePathAny) Get(t testing.TB) []oc.E_TransportTypes_OTN_APPLICATION_CODE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_OTN_APPLICATION_CODE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/otn-compliance-code with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OtnComplianceCodePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_OTN_APPLICATION_CODE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_OTN_APPLICATION_CODE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OtnComplianceCodePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE) bool) *oc.E_TransportTypes_OTN_APPLICATION_CODEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_OTN_APPLICATION_CODEWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_OtnComplianceCodePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/otn-compliance-code with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OtnComplianceCodePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE) bool) *oc.E_TransportTypes_OTN_APPLICATION_CODEWatcher {
	t.Helper()
	return watch_Component_Transceiver_OtnComplianceCodePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/otn-compliance-code with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_OtnComplianceCodePath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_OTN_APPLICATION_CODE) *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/otn-compliance-code failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/otn-compliance-code to the batch object.
func (n *Component_Transceiver_OtnComplianceCodePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/otn-compliance-code with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OtnComplianceCodePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_OTN_APPLICATION_CODE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_OTN_APPLICATION_CODE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OtnComplianceCodePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE) bool) *oc.E_TransportTypes_OTN_APPLICATION_CODEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_OTN_APPLICATION_CODEWatcher{}
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
			qv := convertComponent_Transceiver_OtnComplianceCodePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/otn-compliance-code with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OtnComplianceCodePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE) bool) *oc.E_TransportTypes_OTN_APPLICATION_CODEWatcher {
	t.Helper()
	return watch_Component_Transceiver_OtnComplianceCodePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/otn-compliance-code to the batch object.
func (n *Component_Transceiver_OtnComplianceCodePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_OtnComplianceCodePath extracts the value of the leaf OtnComplianceCode from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE.
func convertComponent_Transceiver_OtnComplianceCodePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_OTN_APPLICATION_CODE{
		Metadata: md,
	}
	val := parent.OtnComplianceCode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/output-power with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_OutputPowerPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver_OutputPower {
	t.Helper()
	goStruct := &oc.Component_Transceiver_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_OutputPower", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Transceiver_OutputPower{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/output-power with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_OutputPowerPath) Get(t testing.TB) *oc.Component_Transceiver_OutputPower {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/output-power with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_OutputPowerPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver_OutputPower {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver_OutputPower
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_OutputPower", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver_OutputPower{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/output-power with a ONCE subscription.
func (n *Component_Transceiver_OutputPowerPathAny) Get(t testing.TB) []*oc.Component_Transceiver_OutputPower {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver_OutputPower
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPowerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_OutputPower {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_OutputPower{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_OutputPower) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Transceiver_OutputPower{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Transceiver_OutputPower)))
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPowerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_OutputPower) bool) *oc.Component_Transceiver_OutputPowerWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_OutputPowerWatcher{}
	gs := &oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_OutputPower", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_Transceiver_OutputPower{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_OutputPower)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPowerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_OutputPower) bool) *oc.Component_Transceiver_OutputPowerWatcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPowerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/output-power with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_OutputPowerPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Transceiver_OutputPower) *oc.QualifiedComponent_Transceiver_OutputPower {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Transceiver_OutputPower) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/output-power failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power to the batch object.
func (n *Component_Transceiver_OutputPowerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPowerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_OutputPower {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_OutputPower{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_OutputPower) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPowerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_OutputPower) bool) *oc.Component_Transceiver_OutputPowerWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_OutputPowerWatcher{}
	structs := map[string]*oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_OutputPower", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_Transceiver_OutputPower{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_OutputPower)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPowerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_OutputPower) bool) *oc.Component_Transceiver_OutputPowerWatcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPowerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power to the batch object.
func (n *Component_Transceiver_OutputPowerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_OutputPower_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_OutputPower_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_OutputPower_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_OutputPower_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_OutputPower_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/avg with a ONCE subscription.
func (n *Component_Transceiver_OutputPower_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_OutputPower_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/output-power/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_OutputPower_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/output-power/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/avg to the batch object.
func (n *Component_Transceiver_OutputPower_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_OutputPower_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/avg to the batch object.
func (n *Component_Transceiver_OutputPower_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_OutputPower_AvgPath extracts the value of the leaf Avg from its parent oc.Component_Transceiver_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_OutputPower_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_OutputPower) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_OutputPower_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_OutputPower_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_OutputPower_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_OutputPower_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_OutputPower_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/instant with a ONCE subscription.
func (n *Component_Transceiver_OutputPower_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_OutputPower_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/output-power/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_OutputPower_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/output-power/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/instant to the batch object.
func (n *Component_Transceiver_OutputPower_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_OutputPower_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/instant to the batch object.
func (n *Component_Transceiver_OutputPower_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_OutputPower_InstantPath extracts the value of the leaf Instant from its parent oc.Component_Transceiver_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_OutputPower_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_OutputPower) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_OutputPower_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_OutputPower_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_OutputPower_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_OutputPower_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_OutputPower_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/interval with a ONCE subscription.
func (n *Component_Transceiver_OutputPower_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_OutputPower_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/output-power/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_OutputPower_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/output-power/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/interval to the batch object.
func (n *Component_Transceiver_OutputPower_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_OutputPower_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/interval to the batch object.
func (n *Component_Transceiver_OutputPower_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_OutputPower_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_Transceiver_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_OutputPower_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_OutputPower) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_OutputPower_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_OutputPower_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_OutputPower_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_OutputPower_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_OutputPower_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/max with a ONCE subscription.
func (n *Component_Transceiver_OutputPower_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_OutputPower_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/output-power/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_OutputPower_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/output-power/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/max to the batch object.
func (n *Component_Transceiver_OutputPower_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_OutputPower_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/max to the batch object.
func (n *Component_Transceiver_OutputPower_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_OutputPower_MaxPath extracts the value of the leaf Max from its parent oc.Component_Transceiver_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_OutputPower_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_OutputPower) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_OutputPower_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_OutputPower_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_OutputPower_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_OutputPower_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_OutputPower_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/max-time with a ONCE subscription.
func (n *Component_Transceiver_OutputPower_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_OutputPower_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/output-power/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_OutputPower_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/output-power/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/max-time to the batch object.
func (n *Component_Transceiver_OutputPower_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_OutputPower_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/max-time to the batch object.
func (n *Component_Transceiver_OutputPower_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_OutputPower_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_Transceiver_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_OutputPower_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_OutputPower) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_OutputPower_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_OutputPower_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_OutputPower_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_OutputPower_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_OutputPower_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/min with a ONCE subscription.
func (n *Component_Transceiver_OutputPower_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_OutputPower_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/output-power/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_OutputPower_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/output-power/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/min to the batch object.
func (n *Component_Transceiver_OutputPower_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_OutputPower_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/min to the batch object.
func (n *Component_Transceiver_OutputPower_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_OutputPower_MinPath extracts the value of the leaf Min from its parent oc.Component_Transceiver_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_OutputPower_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_OutputPower) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_OutputPower_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_OutputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_OutputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_OutputPower_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/output-power/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_OutputPower_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_OutputPower_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_OutputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_OutputPower", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_OutputPower_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/output-power/min-time with a ONCE subscription.
func (n *Component_Transceiver_OutputPower_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_OutputPower", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_OutputPower_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/output-power/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_OutputPower_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/output-power/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/min-time to the batch object.
func (n *Component_Transceiver_OutputPower_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/output-power/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_OutputPower_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_OutputPower_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Transceiver_OutputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Transceiver_OutputPower{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_OutputPower", structs[pre], queryPath, true, false)
			qv := convertComponent_Transceiver_OutputPower_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/output-power/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_OutputPower_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_OutputPower_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/output-power/min-time to the batch object.
func (n *Component_Transceiver_OutputPower_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_OutputPower_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_Transceiver_OutputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_OutputPower_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_OutputPower) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PostFecBerPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver_PostFecBer {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PostFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PostFecBer", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Transceiver_PostFecBer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PostFecBerPath) Get(t testing.TB) *oc.Component_Transceiver_PostFecBer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PostFecBerPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver_PostFecBer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver_PostFecBer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_PostFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_PostFecBer", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver_PostFecBer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber with a ONCE subscription.
func (n *Component_Transceiver_PostFecBerPathAny) Get(t testing.TB) []*oc.Component_Transceiver_PostFecBer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver_PostFecBer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_PostFecBer {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_PostFecBer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_PostFecBer) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Transceiver_PostFecBer{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Transceiver_PostFecBer)))
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_PostFecBer) bool) *oc.Component_Transceiver_PostFecBerWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_PostFecBerWatcher{}
	gs := &oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PostFecBer", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_Transceiver_PostFecBer{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_PostFecBer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_PostFecBer) bool) *oc.Component_Transceiver_PostFecBerWatcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/post-fec-ber with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PostFecBerPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Transceiver_PostFecBer) *oc.QualifiedComponent_Transceiver_PostFecBer {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Transceiver_PostFecBer) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/post-fec-ber failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber to the batch object.
func (n *Component_Transceiver_PostFecBerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_PostFecBer {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_PostFecBer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_PostFecBer) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_PostFecBer) bool) *oc.Component_Transceiver_PostFecBerWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_PostFecBerWatcher{}
	structs := map[string]*oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Transceiver_PostFecBer", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_Transceiver_PostFecBer{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_PostFecBer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_PostFecBer) bool) *oc.Component_Transceiver_PostFecBerWatcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber to the batch object.
func (n *Component_Transceiver_PostFecBerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PostFecBer_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PostFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PostFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PostFecBer_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PostFecBer_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PostFecBer_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
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
		qv := convertComponent_Transceiver_PostFecBer_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg with a ONCE subscription.
func (n *Component_Transceiver_PostFecBer_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PostFecBer_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PostFecBer_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg to the batch object.
func (n *Component_Transceiver_PostFecBer_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
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
			qv := convertComponent_Transceiver_PostFecBer_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/avg to the batch object.
func (n *Component_Transceiver_PostFecBer_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PostFecBer_AvgPath extracts the value of the leaf Avg from its parent oc.Component_Transceiver_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_PostFecBer_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PostFecBer) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PostFecBer_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PostFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PostFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PostFecBer_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PostFecBer_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PostFecBer_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
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
		qv := convertComponent_Transceiver_PostFecBer_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant with a ONCE subscription.
func (n *Component_Transceiver_PostFecBer_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PostFecBer_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PostFecBer_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant to the batch object.
func (n *Component_Transceiver_PostFecBer_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
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
			qv := convertComponent_Transceiver_PostFecBer_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/instant to the batch object.
func (n *Component_Transceiver_PostFecBer_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PostFecBer_InstantPath extracts the value of the leaf Instant from its parent oc.Component_Transceiver_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_PostFecBer_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PostFecBer) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PostFecBer_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PostFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PostFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PostFecBer_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PostFecBer_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PostFecBer_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Transceiver_PostFecBer_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval with a ONCE subscription.
func (n *Component_Transceiver_PostFecBer_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PostFecBer_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PostFecBer_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval to the batch object.
func (n *Component_Transceiver_PostFecBer_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertComponent_Transceiver_PostFecBer_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/interval to the batch object.
func (n *Component_Transceiver_PostFecBer_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PostFecBer_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_Transceiver_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_PostFecBer_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PostFecBer) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_PostFecBer_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_PostFecBer{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_PostFecBer", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_PostFecBer_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_PostFecBer_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_PostFecBer_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
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
		qv := convertComponent_Transceiver_PostFecBer_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max with a ONCE subscription.
func (n *Component_Transceiver_PostFecBer_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Transceiver_PostFecBer_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_PostFecBer_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/max to the batch object.
func (n *Component_Transceiver_PostFecBer_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_PostFecBer_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_PostFecBer_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
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
			qv := convertComponent_Transceiver_PostFecBer_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/post-fec-ber/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_PostFecBer_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_PostFecBer_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/post-fec-ber/max to the batch object.
func (n *Component_Transceiver_PostFecBer_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_PostFecBer_MaxPath extracts the value of the leaf Max from its parent oc.Component_Transceiver_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_PostFecBer_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_PostFecBer) *oc.QualifiedFloat64 {
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
