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

// Lookup fetches the value at /openconfig-platform/components/component/state/parent with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_ParentPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_ParentPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/parent with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_ParentPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/parent with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_ParentPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertComponent_ParentPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/parent with a ONCE subscription.
func (n *Component_ParentPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/parent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_ParentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_ParentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_ParentPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/parent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_ParentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_ParentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/parent with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_ParentPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/parent failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/parent to the batch object.
func (n *Component_ParentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/parent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_ParentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_ParentPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertComponent_ParentPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/parent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_ParentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_ParentPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/parent to the batch object.
func (n *Component_ParentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_ParentPath extracts the value of the leaf Parent from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_ParentPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Parent
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/part-no with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_PartNoPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_PartNoPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/part-no with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_PartNoPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/part-no with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_PartNoPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertComponent_PartNoPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/part-no with a ONCE subscription.
func (n *Component_PartNoPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/part-no with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_PartNoPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_PartNoPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_PartNoPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/part-no with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_PartNoPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_PartNoPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/part-no with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_PartNoPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/part-no failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/part-no to the batch object.
func (n *Component_PartNoPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/part-no with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_PartNoPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_PartNoPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertComponent_PartNoPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/part-no with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_PartNoPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_PartNoPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/part-no to the batch object.
func (n *Component_PartNoPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_PartNoPath extracts the value of the leaf PartNo from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_PartNoPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.PartNo
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_PciePath) Lookup(t testing.TB) *oc.QualifiedComponent_Pcie {
	t.Helper()
	goStruct := &oc.Component_Pcie{}
	md, ok := oc.Lookup(t, n, "Component_Pcie", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Pcie{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_PciePath) Get(t testing.TB) *oc.Component_Pcie {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_PciePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Pcie {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Pcie
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Pcie{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie with a ONCE subscription.
func (n *Component_PciePathAny) Get(t testing.TB) []*oc.Component_Pcie {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Pcie
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_PciePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Pcie {
	t.Helper()
	c := &oc.CollectionComponent_Pcie{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Pcie) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Pcie{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Pcie)))
		return false
	})
	return c
}

func watch_Component_PciePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Pcie) bool) *oc.Component_PcieWatcher {
	t.Helper()
	w := &oc.Component_PcieWatcher{}
	gs := &oc.Component_Pcie{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_Pcie{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Pcie)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_PciePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Pcie) bool) *oc.Component_PcieWatcher {
	t.Helper()
	return watch_Component_PciePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_PciePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Pcie) *oc.QualifiedComponent_Pcie {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Pcie) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie to the batch object.
func (n *Component_PciePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_PciePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Pcie {
	t.Helper()
	c := &oc.CollectionComponent_Pcie{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Pcie) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_PciePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Pcie) bool) *oc.Component_PcieWatcher {
	t.Helper()
	w := &oc.Component_PcieWatcher{}
	structs := map[string]*oc.Component_Pcie{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_Pcie{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Pcie)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_PciePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Pcie) bool) *oc.Component_PcieWatcher {
	t.Helper()
	return watch_Component_PciePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie to the batch object.
func (n *Component_PciePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_CorrectableErrorsPath) Lookup(t testing.TB) *oc.QualifiedComponent_Pcie_CorrectableErrors {
	t.Helper()
	goStruct := &oc.Component_Pcie_CorrectableErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_CorrectableErrors", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Pcie_CorrectableErrors{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_CorrectableErrorsPath) Get(t testing.TB) *oc.Component_Pcie_CorrectableErrors {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_CorrectableErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Pcie_CorrectableErrors {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Pcie_CorrectableErrors
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_CorrectableErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_CorrectableErrors", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Pcie_CorrectableErrors{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors with a ONCE subscription.
func (n *Component_Pcie_CorrectableErrorsPathAny) Get(t testing.TB) []*oc.Component_Pcie_CorrectableErrors {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Pcie_CorrectableErrors
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Pcie_CorrectableErrors {
	t.Helper()
	c := &oc.CollectionComponent_Pcie_CorrectableErrors{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Pcie_CorrectableErrors) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Pcie_CorrectableErrors{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Pcie_CorrectableErrors)))
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_CorrectableErrors) bool) *oc.Component_Pcie_CorrectableErrorsWatcher {
	t.Helper()
	w := &oc.Component_Pcie_CorrectableErrorsWatcher{}
	gs := &oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_CorrectableErrors", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_Pcie_CorrectableErrors{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Pcie_CorrectableErrors)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_CorrectableErrors) bool) *oc.Component_Pcie_CorrectableErrorsWatcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/correctable-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_CorrectableErrorsPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Pcie_CorrectableErrors) *oc.QualifiedComponent_Pcie_CorrectableErrors {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Pcie_CorrectableErrors) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/correctable-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors to the batch object.
func (n *Component_Pcie_CorrectableErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Pcie_CorrectableErrors {
	t.Helper()
	c := &oc.CollectionComponent_Pcie_CorrectableErrors{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Pcie_CorrectableErrors) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_CorrectableErrors) bool) *oc.Component_Pcie_CorrectableErrorsWatcher {
	t.Helper()
	w := &oc.Component_Pcie_CorrectableErrorsWatcher{}
	structs := map[string]*oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_CorrectableErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_CorrectableErrors", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_Pcie_CorrectableErrors{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Pcie_CorrectableErrors)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_CorrectableErrors) bool) *oc.Component_Pcie_CorrectableErrorsWatcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors to the batch object.
func (n *Component_Pcie_CorrectableErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_CorrectableErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_CorrectableErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_CorrectableErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_CorrectableErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors with a ONCE subscription.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_CorrectableErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_CorrectableErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_CorrectableErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/advisory-non-fatal-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath extracts the value of the leaf AdvisoryNonFatalErrors from its parent oc.Component_Pcie_CorrectableErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_CorrectableErrors_AdvisoryNonFatalErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_CorrectableErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AdvisoryNonFatalErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_CorrectableErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_CorrectableErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_CorrectableErrors_BadDllpErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_CorrectableErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_CorrectableErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_CorrectableErrors_BadDllpErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors with a ONCE subscription.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_BadDllpErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_CorrectableErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_CorrectableErrors_BadDllpErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_BadDllpErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_BadDllpErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_CorrectableErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_CorrectableErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_CorrectableErrors_BadDllpErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_BadDllpErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/bad-dllp-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_BadDllpErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_CorrectableErrors_BadDllpErrorsPath extracts the value of the leaf BadDllpErrors from its parent oc.Component_Pcie_CorrectableErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_CorrectableErrors_BadDllpErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_CorrectableErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.BadDllpErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_CorrectableErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_CorrectableErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_CorrectableErrors_BadTlpErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_CorrectableErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_CorrectableErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_CorrectableErrors_BadTlpErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors with a ONCE subscription.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_BadTlpErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_CorrectableErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_CorrectableErrors_BadTlpErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_BadTlpErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_BadTlpErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_CorrectableErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_CorrectableErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_CorrectableErrors_BadTlpErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_BadTlpErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/bad-tlp-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_BadTlpErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_CorrectableErrors_BadTlpErrorsPath extracts the value of the leaf BadTlpErrors from its parent oc.Component_Pcie_CorrectableErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_CorrectableErrors_BadTlpErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_CorrectableErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.BadTlpErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_CorrectableErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_CorrectableErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_CorrectableErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_CorrectableErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors with a ONCE subscription.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_CorrectableErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_CorrectableErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_CorrectableErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/hdr-log-overflow-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_HdrLogOverflowErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath extracts the value of the leaf HdrLogOverflowErrors from its parent oc.Component_Pcie_CorrectableErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_CorrectableErrors_HdrLogOverflowErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_CorrectableErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.HdrLogOverflowErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_CorrectableErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_CorrectableErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_CorrectableErrors_InternalErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_CorrectableErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_CorrectableErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_CorrectableErrors_InternalErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors with a ONCE subscription.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_InternalErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_CorrectableErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_CorrectableErrors_InternalErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_InternalErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_InternalErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_CorrectableErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_CorrectableErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_CorrectableErrors_InternalErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_InternalErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/internal-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_InternalErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_CorrectableErrors_InternalErrorsPath extracts the value of the leaf InternalErrors from its parent oc.Component_Pcie_CorrectableErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_CorrectableErrors_InternalErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_CorrectableErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InternalErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_CorrectableErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_CorrectableErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_CorrectableErrors_ReceiverErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_CorrectableErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_CorrectableErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_CorrectableErrors_ReceiverErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors with a ONCE subscription.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_ReceiverErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_CorrectableErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_CorrectableErrors_ReceiverErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_ReceiverErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_ReceiverErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_CorrectableErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_CorrectableErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_CorrectableErrors_ReceiverErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_ReceiverErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/receiver-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_ReceiverErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_CorrectableErrors_ReceiverErrorsPath extracts the value of the leaf ReceiverErrors from its parent oc.Component_Pcie_CorrectableErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_CorrectableErrors_ReceiverErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_CorrectableErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ReceiverErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_CorrectableErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_CorrectableErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_CorrectableErrors_RelayRolloverErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_CorrectableErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_CorrectableErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_CorrectableErrors_RelayRolloverErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors with a ONCE subscription.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_RelayRolloverErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_CorrectableErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_CorrectableErrors_RelayRolloverErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_RelayRolloverErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_RelayRolloverErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_CorrectableErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_CorrectableErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_CorrectableErrors_RelayRolloverErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_RelayRolloverErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/relay-rollover-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_RelayRolloverErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_CorrectableErrors_RelayRolloverErrorsPath extracts the value of the leaf RelayRolloverErrors from its parent oc.Component_Pcie_CorrectableErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_CorrectableErrors_RelayRolloverErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_CorrectableErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.RelayRolloverErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_CorrectableErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_CorrectableErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_CorrectableErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_CorrectableErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors with a ONCE subscription.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_CorrectableErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_CorrectableErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_CorrectableErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/replay-timeout-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_ReplayTimeoutErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath extracts the value of the leaf ReplayTimeoutErrors from its parent oc.Component_Pcie_CorrectableErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_CorrectableErrors_ReplayTimeoutErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_CorrectableErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ReplayTimeoutErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_CorrectableErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_CorrectableErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_CorrectableErrors_TotalErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_CorrectableErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_CorrectableErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_CorrectableErrors_TotalErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors with a ONCE subscription.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_TotalErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_CorrectableErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_CorrectableErrors_TotalErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_TotalErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_CorrectableErrors_TotalErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_CorrectableErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_CorrectableErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_CorrectableErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_CorrectableErrors_TotalErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_CorrectableErrors_TotalErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/correctable-errors/total-errors to the batch object.
func (n *Component_Pcie_CorrectableErrors_TotalErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_CorrectableErrors_TotalErrorsPath extracts the value of the leaf TotalErrors from its parent oc.Component_Pcie_CorrectableErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_CorrectableErrors_TotalErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_CorrectableErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TotalErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrorsPath) Lookup(t testing.TB) *oc.QualifiedComponent_Pcie_FatalErrors {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Pcie_FatalErrors{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrorsPath) Get(t testing.TB) *oc.Component_Pcie_FatalErrors {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Pcie_FatalErrors {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Pcie_FatalErrors
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Pcie_FatalErrors{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrorsPathAny) Get(t testing.TB) []*oc.Component_Pcie_FatalErrors {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Pcie_FatalErrors
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Pcie_FatalErrors {
	t.Helper()
	c := &oc.CollectionComponent_Pcie_FatalErrors{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Pcie_FatalErrors) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Pcie_FatalErrors{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Pcie_FatalErrors)))
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_FatalErrors) bool) *oc.Component_Pcie_FatalErrorsWatcher {
	t.Helper()
	w := &oc.Component_Pcie_FatalErrorsWatcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, false, false)
		qv := (&oc.QualifiedComponent_Pcie_FatalErrors{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Pcie_FatalErrors)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_FatalErrors) bool) *oc.Component_Pcie_FatalErrorsWatcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrorsPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Pcie_FatalErrors) *oc.QualifiedComponent_Pcie_FatalErrors {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Pcie_FatalErrors) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors to the batch object.
func (n *Component_Pcie_FatalErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Pcie_FatalErrors {
	t.Helper()
	c := &oc.CollectionComponent_Pcie_FatalErrors{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Pcie_FatalErrors) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_FatalErrors) bool) *oc.Component_Pcie_FatalErrorsWatcher {
	t.Helper()
	w := &oc.Component_Pcie_FatalErrorsWatcher{}
	structs := map[string]*oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_FatalErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_FatalErrors", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedComponent_Pcie_FatalErrors{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Pcie_FatalErrors)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Pcie_FatalErrors) bool) *oc.Component_Pcie_FatalErrorsWatcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors to the batch object.
func (n *Component_Pcie_FatalErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_AcsViolationErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_AcsViolationErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors to the batch object.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_AcsViolationErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_FatalErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_FatalErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_AcsViolationErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/acs-violation-errors to the batch object.
func (n *Component_Pcie_FatalErrors_AcsViolationErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath extracts the value of the leaf AcsViolationErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_AcsViolationErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AcsViolationErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors to the batch object.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_FatalErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_FatalErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/atomic-op-blocked-errors to the batch object.
func (n *Component_Pcie_FatalErrors_AtomicOpBlockedErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath extracts the value of the leaf AtomicOpBlockedErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_AtomicOpBlockedErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AtomicOpBlockedErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_BlockedTlpErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_BlockedTlpErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors to the batch object.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_FatalErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_FatalErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/blocked-tlp-errors to the batch object.
func (n *Component_Pcie_FatalErrors_BlockedTlpErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath extracts the value of the leaf BlockedTlpErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_BlockedTlpErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.BlockedTlpErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Pcie_FatalErrors{}
	md, ok := oc.Lookup(t, n, "Component_Pcie_FatalErrors", goStruct, true, false)
	if ok {
		return convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Pcie_FatalErrors{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Pcie_FatalErrors", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a ONCE subscription.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_CompletionAbortErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Pcie_FatalErrors", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_CompletionAbortErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors to the batch object.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.Component_Pcie_FatalErrors{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Component_Pcie_FatalErrors{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Component_Pcie_FatalErrors", structs[pre], queryPath, true, false)
			qv := convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/pcie/fatal-errors/completion-abort-errors to the batch object.
func (n *Component_Pcie_FatalErrors_CompletionAbortErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath extracts the value of the leaf CompletionAbortErrors from its parent oc.Component_Pcie_FatalErrors
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Pcie_FatalErrors_CompletionAbortErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Pcie_FatalErrors) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CompletionAbortErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
