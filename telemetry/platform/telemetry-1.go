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

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Memory{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Memory", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Memory", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a ONCE subscription.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Memory", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors to the batch object.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Memory{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Memory", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/memory/state/uncorrected-parity-errors to the batch object.
func (n *Component_IntegratedCircuit_Memory_UncorrectedParityErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath extracts the value of the leaf UncorrectedParityErrors from its parent oc.Component_IntegratedCircuit_Memory
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Memory_UncorrectedParityErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Memory) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.UncorrectedParityErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/state/node-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_NodeIdPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_NodeIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/state/node-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_NodeIdPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_NodeIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a ONCE subscription.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_NodeIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_NodeIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_NodeIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_NodeIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_NodeIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_NodeIdPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/state/node-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/state/node-id to the batch object.
func (n *Component_IntegratedCircuit_NodeIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_NodeIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_NodeIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/state/node-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_NodeIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/state/node-id to the batch object.
func (n *Component_IntegratedCircuit_NodeIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_NodeIdPath extracts the value of the leaf NodeId from its parent oc.Component_IntegratedCircuit
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_NodeIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.NodeId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_UtilizationPath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit_Utilization {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_UtilizationPath) Get(t testing.TB) *oc.Component_IntegratedCircuit_Utilization {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit_Utilization {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit_Utilization
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization with a ONCE subscription.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit_Utilization {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit_Utilization
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_UtilizationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_Utilization {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_Utilization{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_IntegratedCircuit_Utilization)))
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_UtilizationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool) *oc.Component_IntegratedCircuit_UtilizationWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_UtilizationWatcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_Utilization)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_UtilizationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool) *oc.Component_IntegratedCircuit_UtilizationWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_UtilizationPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_UtilizationPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_IntegratedCircuit_Utilization) *oc.QualifiedComponent_IntegratedCircuit_Utilization {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization to the batch object.
func (n *Component_IntegratedCircuit_UtilizationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_Utilization {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_Utilization{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_UtilizationPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool) *oc.Component_IntegratedCircuit_UtilizationWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_UtilizationWatcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_Utilization)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization) bool) *oc.Component_IntegratedCircuit_UtilizationWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_UtilizationPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization to the batch object.
func (n *Component_IntegratedCircuit_UtilizationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Lookup(t testing.TB) *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Get(t testing.TB) *oc.Component_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Get(t testing.TB) []*oc.Component_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_IntegratedCircuit_Utilization_Resource
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_Utilization_Resource{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_IntegratedCircuit_Utilization_Resource)))
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_ResourcePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool) *oc.Component_IntegratedCircuit_Utilization_ResourceWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_Utilization_ResourceWatcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool) *oc.Component_IntegratedCircuit_Utilization_ResourceWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_ResourcePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource to the batch object.
func (n *Component_IntegratedCircuit_Utilization_ResourcePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_IntegratedCircuit_Utilization_Resource {
	t.Helper()
	c := &oc.CollectionComponent_IntegratedCircuit_Utilization_Resource{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_ResourcePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool) *oc.Component_IntegratedCircuit_Utilization_ResourceWatcher {
	t.Helper()
	w := &oc.Component_IntegratedCircuit_Utilization_ResourceWatcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_IntegratedCircuit_Utilization_Resource) bool) *oc.Component_IntegratedCircuit_Utilization_ResourceWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_ResourcePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource to the batch object.
func (n *Component_IntegratedCircuit_Utilization_ResourcePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_CommittedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_CommittedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/committed to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_CommittedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath extracts the value of the leaf Committed from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_CommittedPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Committed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_FreePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_FreePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_FreePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_FreePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_FreePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_FreePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_FreePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_FreePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/free to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_FreePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_FreePath extracts the value of the leaf Free from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_FreePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Free
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/high-watermark to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_HighWatermarkPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath extracts the value of the leaf HighWatermark from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_HighWatermarkPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.HighWatermark
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/last-high-watermark to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath extracts the value of the leaf LastHighWatermark from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_LastHighWatermarkPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LastHighWatermark
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/max-limit to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_MaxLimitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath extracts the value of the leaf MaxLimit from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_MaxLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/name to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_NamePath extracts the value of the leaf Name from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_IntegratedCircuit_Utilization_Resource_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Name
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	md, ok := oc.Lookup(t, n, "Component_IntegratedCircuit_Utilization_Resource", goStruct, true, false)
	if ok {
		return convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_IntegratedCircuit_Utilization_Resource{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a ONCE subscription.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_UsedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_UsedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_IntegratedCircuit_Utilization_Resource_UsedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_IntegratedCircuit_Utilization_Resource{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_IntegratedCircuit_Utilization_Resource{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_IntegratedCircuit_Utilization_Resource", structs[pre], queryPath, true, false)
			qv := convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_IntegratedCircuit_Utilization_Resource_UsedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/integrated-circuit/utilization/resources/resource/state/used to the batch object.
func (n *Component_IntegratedCircuit_Utilization_Resource_UsedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath extracts the value of the leaf Used from its parent oc.Component_IntegratedCircuit_Utilization_Resource
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_IntegratedCircuit_Utilization_Resource_UsedPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_IntegratedCircuit_Utilization_Resource) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Used
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/last-reboot-reason with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LastRebootReasonPath) Lookup(t testing.TB) *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_LastRebootReasonPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/last-reboot-reason with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LastRebootReasonPath) Get(t testing.TB) oc.E_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/last-reboot-reason with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LastRebootReasonPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_LastRebootReasonPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/last-reboot-reason with a ONCE subscription.
func (n *Component_LastRebootReasonPathAny) Get(t testing.TB) []oc.E_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PlatformTypes_COMPONENT_REBOOT_REASON
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-reboot-reason with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastRebootReasonPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastRebootReasonPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool) *oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_LastRebootReasonPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-reboot-reason with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastRebootReasonPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool) *oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher {
	t.Helper()
	return watch_Component_LastRebootReasonPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/last-reboot-reason with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LastRebootReasonPath) Await(t testing.TB, timeout time.Duration, val oc.E_PlatformTypes_COMPONENT_REBOOT_REASON) *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/last-reboot-reason failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/last-reboot-reason to the batch object.
func (n *Component_LastRebootReasonPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-reboot-reason with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastRebootReasonPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastRebootReasonPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool) *oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher{}
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
			qv := convertComponent_LastRebootReasonPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-reboot-reason with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastRebootReasonPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) bool) *oc.E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher {
	t.Helper()
	return watch_Component_LastRebootReasonPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/last-reboot-reason to the batch object.
func (n *Component_LastRebootReasonPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_LastRebootReasonPath extracts the value of the leaf LastRebootReason from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON.
func convertComponent_LastRebootReasonPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	qv := &oc.QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON{
		Metadata: md,
	}
	val := parent.LastRebootReason
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/last-reboot-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LastRebootTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_LastRebootTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/last-reboot-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LastRebootTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/last-reboot-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LastRebootTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_LastRebootTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/last-reboot-time with a ONCE subscription.
func (n *Component_LastRebootTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-reboot-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastRebootTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastRebootTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_LastRebootTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-reboot-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastRebootTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_LastRebootTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/last-reboot-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LastRebootTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/last-reboot-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/last-reboot-time to the batch object.
func (n *Component_LastRebootTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-reboot-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastRebootTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastRebootTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertComponent_LastRebootTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-reboot-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastRebootTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_LastRebootTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/last-reboot-time to the batch object.
func (n *Component_LastRebootTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_LastRebootTimePath extracts the value of the leaf LastRebootTime from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_LastRebootTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LastRebootTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/last-switchover-reason with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LastSwitchoverReasonPath) Lookup(t testing.TB) *oc.QualifiedComponent_LastSwitchoverReason {
	t.Helper()
	goStruct := &oc.Component_LastSwitchoverReason{}
	md, ok := oc.Lookup(t, n, "Component_LastSwitchoverReason", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_LastSwitchoverReason{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/last-switchover-reason with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LastSwitchoverReasonPath) Get(t testing.TB) *oc.Component_LastSwitchoverReason {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/last-switchover-reason with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LastSwitchoverReasonPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_LastSwitchoverReason {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_LastSwitchoverReason
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_LastSwitchoverReason{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_LastSwitchoverReason", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_LastSwitchoverReason{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/last-switchover-reason with a ONCE subscription.
func (n *Component_LastSwitchoverReasonPathAny) Get(t testing.TB) []*oc.Component_LastSwitchoverReason {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_LastSwitchoverReason
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-reason with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverReasonPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_LastSwitchoverReason {
	t.Helper()
	c := &oc.CollectionComponent_LastSwitchoverReason{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_LastSwitchoverReason) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_LastSwitchoverReason{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_LastSwitchoverReason)))
		return false
	})
	return c
}

func watch_Component_LastSwitchoverReasonPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_LastSwitchoverReason) bool) *oc.Component_LastSwitchoverReasonWatcher {
	t.Helper()
	w := &oc.Component_LastSwitchoverReasonWatcher{}
	gs := &oc.Component_LastSwitchoverReason{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_LastSwitchoverReason", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_LastSwitchoverReason{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_LastSwitchoverReason)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-reason with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverReasonPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_LastSwitchoverReason) bool) *oc.Component_LastSwitchoverReasonWatcher {
	t.Helper()
	return watch_Component_LastSwitchoverReasonPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/last-switchover-reason with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LastSwitchoverReasonPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_LastSwitchoverReason) *oc.QualifiedComponent_LastSwitchoverReason {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_LastSwitchoverReason) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/last-switchover-reason failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-reason to the batch object.
func (n *Component_LastSwitchoverReasonPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-reason with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverReasonPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_LastSwitchoverReason {
	t.Helper()
	c := &oc.CollectionComponent_LastSwitchoverReason{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_LastSwitchoverReason) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastSwitchoverReasonPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_LastSwitchoverReason) bool) *oc.Component_LastSwitchoverReasonWatcher {
	t.Helper()
	w := &oc.Component_LastSwitchoverReasonWatcher{}
	structs := map[string]*oc.Component_LastSwitchoverReason{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_LastSwitchoverReason{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_LastSwitchoverReason", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_LastSwitchoverReason{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_LastSwitchoverReason)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-reason with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverReasonPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_LastSwitchoverReason) bool) *oc.Component_LastSwitchoverReasonWatcher {
	t.Helper()
	return watch_Component_LastSwitchoverReasonPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-reason to the batch object.
func (n *Component_LastSwitchoverReasonPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/state/last-switchover-reason/details with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LastSwitchoverReason_DetailsPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_LastSwitchoverReason{}
	md, ok := oc.Lookup(t, n, "Component_LastSwitchoverReason", goStruct, true, false)
	if ok {
		return convertComponent_LastSwitchoverReason_DetailsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/last-switchover-reason/details with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LastSwitchoverReason_DetailsPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LastSwitchoverReason_DetailsPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_LastSwitchoverReason{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_LastSwitchoverReason", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_LastSwitchoverReason_DetailsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a ONCE subscription.
func (n *Component_LastSwitchoverReason_DetailsPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverReason_DetailsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastSwitchoverReason_DetailsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_LastSwitchoverReason{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_LastSwitchoverReason", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_LastSwitchoverReason_DetailsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverReason_DetailsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_LastSwitchoverReason_DetailsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/last-switchover-reason/details with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LastSwitchoverReason_DetailsPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/last-switchover-reason/details failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-reason/details to the batch object.
func (n *Component_LastSwitchoverReason_DetailsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverReason_DetailsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastSwitchoverReason_DetailsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Component_LastSwitchoverReason{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_LastSwitchoverReason{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_LastSwitchoverReason", structs[pre], queryPath, true, false)
			qv := convertComponent_LastSwitchoverReason_DetailsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-reason/details with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverReason_DetailsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_LastSwitchoverReason_DetailsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-reason/details to the batch object.
func (n *Component_LastSwitchoverReason_DetailsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_LastSwitchoverReason_DetailsPath extracts the value of the leaf Details from its parent oc.Component_LastSwitchoverReason
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_LastSwitchoverReason_DetailsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_LastSwitchoverReason) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Details
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/last-switchover-reason/trigger with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LastSwitchoverReason_TriggerPath) Lookup(t testing.TB) *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	t.Helper()
	goStruct := &oc.Component_LastSwitchoverReason{}
	md, ok := oc.Lookup(t, n, "Component_LastSwitchoverReason", goStruct, true, false)
	if ok {
		return convertComponent_LastSwitchoverReason_TriggerPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/last-switchover-reason/trigger with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LastSwitchoverReason_TriggerPath) Get(t testing.TB) oc.E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/last-switchover-reason/trigger with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LastSwitchoverReason_TriggerPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_LastSwitchoverReason{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_LastSwitchoverReason", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_LastSwitchoverReason_TriggerPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/last-switchover-reason/trigger with a ONCE subscription.
func (n *Component_LastSwitchoverReason_TriggerPathAny) Get(t testing.TB) []oc.E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-reason/trigger with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverReason_TriggerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastSwitchoverReason_TriggerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) bool) *oc.E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTriggerWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTriggerWatcher{}
	gs := &oc.Component_LastSwitchoverReason{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_LastSwitchoverReason", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_LastSwitchoverReason_TriggerPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-reason/trigger with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverReason_TriggerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) bool) *oc.E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTriggerWatcher {
	t.Helper()
	return watch_Component_LastSwitchoverReason_TriggerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/last-switchover-reason/trigger with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LastSwitchoverReason_TriggerPath) Await(t testing.TB, timeout time.Duration, val oc.E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/last-switchover-reason/trigger failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-reason/trigger to the batch object.
func (n *Component_LastSwitchoverReason_TriggerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-reason/trigger with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverReason_TriggerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastSwitchoverReason_TriggerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) bool) *oc.E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTriggerWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTriggerWatcher{}
	structs := map[string]*oc.Component_LastSwitchoverReason{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_LastSwitchoverReason{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_LastSwitchoverReason", structs[pre], queryPath, true, false)
			qv := convertComponent_LastSwitchoverReason_TriggerPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-reason/trigger with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverReason_TriggerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) bool) *oc.E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTriggerWatcher {
	t.Helper()
	return watch_Component_LastSwitchoverReason_TriggerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-reason/trigger to the batch object.
func (n *Component_LastSwitchoverReason_TriggerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_LastSwitchoverReason_TriggerPath extracts the value of the leaf Trigger from its parent oc.Component_LastSwitchoverReason
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger.
func convertComponent_LastSwitchoverReason_TriggerPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_LastSwitchoverReason) *oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	t.Helper()
	qv := &oc.QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger{
		Metadata: md,
	}
	val := parent.Trigger
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/last-switchover-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LastSwitchoverTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_LastSwitchoverTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/last-switchover-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LastSwitchoverTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/last-switchover-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LastSwitchoverTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_LastSwitchoverTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/last-switchover-time with a ONCE subscription.
func (n *Component_LastSwitchoverTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastSwitchoverTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_LastSwitchoverTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_LastSwitchoverTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/last-switchover-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LastSwitchoverTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/last-switchover-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-time to the batch object.
func (n *Component_LastSwitchoverTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/last-switchover-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LastSwitchoverTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LastSwitchoverTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertComponent_LastSwitchoverTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/last-switchover-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LastSwitchoverTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_LastSwitchoverTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/last-switchover-time to the batch object.
func (n *Component_LastSwitchoverTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_LastSwitchoverTimePath extracts the value of the leaf LastSwitchoverTime from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_LastSwitchoverTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.LastSwitchoverTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/location with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_LocationPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_LocationPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/location with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_LocationPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/location with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_LocationPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_LocationPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/location with a ONCE subscription.
func (n *Component_LocationPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/location with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LocationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LocationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_LocationPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/location with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LocationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_LocationPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/location with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_LocationPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/location failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/location to the batch object.
func (n *Component_LocationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/location with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_LocationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_LocationPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertComponent_LocationPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/location with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_LocationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_LocationPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/location to the batch object.
func (n *Component_LocationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_LocationPath extracts the value of the leaf Location from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_LocationPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Location
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/memory with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_MemoryPath) Lookup(t testing.TB) *oc.QualifiedComponent_Memory {
	t.Helper()
	goStruct := &oc.Component_Memory{}
	md, ok := oc.Lookup(t, n, "Component_Memory", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Memory{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/memory with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_MemoryPath) Get(t testing.TB) *oc.Component_Memory {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/memory with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_MemoryPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Memory {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Memory
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Memory", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Memory{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/memory with a ONCE subscription.
func (n *Component_MemoryPathAny) Get(t testing.TB) []*oc.Component_Memory {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Memory
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/memory with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_MemoryPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Memory {
	t.Helper()
	c := &oc.CollectionComponent_Memory{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Memory) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Memory{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Memory)))
		return false
	})
	return c
}

func watch_Component_MemoryPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Memory) bool) *oc.Component_MemoryWatcher {
	t.Helper()
	w := &oc.Component_MemoryWatcher{}
	gs := &oc.Component_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Memory", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_Memory{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Memory)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/memory with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_MemoryPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Memory) bool) *oc.Component_MemoryWatcher {
	t.Helper()
	return watch_Component_MemoryPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/memory with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_MemoryPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Memory) *oc.QualifiedComponent_Memory {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Memory) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/memory failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/memory to the batch object.
func (n *Component_MemoryPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/memory with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_MemoryPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Memory {
	t.Helper()
	c := &oc.CollectionComponent_Memory{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Memory) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_MemoryPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Memory) bool) *oc.Component_MemoryWatcher {
	t.Helper()
	w := &oc.Component_MemoryWatcher{}
	structs := map[string]*oc.Component_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Memory{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Memory", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_Memory{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Memory)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/memory with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_MemoryPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Memory) bool) *oc.Component_MemoryWatcher {
	t.Helper()
	return watch_Component_MemoryPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/memory to the batch object.
func (n *Component_MemoryPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/state/memory/available with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Memory_AvailablePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Memory{}
	md, ok := oc.Lookup(t, n, "Component_Memory", goStruct, true, false)
	if ok {
		return convertComponent_Memory_AvailablePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/memory/available with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Memory_AvailablePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/memory/available with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Memory_AvailablePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Memory", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Memory_AvailablePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/memory/available with a ONCE subscription.
func (n *Component_Memory_AvailablePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/memory/available with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Memory_AvailablePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Memory_AvailablePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Memory", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Memory_AvailablePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/memory/available with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Memory_AvailablePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Memory_AvailablePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/memory/available with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Memory_AvailablePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/memory/available failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/memory/available to the batch object.
func (n *Component_Memory_AvailablePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/memory/available with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Memory_AvailablePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Memory_AvailablePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Memory{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Memory", structs[pre], queryPath, true, false)
			qv := convertComponent_Memory_AvailablePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/memory/available with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Memory_AvailablePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Memory_AvailablePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/memory/available to the batch object.
func (n *Component_Memory_AvailablePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Memory_AvailablePath extracts the value of the leaf Available from its parent oc.Component_Memory
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Memory_AvailablePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Memory) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Available
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/memory/utilized with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Memory_UtilizedPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Memory{}
	md, ok := oc.Lookup(t, n, "Component_Memory", goStruct, true, false)
	if ok {
		return convertComponent_Memory_UtilizedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/memory/utilized with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Memory_UtilizedPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/memory/utilized with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Memory_UtilizedPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Memory", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Memory_UtilizedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/memory/utilized with a ONCE subscription.
func (n *Component_Memory_UtilizedPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/memory/utilized with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Memory_UtilizedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Memory_UtilizedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Memory", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Memory_UtilizedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/memory/utilized with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Memory_UtilizedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Memory_UtilizedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/memory/utilized with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Memory_UtilizedPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/memory/utilized failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/memory/utilized to the batch object.
func (n *Component_Memory_UtilizedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/memory/utilized with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Memory_UtilizedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Memory_UtilizedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Memory{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Memory{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Memory", structs[pre], queryPath, true, false)
			qv := convertComponent_Memory_UtilizedPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/memory/utilized with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Memory_UtilizedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Memory_UtilizedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/memory/utilized to the batch object.
func (n *Component_Memory_UtilizedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Memory_UtilizedPath extracts the value of the leaf Utilized from its parent oc.Component_Memory
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Memory_UtilizedPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Memory) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Utilized
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/mfg-date with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_MfgDatePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_MfgDatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/mfg-date with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_MfgDatePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/mfg-date with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_MfgDatePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_MfgDatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/mfg-date with a ONCE subscription.
func (n *Component_MfgDatePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/mfg-date with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_MfgDatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_MfgDatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_MfgDatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/mfg-date with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_MfgDatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_MfgDatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/mfg-date with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_MfgDatePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/mfg-date failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/mfg-date to the batch object.
func (n *Component_MfgDatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/mfg-date with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_MfgDatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_MfgDatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertComponent_MfgDatePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/mfg-date with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_MfgDatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_MfgDatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/mfg-date to the batch object.
func (n *Component_MfgDatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_MfgDatePath extracts the value of the leaf MfgDate from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_MfgDatePath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MfgDate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/mfg-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_MfgNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_MfgNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/mfg-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_MfgNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/mfg-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_MfgNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_MfgNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/mfg-name with a ONCE subscription.
func (n *Component_MfgNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/mfg-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_MfgNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_MfgNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_MfgNamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/mfg-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_MfgNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_MfgNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/mfg-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_MfgNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/mfg-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/mfg-name to the batch object.
func (n *Component_MfgNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/mfg-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_MfgNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_MfgNamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertComponent_MfgNamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/mfg-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_MfgNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_MfgNamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/mfg-name to the batch object.
func (n *Component_MfgNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_MfgNamePath extracts the value of the leaf MfgName from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_MfgNamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MfgName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/name with a ONCE subscription.
func (n *Component_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/name to the batch object.
func (n *Component_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertComponent_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/name to the batch object.
func (n *Component_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_NamePath extracts the value of the leaf Name from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Name
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/oper-status with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OperStatusPath) Lookup(t testing.TB) *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_OperStatusPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/oper-status with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OperStatusPath) Get(t testing.TB) oc.E_PlatformTypes_COMPONENT_OPER_STATUS {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/oper-status with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OperStatusPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OperStatusPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/oper-status with a ONCE subscription.
func (n *Component_OperStatusPathAny) Get(t testing.TB) []oc.E_PlatformTypes_COMPONENT_OPER_STATUS {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PlatformTypes_COMPONENT_OPER_STATUS
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/oper-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OperStatusPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_COMPONENT_OPER_STATUS {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_COMPONENT_OPER_STATUS{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OperStatusPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) bool) *oc.E_PlatformTypes_COMPONENT_OPER_STATUSWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_COMPONENT_OPER_STATUSWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OperStatusPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/oper-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OperStatusPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) bool) *oc.E_PlatformTypes_COMPONENT_OPER_STATUSWatcher {
	t.Helper()
	return watch_Component_OperStatusPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/oper-status with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OperStatusPath) Await(t testing.TB, timeout time.Duration, val oc.E_PlatformTypes_COMPONENT_OPER_STATUS) *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/oper-status failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/oper-status to the batch object.
func (n *Component_OperStatusPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/oper-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OperStatusPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_COMPONENT_OPER_STATUS {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_COMPONENT_OPER_STATUS{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OperStatusPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) bool) *oc.E_PlatformTypes_COMPONENT_OPER_STATUSWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_COMPONENT_OPER_STATUSWatcher{}
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
			qv := convertComponent_OperStatusPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/oper-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OperStatusPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) bool) *oc.E_PlatformTypes_COMPONENT_OPER_STATUSWatcher {
	t.Helper()
	return watch_Component_OperStatusPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/oper-status to the batch object.
func (n *Component_OperStatusPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OperStatusPath extracts the value of the leaf OperStatus from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS.
func convertComponent_OperStatusPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS {
	t.Helper()
	qv := &oc.QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS{
		Metadata: md,
	}
	val := parent.OperStatus
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannelPath) Lookup(t testing.TB) *oc.QualifiedComponent_OpticalChannel {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_OpticalChannel{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannelPath) Get(t testing.TB) *oc.Component_OpticalChannel {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_OpticalChannel {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_OpticalChannel
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_OpticalChannel{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel with a ONCE subscription.
func (n *Component_OpticalChannelPathAny) Get(t testing.TB) []*oc.Component_OpticalChannel {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_OpticalChannel
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_OpticalChannel{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_OpticalChannel)))
		return false
	})
	return c
}

func watch_Component_OpticalChannelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel) bool) *oc.Component_OpticalChannelWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannelWatcher{}
	gs := &oc.Component_OpticalChannel{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_OpticalChannel{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel) bool) *oc.Component_OpticalChannelWatcher {
	t.Helper()
	return watch_Component_OpticalChannelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannelPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_OpticalChannel) *oc.QualifiedComponent_OpticalChannel {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_OpticalChannel) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel to the batch object.
func (n *Component_OpticalChannelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannelPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel) bool) *oc.Component_OpticalChannelWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannelWatcher{}
	structs := map[string]*oc.Component_OpticalChannel{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_OpticalChannel{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel) bool) *oc.Component_OpticalChannelWatcher {
	t.Helper()
	return watch_Component_OpticalChannelPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel to the batch object.
func (n *Component_OpticalChannelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPath) Lookup(t testing.TB) *oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPath) Get(t testing.TB) *oc.Component_OpticalChannel_CarrierFrequencyOffset {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset with a ONCE subscription.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPathAny) Get(t testing.TB) []*oc.Component_OpticalChannel_CarrierFrequencyOffset {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_OpticalChannel_CarrierFrequencyOffset
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_CarrierFrequencyOffset {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_CarrierFrequencyOffset{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_OpticalChannel_CarrierFrequencyOffset)))
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffsetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset) bool) *oc.Component_OpticalChannel_CarrierFrequencyOffsetWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_CarrierFrequencyOffsetWatcher{}
	gs := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset) bool) *oc.Component_OpticalChannel_CarrierFrequencyOffsetWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffsetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_OpticalChannel_CarrierFrequencyOffset) *oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_CarrierFrequencyOffset {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_CarrierFrequencyOffset{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffsetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset) bool) *oc.Component_OpticalChannel_CarrierFrequencyOffsetWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_CarrierFrequencyOffsetWatcher{}
	structs := map[string]*oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_CarrierFrequencyOffset) bool) *oc.Component_OpticalChannel_CarrierFrequencyOffsetWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffsetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffsetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_CarrierFrequencyOffset_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg with a ONCE subscription.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_CarrierFrequencyOffset_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/avg to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_CarrierFrequencyOffset_AvgPath extracts the value of the leaf Avg from its parent oc.Component_OpticalChannel_CarrierFrequencyOffset
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_CarrierFrequencyOffset_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_CarrierFrequencyOffset) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_CarrierFrequencyOffset_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant with a ONCE subscription.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_CarrierFrequencyOffset_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/instant to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_CarrierFrequencyOffset_InstantPath extracts the value of the leaf Instant from its parent oc.Component_OpticalChannel_CarrierFrequencyOffset
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_CarrierFrequencyOffset_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_CarrierFrequencyOffset) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_CarrierFrequencyOffset_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval with a ONCE subscription.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_CarrierFrequencyOffset_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/interval to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_CarrierFrequencyOffset_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_OpticalChannel_CarrierFrequencyOffset
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_CarrierFrequencyOffset_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_CarrierFrequencyOffset) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max with a ONCE subscription.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxPath extracts the value of the leaf Max from its parent oc.Component_OpticalChannel_CarrierFrequencyOffset
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_CarrierFrequencyOffset) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time with a ONCE subscription.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/max-time to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_OpticalChannel_CarrierFrequencyOffset
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_CarrierFrequencyOffset_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_CarrierFrequencyOffset) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_CarrierFrequencyOffset_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min with a ONCE subscription.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_CarrierFrequencyOffset_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_CarrierFrequencyOffset_MinPath extracts the value of the leaf Min from its parent oc.Component_OpticalChannel_CarrierFrequencyOffset
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_CarrierFrequencyOffset_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_CarrierFrequencyOffset) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_CarrierFrequencyOffset_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time with a ONCE subscription.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_CarrierFrequencyOffset_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_CarrierFrequencyOffset_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_CarrierFrequencyOffset{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_CarrierFrequencyOffset{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_CarrierFrequencyOffset", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_CarrierFrequencyOffset_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_CarrierFrequencyOffset_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/carrier-frequency-offset/min-time to the batch object.
func (n *Component_OpticalChannel_CarrierFrequencyOffset_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_CarrierFrequencyOffset_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_OpticalChannel_CarrierFrequencyOffset
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_CarrierFrequencyOffset_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_CarrierFrequencyOffset) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ChromaticDispersionPath) Lookup(t testing.TB) *oc.QualifiedComponent_OpticalChannel_ChromaticDispersion {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ChromaticDispersion", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_OpticalChannel_ChromaticDispersion{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ChromaticDispersionPath) Get(t testing.TB) *oc.Component_OpticalChannel_ChromaticDispersion {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ChromaticDispersionPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_OpticalChannel_ChromaticDispersion {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_OpticalChannel_ChromaticDispersion
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_OpticalChannel_ChromaticDispersion{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion with a ONCE subscription.
func (n *Component_OpticalChannel_ChromaticDispersionPathAny) Get(t testing.TB) []*oc.Component_OpticalChannel_ChromaticDispersion {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_OpticalChannel_ChromaticDispersion
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_ChromaticDispersion {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_ChromaticDispersion{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_ChromaticDispersion) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_OpticalChannel_ChromaticDispersion{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_OpticalChannel_ChromaticDispersion)))
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ChromaticDispersion) bool) *oc.Component_OpticalChannel_ChromaticDispersionWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_ChromaticDispersionWatcher{}
	gs := &oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_OpticalChannel_ChromaticDispersion{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_ChromaticDispersion)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ChromaticDispersion) bool) *oc.Component_OpticalChannel_ChromaticDispersionWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ChromaticDispersionPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_OpticalChannel_ChromaticDispersion) *oc.QualifiedComponent_OpticalChannel_ChromaticDispersion {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_OpticalChannel_ChromaticDispersion) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_ChromaticDispersion {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_ChromaticDispersion{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_ChromaticDispersion) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ChromaticDispersion) bool) *oc.Component_OpticalChannel_ChromaticDispersionWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_ChromaticDispersionWatcher{}
	structs := map[string]*oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ChromaticDispersion{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_OpticalChannel_ChromaticDispersion{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_ChromaticDispersion)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_ChromaticDispersion) bool) *oc.Component_OpticalChannel_ChromaticDispersionWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ChromaticDispersion", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ChromaticDispersion_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ChromaticDispersion_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg with a ONCE subscription.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ChromaticDispersion_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ChromaticDispersion{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ChromaticDispersion_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/avg to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ChromaticDispersion_AvgPath extracts the value of the leaf Avg from its parent oc.Component_OpticalChannel_ChromaticDispersion
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ChromaticDispersion_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ChromaticDispersion) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ChromaticDispersion", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ChromaticDispersion_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ChromaticDispersion_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant with a ONCE subscription.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ChromaticDispersion_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ChromaticDispersion{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ChromaticDispersion_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/instant to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ChromaticDispersion_InstantPath extracts the value of the leaf Instant from its parent oc.Component_OpticalChannel_ChromaticDispersion
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ChromaticDispersion_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ChromaticDispersion) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ChromaticDispersion", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ChromaticDispersion_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ChromaticDispersion_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval with a ONCE subscription.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ChromaticDispersion_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ChromaticDispersion{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ChromaticDispersion_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/interval to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ChromaticDispersion_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_OpticalChannel_ChromaticDispersion
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ChromaticDispersion_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ChromaticDispersion) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ChromaticDispersion", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ChromaticDispersion_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ChromaticDispersion_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max with a ONCE subscription.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ChromaticDispersion_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ChromaticDispersion{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ChromaticDispersion_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ChromaticDispersion_MaxPath extracts the value of the leaf Max from its parent oc.Component_OpticalChannel_ChromaticDispersion
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ChromaticDispersion_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ChromaticDispersion) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ChromaticDispersion", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ChromaticDispersion_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ChromaticDispersion_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time with a ONCE subscription.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ChromaticDispersion_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ChromaticDispersion{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ChromaticDispersion_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/max-time to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ChromaticDispersion_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_OpticalChannel_ChromaticDispersion
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ChromaticDispersion_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ChromaticDispersion) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ChromaticDispersion", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ChromaticDispersion_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ChromaticDispersion_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min with a ONCE subscription.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ChromaticDispersion_MinPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_MinPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ChromaticDispersion{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ChromaticDispersion_MinPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_MinPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ChromaticDispersion_MinPath extracts the value of the leaf Min from its parent oc.Component_OpticalChannel_ChromaticDispersion
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_ChromaticDispersion_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ChromaticDispersion) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_ChromaticDispersion", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_ChromaticDispersion_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_ChromaticDispersion{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_ChromaticDispersion_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time with a ONCE subscription.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_ChromaticDispersion_MinTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_ChromaticDispersion_MinTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_ChromaticDispersion{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_ChromaticDispersion{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_ChromaticDispersion", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_ChromaticDispersion_MinTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_ChromaticDispersion_MinTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/chromatic-dispersion/min-time to the batch object.
func (n *Component_OpticalChannel_ChromaticDispersion_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_ChromaticDispersion_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_OpticalChannel_ChromaticDispersion
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_ChromaticDispersion_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_ChromaticDispersion) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/esnr with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_EsnrPath) Lookup(t testing.TB) *oc.QualifiedComponent_OpticalChannel_Esnr {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Esnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Esnr", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_OpticalChannel_Esnr{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/esnr with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_EsnrPath) Get(t testing.TB) *oc.Component_OpticalChannel_Esnr {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/esnr with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_EsnrPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_OpticalChannel_Esnr {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_OpticalChannel_Esnr
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Esnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Esnr", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_OpticalChannel_Esnr{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/esnr with a ONCE subscription.
func (n *Component_OpticalChannel_EsnrPathAny) Get(t testing.TB) []*oc.Component_OpticalChannel_Esnr {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_OpticalChannel_Esnr
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/esnr with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_EsnrPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_Esnr {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_Esnr{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_Esnr) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_OpticalChannel_Esnr{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_OpticalChannel_Esnr)))
		return false
	})
	return c
}

func watch_Component_OpticalChannel_EsnrPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_Esnr) bool) *oc.Component_OpticalChannel_EsnrWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_EsnrWatcher{}
	gs := &oc.Component_OpticalChannel_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Esnr", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_OpticalChannel_Esnr{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_Esnr)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/esnr with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_EsnrPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_Esnr) bool) *oc.Component_OpticalChannel_EsnrWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_EsnrPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/esnr with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_EsnrPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_OpticalChannel_Esnr) *oc.QualifiedComponent_OpticalChannel_Esnr {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_OpticalChannel_Esnr) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/esnr failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/esnr to the batch object.
func (n *Component_OpticalChannel_EsnrPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/esnr with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_EsnrPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_OpticalChannel_Esnr {
	t.Helper()
	c := &oc.CollectionComponent_OpticalChannel_Esnr{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_OpticalChannel_Esnr) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_EsnrPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_Esnr) bool) *oc.Component_OpticalChannel_EsnrWatcher {
	t.Helper()
	w := &oc.Component_OpticalChannel_EsnrWatcher{}
	structs := map[string]*oc.Component_OpticalChannel_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Esnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Esnr", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_OpticalChannel_Esnr{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_OpticalChannel_Esnr)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/esnr with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_EsnrPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_OpticalChannel_Esnr) bool) *oc.Component_OpticalChannel_EsnrWatcher {
	t.Helper()
	return watch_Component_OpticalChannel_EsnrPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/esnr to the batch object.
func (n *Component_OpticalChannel_EsnrPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/esnr/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_Esnr_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Esnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Esnr", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_Esnr_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/esnr/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_Esnr_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/esnr/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_Esnr_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Esnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Esnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_Esnr_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/esnr/avg with a ONCE subscription.
func (n *Component_OpticalChannel_Esnr_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/esnr/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Esnr_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Esnr_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Esnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_Esnr_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/esnr/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Esnr_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Esnr_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/esnr/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_Esnr_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/esnr/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/esnr/avg to the batch object.
func (n *Component_OpticalChannel_Esnr_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/esnr/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Esnr_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Esnr_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Esnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Esnr", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_Esnr_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/esnr/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Esnr_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Esnr_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/esnr/avg to the batch object.
func (n *Component_OpticalChannel_Esnr_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_Esnr_AvgPath extracts the value of the leaf Avg from its parent oc.Component_OpticalChannel_Esnr
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_Esnr_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_Esnr) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/esnr/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_Esnr_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Esnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Esnr", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_Esnr_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/esnr/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_Esnr_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/esnr/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_Esnr_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Esnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Esnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_Esnr_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/esnr/instant with a ONCE subscription.
func (n *Component_OpticalChannel_Esnr_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/esnr/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Esnr_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Esnr_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_OpticalChannel_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Esnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_Esnr_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/esnr/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Esnr_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Esnr_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/esnr/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_Esnr_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/esnr/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/esnr/instant to the batch object.
func (n *Component_OpticalChannel_Esnr_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/esnr/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Esnr_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Esnr_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Esnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Esnr", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_Esnr_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/esnr/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Esnr_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Esnr_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/esnr/instant to the batch object.
func (n *Component_OpticalChannel_Esnr_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_Esnr_InstantPath extracts the value of the leaf Instant from its parent oc.Component_OpticalChannel_Esnr
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_OpticalChannel_Esnr_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_Esnr) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/optical-channel/state/esnr/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_OpticalChannel_Esnr_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_OpticalChannel_Esnr{}
	md, ok := oc.Lookup(t, n, "Component_OpticalChannel_Esnr", goStruct, true, false)
	if ok {
		return convertComponent_OpticalChannel_Esnr_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/optical-channel/state/esnr/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_OpticalChannel_Esnr_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/optical-channel/state/esnr/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_OpticalChannel_Esnr_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_OpticalChannel_Esnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_OpticalChannel_Esnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_OpticalChannel_Esnr_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/optical-channel/state/esnr/interval with a ONCE subscription.
func (n *Component_OpticalChannel_Esnr_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/esnr/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Esnr_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Esnr_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_OpticalChannel_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_OpticalChannel_Esnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_OpticalChannel_Esnr_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/esnr/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Esnr_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Esnr_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/optical-channel/state/esnr/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_OpticalChannel_Esnr_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/optical-channel/state/esnr/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/esnr/interval to the batch object.
func (n *Component_OpticalChannel_Esnr_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/optical-channel/state/esnr/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_OpticalChannel_Esnr_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_OpticalChannel_Esnr_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_OpticalChannel_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_OpticalChannel_Esnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_OpticalChannel_Esnr", structs[pre], queryPath, true, false)
			qv := convertComponent_OpticalChannel_Esnr_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/optical-channel/state/esnr/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_OpticalChannel_Esnr_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_OpticalChannel_Esnr_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/optical-channel/state/esnr/interval to the batch object.
func (n *Component_OpticalChannel_Esnr_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_OpticalChannel_Esnr_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_OpticalChannel_Esnr
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_OpticalChannel_Esnr_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_OpticalChannel_Esnr) *oc.QualifiedUint64 {
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
