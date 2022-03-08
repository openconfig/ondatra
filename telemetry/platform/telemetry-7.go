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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/ethernet-pmd with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_EthernetPmdPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_EthernetPmdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/ethernet-pmd with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_EthernetPmdPath) Get(t testing.TB) oc.E_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_EthernetPmdPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_EthernetPmdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd with a ONCE subscription.
func (n *Component_Transceiver_EthernetPmdPathAny) Get(t testing.TB) []oc.E_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_ETHERNET_PMD_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_EthernetPmdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_ETHERNET_PMD_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_EthernetPmdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool) *oc.E_TransportTypes_ETHERNET_PMD_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_ETHERNET_PMD_TYPEWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_EthernetPmdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_EthernetPmdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool) *oc.E_TransportTypes_ETHERNET_PMD_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_EthernetPmdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_EthernetPmdPath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_ETHERNET_PMD_TYPE) *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/ethernet-pmd failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/ethernet-pmd to the batch object.
func (n *Component_Transceiver_EthernetPmdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_EthernetPmdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_ETHERNET_PMD_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_EthernetPmdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool) *oc.E_TransportTypes_ETHERNET_PMD_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_EthernetPmdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/ethernet-pmd to the batch object.
func (n *Component_Transceiver_EthernetPmdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_EthernetPmdPath extracts the value of the leaf EthernetPmd from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE.
func convertComponent_Transceiver_EthernetPmdPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE{
		Metadata: md,
	}
	val := parent.EthernetPmd
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_EthernetPmdPreconfPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Get(t testing.TB) oc.E_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_EthernetPmdPreconfPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_EthernetPmdPreconfPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf with a ONCE subscription.
func (n *Component_Transceiver_EthernetPmdPreconfPathAny) Get(t testing.TB) []oc.E_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_ETHERNET_PMD_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_ETHERNET_PMD_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_EthernetPmdPreconfPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool) *oc.E_TransportTypes_ETHERNET_PMD_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_ETHERNET_PMD_TYPEWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_EthernetPmdPreconfPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool) *oc.E_TransportTypes_ETHERNET_PMD_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_EthernetPmdPreconfPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_ETHERNET_PMD_TYPE) *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf to the batch object.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_EthernetPmdPreconfPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_ETHERNET_PMD_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_EthernetPmdPreconfPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE) bool) *oc.E_TransportTypes_ETHERNET_PMD_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_EthernetPmdPreconfPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/ethernet-pmd-preconf to the batch object.
func (n *Component_Transceiver_EthernetPmdPreconfPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_EthernetPmdPreconfPath extracts the value of the leaf EthernetPmdPreconf from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE.
func convertComponent_Transceiver_EthernetPmdPreconfPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE{
		Metadata: md,
	}
	val := parent.EthernetPmdPreconf
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/fault-condition with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FaultConditionPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_FaultConditionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/fault-condition with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FaultConditionPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/fault-condition with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FaultConditionPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FaultConditionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/fault-condition with a ONCE subscription.
func (n *Component_Transceiver_FaultConditionPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fault-condition with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FaultConditionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_FaultConditionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_FaultConditionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fault-condition with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FaultConditionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Component_Transceiver_FaultConditionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/fault-condition with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_FaultConditionPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/fault-condition failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fault-condition to the batch object.
func (n *Component_Transceiver_FaultConditionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fault-condition with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FaultConditionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fault-condition with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FaultConditionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Component_Transceiver_FaultConditionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fault-condition to the batch object.
func (n *Component_Transceiver_FaultConditionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_FaultConditionPath extracts the value of the leaf FaultCondition from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertComponent_Transceiver_FaultConditionPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.FaultCondition
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/fec-corrected-bits with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FecCorrectedBitsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_FecCorrectedBitsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/fec-corrected-bits with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FecCorrectedBitsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bits with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FecCorrectedBitsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FecCorrectedBitsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bits with a ONCE subscription.
func (n *Component_Transceiver_FecCorrectedBitsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bits with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecCorrectedBitsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_FecCorrectedBitsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_FecCorrectedBitsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bits with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecCorrectedBitsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_FecCorrectedBitsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bits with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_FecCorrectedBitsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/fec-corrected-bits failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-corrected-bits to the batch object.
func (n *Component_Transceiver_FecCorrectedBitsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bits with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecCorrectedBitsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bits with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecCorrectedBitsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_FecCorrectedBitsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-corrected-bits to the batch object.
func (n *Component_Transceiver_FecCorrectedBitsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_FecCorrectedBitsPath extracts the value of the leaf FecCorrectedBits from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_FecCorrectedBitsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FecCorrectedBits
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FecCorrectedBytesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_FecCorrectedBytesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FecCorrectedBytesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FecCorrectedBytesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FecCorrectedBytesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes with a ONCE subscription.
func (n *Component_Transceiver_FecCorrectedBytesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecCorrectedBytesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_FecCorrectedBytesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_FecCorrectedBytesPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecCorrectedBytesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_FecCorrectedBytesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_FecCorrectedBytesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes to the batch object.
func (n *Component_Transceiver_FecCorrectedBytesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecCorrectedBytesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecCorrectedBytesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_FecCorrectedBytesPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-corrected-bytes to the batch object.
func (n *Component_Transceiver_FecCorrectedBytesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_FecCorrectedBytesPath extracts the value of the leaf FecCorrectedBytes from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_FecCorrectedBytesPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FecCorrectedBytes
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/fec-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FecModePath) Lookup(t testing.TB) *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_FecModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/fec-mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FecModePath) Get(t testing.TB) oc.E_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/fec-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FecModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FecModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/fec-mode with a ONCE subscription.
func (n *Component_Transceiver_FecModePathAny) Get(t testing.TB) []oc.E_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PlatformTypes_FEC_MODE_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_FEC_MODE_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_FecModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE) bool) *oc.E_PlatformTypes_FEC_MODE_TYPEWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_FEC_MODE_TYPEWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_FecModePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE) bool) *oc.E_PlatformTypes_FEC_MODE_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_FecModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/fec-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_FecModePath) Await(t testing.TB, timeout time.Duration, val oc.E_PlatformTypes_FEC_MODE_TYPE) *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/fec-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-mode to the batch object.
func (n *Component_Transceiver_FecModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_FEC_MODE_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE) bool) *oc.E_PlatformTypes_FEC_MODE_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_FecModePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-mode to the batch object.
func (n *Component_Transceiver_FecModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_FecModePath extracts the value of the leaf FecMode from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE.
func convertComponent_Transceiver_FecModePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE{
		Metadata: md,
	}
	val := parent.FecMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/fec-status with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FecStatusPath) Lookup(t testing.TB) *oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_FecStatusPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/fec-status with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FecStatusPath) Get(t testing.TB) oc.E_PlatformTypes_FEC_STATUS_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/fec-status with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FecStatusPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FecStatusPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/fec-status with a ONCE subscription.
func (n *Component_Transceiver_FecStatusPathAny) Get(t testing.TB) []oc.E_PlatformTypes_FEC_STATUS_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PlatformTypes_FEC_STATUS_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecStatusPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_FEC_STATUS_TYPE {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_FEC_STATUS_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_FecStatusPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE) bool) *oc.E_PlatformTypes_FEC_STATUS_TYPEWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_FEC_STATUS_TYPEWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_FecStatusPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecStatusPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE) bool) *oc.E_PlatformTypes_FEC_STATUS_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_FecStatusPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/fec-status with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_FecStatusPath) Await(t testing.TB, timeout time.Duration, val oc.E_PlatformTypes_FEC_STATUS_TYPE) *oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/fec-status failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-status to the batch object.
func (n *Component_Transceiver_FecStatusPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecStatusPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_FEC_STATUS_TYPE {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_FEC_STATUS_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecStatusPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE) bool) *oc.E_PlatformTypes_FEC_STATUS_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_FecStatusPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-status to the batch object.
func (n *Component_Transceiver_FecStatusPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_FecStatusPath extracts the value of the leaf FecStatus from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE.
func convertComponent_Transceiver_FecStatusPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_PlatformTypes_FEC_STATUS_TYPE{
		Metadata: md,
	}
	val := parent.FecStatus
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FecUncorrectableBlocksPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_FecUncorrectableBlocksPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FecUncorrectableBlocksPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FecUncorrectableBlocksPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FecUncorrectableBlocksPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks with a ONCE subscription.
func (n *Component_Transceiver_FecUncorrectableBlocksPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecUncorrectableBlocksPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_FecUncorrectableBlocksPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_FecUncorrectableBlocksPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecUncorrectableBlocksPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_FecUncorrectableBlocksPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_FecUncorrectableBlocksPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks to the batch object.
func (n *Component_Transceiver_FecUncorrectableBlocksPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecUncorrectableBlocksPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecUncorrectableBlocksPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_FecUncorrectableBlocksPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-blocks to the batch object.
func (n *Component_Transceiver_FecUncorrectableBlocksPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_FecUncorrectableBlocksPath extracts the value of the leaf FecUncorrectableBlocks from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_FecUncorrectableBlocksPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FecUncorrectableBlocks
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FecUncorrectableWordsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_FecUncorrectableWordsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FecUncorrectableWordsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FecUncorrectableWordsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FecUncorrectableWordsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words with a ONCE subscription.
func (n *Component_Transceiver_FecUncorrectableWordsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecUncorrectableWordsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_FecUncorrectableWordsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_FecUncorrectableWordsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecUncorrectableWordsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_FecUncorrectableWordsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_FecUncorrectableWordsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words to the batch object.
func (n *Component_Transceiver_FecUncorrectableWordsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FecUncorrectableWordsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FecUncorrectableWordsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_FecUncorrectableWordsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/fec-uncorrectable-words to the batch object.
func (n *Component_Transceiver_FecUncorrectableWordsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_FecUncorrectableWordsPath extracts the value of the leaf FecUncorrectableWords from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_FecUncorrectableWordsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.FecUncorrectableWords
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/form-factor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FormFactorPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_FormFactorPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/form-factor with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FormFactorPath) Get(t testing.TB) oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/form-factor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FormFactorPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FormFactorPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/form-factor with a ONCE subscription.
func (n *Component_Transceiver_FormFactorPathAny) Get(t testing.TB) []oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/form-factor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FormFactorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_FormFactorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool) *oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_FormFactorPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/form-factor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FormFactorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool) *oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_FormFactorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/form-factor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_FormFactorPath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/form-factor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/form-factor to the batch object.
func (n *Component_Transceiver_FormFactorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/form-factor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FormFactorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/form-factor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FormFactorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool) *oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_FormFactorPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/form-factor to the batch object.
func (n *Component_Transceiver_FormFactorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_FormFactorPath extracts the value of the leaf FormFactor from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE.
func convertComponent_Transceiver_FormFactorPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE{
		Metadata: md,
	}
	val := parent.FormFactor
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/form-factor-preconf with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FormFactorPreconfPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_FormFactorPreconfPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/form-factor-preconf with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FormFactorPreconfPath) Get(t testing.TB) oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/form-factor-preconf with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FormFactorPreconfPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FormFactorPreconfPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/form-factor-preconf with a ONCE subscription.
func (n *Component_Transceiver_FormFactorPreconfPathAny) Get(t testing.TB) []oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/form-factor-preconf with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FormFactorPreconfPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_FormFactorPreconfPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool) *oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, true, false)
		return convertComponent_Transceiver_FormFactorPreconfPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/form-factor-preconf with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FormFactorPreconfPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool) *oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_FormFactorPreconfPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/form-factor-preconf with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_FormFactorPreconfPath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/form-factor-preconf failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/form-factor-preconf to the batch object.
func (n *Component_Transceiver_FormFactorPreconfPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/form-factor-preconf with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_FormFactorPreconfPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/form-factor-preconf with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_FormFactorPreconfPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) bool) *oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher {
	t.Helper()
	return watch_Component_Transceiver_FormFactorPreconfPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/form-factor-preconf to the batch object.
func (n *Component_Transceiver_FormFactorPreconfPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_FormFactorPreconfPath extracts the value of the leaf FormFactorPreconf from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE.
func convertComponent_Transceiver_FormFactorPreconfPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE{
		Metadata: md,
	}
	val := parent.FormFactorPreconf
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/input-power with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_InputPowerPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver_InputPower {
	t.Helper()
	goStruct := &oc.Component_Transceiver_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_InputPower", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Transceiver_InputPower{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/input-power with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_InputPowerPath) Get(t testing.TB) *oc.Component_Transceiver_InputPower {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/input-power with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_InputPowerPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver_InputPower {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver_InputPower
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_InputPower{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_InputPower", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver_InputPower{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/input-power with a ONCE subscription.
func (n *Component_Transceiver_InputPowerPathAny) Get(t testing.TB) []*oc.Component_Transceiver_InputPower {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver_InputPower
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPowerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_InputPower {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_InputPower{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_InputPower) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Transceiver_InputPower{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Transceiver_InputPower)))
		return false
	})
	return c
}

func watch_Component_Transceiver_InputPowerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_InputPower) bool) *oc.Component_Transceiver_InputPowerWatcher {
	t.Helper()
	w := &oc.Component_Transceiver_InputPowerWatcher{}
	gs := &oc.Component_Transceiver_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_InputPower", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Transceiver_InputPower{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver_InputPower)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPowerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_InputPower) bool) *oc.Component_Transceiver_InputPowerWatcher {
	t.Helper()
	return watch_Component_Transceiver_InputPowerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/input-power with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_InputPowerPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Transceiver_InputPower) *oc.QualifiedComponent_Transceiver_InputPower {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Transceiver_InputPower) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/input-power failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power to the batch object.
func (n *Component_Transceiver_InputPowerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPowerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver_InputPower {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver_InputPower{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver_InputPower) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPowerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_InputPower) bool) *oc.Component_Transceiver_InputPowerWatcher {
	t.Helper()
	return watch_Component_Transceiver_InputPowerPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power to the batch object.
func (n *Component_Transceiver_InputPowerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_InputPower_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_InputPower_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/avg with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_InputPower_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_InputPower_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
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
		qv := convertComponent_Transceiver_InputPower_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/avg with a ONCE subscription.
func (n *Component_Transceiver_InputPower_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_InputPower_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_InputPower_AvgPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/input-power/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_InputPower_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/input-power/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/avg to the batch object.
func (n *Component_Transceiver_InputPower_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_AvgPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/avg to the batch object.
func (n *Component_Transceiver_InputPower_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_InputPower_AvgPath extracts the value of the leaf Avg from its parent oc.Component_Transceiver_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_InputPower_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_InputPower) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_InputPower_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_InputPower_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/instant with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_InputPower_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_InputPower_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
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
		qv := convertComponent_Transceiver_InputPower_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/instant with a ONCE subscription.
func (n *Component_Transceiver_InputPower_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_InputPower_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_InputPower_InstantPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/input-power/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_InputPower_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/input-power/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/instant to the batch object.
func (n *Component_Transceiver_InputPower_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_InstantPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/instant to the batch object.
func (n *Component_Transceiver_InputPower_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_InputPower_InstantPath extracts the value of the leaf Instant from its parent oc.Component_Transceiver_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_InputPower_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_InputPower) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_InputPower_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_InputPower_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_InputPower_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_InputPower_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Transceiver_InputPower_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/interval with a ONCE subscription.
func (n *Component_Transceiver_InputPower_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_InputPower_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_InputPower_IntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/input-power/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_InputPower_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/input-power/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/interval to the batch object.
func (n *Component_Transceiver_InputPower_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_IntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/interval to the batch object.
func (n *Component_Transceiver_InputPower_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_InputPower_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_Transceiver_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_InputPower_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_InputPower) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_InputPower_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_InputPower_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/max with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_InputPower_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_InputPower_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
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
		qv := convertComponent_Transceiver_InputPower_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/max with a ONCE subscription.
func (n *Component_Transceiver_InputPower_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_InputPower_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Transceiver_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_InputPower_MaxPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/input-power/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_InputPower_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/input-power/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/max to the batch object.
func (n *Component_Transceiver_InputPower_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_MaxPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/max to the batch object.
func (n *Component_Transceiver_InputPower_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_InputPower_MaxPath extracts the value of the leaf Max from its parent oc.Component_Transceiver_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_InputPower_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_InputPower) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_InputPower_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_InputPower{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_InputPower", goStruct, true, false)
	if ok {
		return convertComponent_Transceiver_InputPower_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/state/input-power/max-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_InputPower_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_InputPower_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertComponent_Transceiver_InputPower_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/state/input-power/max-time with a ONCE subscription.
func (n *Component_Transceiver_InputPower_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Transceiver_InputPower_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Transceiver_InputPower{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_InputPower_MaxTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver/state/input-power/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Transceiver_InputPower_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver/state/input-power/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/max-time to the batch object.
func (n *Component_Transceiver_InputPower_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver/state/input-power/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Transceiver_InputPower_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_MaxTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver/state/input-power/max-time to the batch object.
func (n *Component_Transceiver_InputPower_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Transceiver_InputPower_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_Transceiver_InputPower
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Transceiver_InputPower_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_InputPower) *oc.QualifiedUint64 {
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_InputPower_MinPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_MinPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_InputPower", gs, queryPath, true, false)
		return convertComponent_Transceiver_InputPower_MinTimePath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/input-power/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_InputPower_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_InputPower_MinTimePath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Transceiver_LaserBiasCurrent{
			Metadata: md,
		}).SetVal(gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver_LaserBiasCurrent) bool) *oc.Component_Transceiver_LaserBiasCurrentWatcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrentPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return convertComponent_Transceiver_LaserBiasCurrent_AvgPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_AvgPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return convertComponent_Transceiver_LaserBiasCurrent_InstantPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_InstantPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return convertComponent_Transceiver_LaserBiasCurrent_IntervalPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_IntervalPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return convertComponent_Transceiver_LaserBiasCurrent_MaxPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MaxPath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return convertComponent_Transceiver_LaserBiasCurrent_MaxTimePath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MaxTimePath(t, n, timeout, predicate)
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
// failing the test fatally is no value is present at the path.
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
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver_LaserBiasCurrent", gs, queryPath, true, false)
		return convertComponent_Transceiver_LaserBiasCurrent_MinPath(t, md, gs), nil
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver/state/laser-bias-current/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Transceiver_LaserBiasCurrent_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Transceiver_LaserBiasCurrent_MinPath(t, n, timeout, predicate)
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
