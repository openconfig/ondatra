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

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, true, false)
	if ok {
		return convertComponent_Port_BreakoutMode_Group_NumPhysicalChannelsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Port_BreakoutMode_Group_NumPhysicalChannelsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels with a ONCE subscription.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Component_Port_BreakoutMode_Group{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Port_BreakoutMode_Group", gs, queryPath, true, false)
		return convertComponent_Port_BreakoutMode_Group_NumPhysicalChannelsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels to the batch object.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/port/breakout-mode/groups/group/state/num-physical-channels to the batch object.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Port_BreakoutMode_Group_NumPhysicalChannelsPath extracts the value of the leaf NumPhysicalChannels from its parent oc.Component_Port_BreakoutMode_Group
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertComponent_Port_BreakoutMode_Group_NumPhysicalChannelsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.NumPhysicalChannels
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/power-supply with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_PowerSupplyPath) Lookup(t testing.TB) *oc.QualifiedComponent_PowerSupply {
	t.Helper()
	goStruct := &oc.Component_PowerSupply{}
	md, ok := oc.Lookup(t, n, "Component_PowerSupply", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_PowerSupply{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/power-supply with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_PowerSupplyPath) Get(t testing.TB) *oc.Component_PowerSupply {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/power-supply with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_PowerSupplyPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_PowerSupply {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_PowerSupply
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_PowerSupply{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_PowerSupply", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_PowerSupply{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/power-supply with a ONCE subscription.
func (n *Component_PowerSupplyPathAny) Get(t testing.TB) []*oc.Component_PowerSupply {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_PowerSupply
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/power-supply with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_PowerSupplyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_PowerSupply {
	t.Helper()
	c := &oc.CollectionComponent_PowerSupply{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_PowerSupply) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_PowerSupply{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_PowerSupply)))
		return false
	})
	return c
}

func watch_Component_PowerSupplyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_PowerSupply) bool) *oc.Component_PowerSupplyWatcher {
	t.Helper()
	w := &oc.Component_PowerSupplyWatcher{}
	gs := &oc.Component_PowerSupply{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_PowerSupply", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_PowerSupply{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_PowerSupply)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/power-supply with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_PowerSupplyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_PowerSupply) bool) *oc.Component_PowerSupplyWatcher {
	t.Helper()
	return watch_Component_PowerSupplyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/power-supply with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_PowerSupplyPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_PowerSupply) *oc.QualifiedComponent_PowerSupply {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_PowerSupply) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/power-supply failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/power-supply to the batch object.
func (n *Component_PowerSupplyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/power-supply with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_PowerSupplyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_PowerSupply {
	t.Helper()
	c := &oc.CollectionComponent_PowerSupply{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_PowerSupply) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/power-supply with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_PowerSupplyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_PowerSupply) bool) *oc.Component_PowerSupplyWatcher {
	t.Helper()
	return watch_Component_PowerSupplyPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/power-supply to the batch object.
func (n *Component_PowerSupplyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/properties/property with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_PropertyPath) Lookup(t testing.TB) *oc.QualifiedComponent_Property {
	t.Helper()
	goStruct := &oc.Component_Property{}
	md, ok := oc.Lookup(t, n, "Component_Property", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Property{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/properties/property with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_PropertyPath) Get(t testing.TB) *oc.Component_Property {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/properties/property with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_PropertyPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Property {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Property
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Property{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Property", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Property{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/properties/property with a ONCE subscription.
func (n *Component_PropertyPathAny) Get(t testing.TB) []*oc.Component_Property {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Property
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/properties/property with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_PropertyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Property {
	t.Helper()
	c := &oc.CollectionComponent_Property{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Property) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Property{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Property)))
		return false
	})
	return c
}

func watch_Component_PropertyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Property) bool) *oc.Component_PropertyWatcher {
	t.Helper()
	w := &oc.Component_PropertyWatcher{}
	gs := &oc.Component_Property{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Property", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Property{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Property)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/properties/property with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_PropertyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Property) bool) *oc.Component_PropertyWatcher {
	t.Helper()
	return watch_Component_PropertyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/properties/property with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_PropertyPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Property) *oc.QualifiedComponent_Property {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Property) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/properties/property failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/properties/property to the batch object.
func (n *Component_PropertyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/properties/property with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_PropertyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Property {
	t.Helper()
	c := &oc.CollectionComponent_Property{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Property) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/properties/property with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_PropertyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Property) bool) *oc.Component_PropertyWatcher {
	t.Helper()
	return watch_Component_PropertyPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/properties/property to the batch object.
func (n *Component_PropertyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/properties/property/state/configurable with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Property_ConfigurablePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Component_Property{}
	md, ok := oc.Lookup(t, n, "Component_Property", goStruct, true, false)
	if ok {
		return convertComponent_Property_ConfigurablePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/properties/property/state/configurable with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Property_ConfigurablePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/properties/property/state/configurable with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Property_ConfigurablePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Property{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Property", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Property_ConfigurablePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/properties/property/state/configurable with a ONCE subscription.
func (n *Component_Property_ConfigurablePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/properties/property/state/configurable with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Property_ConfigurablePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Property_ConfigurablePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Component_Property{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Property", gs, queryPath, true, false)
		return convertComponent_Property_ConfigurablePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/properties/property/state/configurable with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Property_ConfigurablePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Component_Property_ConfigurablePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/properties/property/state/configurable with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Property_ConfigurablePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/properties/property/state/configurable failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/properties/property/state/configurable to the batch object.
func (n *Component_Property_ConfigurablePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/properties/property/state/configurable with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Property_ConfigurablePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/properties/property/state/configurable with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Property_ConfigurablePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Component_Property_ConfigurablePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/properties/property/state/configurable to the batch object.
func (n *Component_Property_ConfigurablePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Property_ConfigurablePath extracts the value of the leaf Configurable from its parent oc.Component_Property
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertComponent_Property_ConfigurablePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Property) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Configurable
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/properties/property/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Property_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Property{}
	md, ok := oc.Lookup(t, n, "Component_Property", goStruct, true, false)
	if ok {
		return convertComponent_Property_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/properties/property/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Property_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/properties/property/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Property_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Property{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Property", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Property_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/properties/property/state/name with a ONCE subscription.
func (n *Component_Property_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/properties/property/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Property_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Property_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_Property{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Property", gs, queryPath, true, false)
		return convertComponent_Property_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/properties/property/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Property_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Property_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/properties/property/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Property_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/properties/property/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/properties/property/state/name to the batch object.
func (n *Component_Property_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/properties/property/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Property_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/properties/property/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Property_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Property_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/properties/property/state/name to the batch object.
func (n *Component_Property_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Property_NamePath extracts the value of the leaf Name from its parent oc.Component_Property
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Property_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Property) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-platform/components/component/properties/property/state/value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Property_ValuePath) Lookup(t testing.TB) *oc.QualifiedComponent_Property_Value_Union {
	t.Helper()
	goStruct := &oc.Component_Property{}
	md, ok := oc.Lookup(t, n, "Component_Property", goStruct, true, false)
	if ok {
		return convertComponent_Property_ValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/properties/property/state/value with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Property_ValuePath) Get(t testing.TB) oc.Component_Property_Value_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/properties/property/state/value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Property_ValuePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Property_Value_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Property_Value_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Property{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Property", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Property_ValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/properties/property/state/value with a ONCE subscription.
func (n *Component_Property_ValuePathAny) Get(t testing.TB) []oc.Component_Property_Value_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Component_Property_Value_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/properties/property/state/value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Property_ValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Property_Value_Union {
	t.Helper()
	c := &oc.CollectionComponent_Property_Value_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Property_Value_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Property_ValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Property_Value_Union) bool) *oc.Component_Property_Value_UnionWatcher {
	t.Helper()
	w := &oc.Component_Property_Value_UnionWatcher{}
	gs := &oc.Component_Property{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Property", gs, queryPath, true, false)
		return convertComponent_Property_ValuePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Property_Value_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/properties/property/state/value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Property_ValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Property_Value_Union) bool) *oc.Component_Property_Value_UnionWatcher {
	t.Helper()
	return watch_Component_Property_ValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/properties/property/state/value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Property_ValuePath) Await(t testing.TB, timeout time.Duration, val oc.Component_Property_Value_Union) *oc.QualifiedComponent_Property_Value_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Property_Value_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/properties/property/state/value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/properties/property/state/value to the batch object.
func (n *Component_Property_ValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/properties/property/state/value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Property_ValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Property_Value_Union {
	t.Helper()
	c := &oc.CollectionComponent_Property_Value_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Property_Value_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/properties/property/state/value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Property_ValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Property_Value_Union) bool) *oc.Component_Property_Value_UnionWatcher {
	t.Helper()
	return watch_Component_Property_ValuePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/properties/property/state/value to the batch object.
func (n *Component_Property_ValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Property_ValuePath extracts the value of the leaf Value from its parent oc.Component_Property
// and combines the update with an existing Metadata to return a *oc.QualifiedComponent_Property_Value_Union.
func convertComponent_Property_ValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Property) *oc.QualifiedComponent_Property_Value_Union {
	t.Helper()
	qv := &oc.QualifiedComponent_Property_Value_Union{
		Metadata: md,
	}
	val := parent.Value
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/removable with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_RemovablePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_RemovablePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/removable with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_RemovablePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/removable with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_RemovablePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_RemovablePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/removable with a ONCE subscription.
func (n *Component_RemovablePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/removable with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_RemovablePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_RemovablePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return convertComponent_RemovablePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/removable with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_RemovablePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Component_RemovablePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/removable with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_RemovablePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/removable failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/removable to the batch object.
func (n *Component_RemovablePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/removable with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_RemovablePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/removable with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_RemovablePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Component_RemovablePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/removable to the batch object.
func (n *Component_RemovablePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_RemovablePath extracts the value of the leaf Removable from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertComponent_RemovablePath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Removable
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/serial-no with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_SerialNoPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_SerialNoPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/serial-no with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_SerialNoPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/serial-no with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_SerialNoPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertComponent_SerialNoPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/serial-no with a ONCE subscription.
func (n *Component_SerialNoPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/serial-no with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_SerialNoPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_SerialNoPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return convertComponent_SerialNoPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/serial-no with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_SerialNoPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_SerialNoPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/serial-no with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_SerialNoPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/serial-no failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/serial-no to the batch object.
func (n *Component_SerialNoPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/serial-no with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_SerialNoPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/serial-no with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_SerialNoPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_SerialNoPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/serial-no to the batch object.
func (n *Component_SerialNoPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_SerialNoPath extracts the value of the leaf SerialNo from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_SerialNoPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-platform/components/component/software-module with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_SoftwareModulePath) Lookup(t testing.TB) *oc.QualifiedComponent_SoftwareModule {
	t.Helper()
	goStruct := &oc.Component_SoftwareModule{}
	md, ok := oc.Lookup(t, n, "Component_SoftwareModule", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_SoftwareModule{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/software-module with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_SoftwareModulePath) Get(t testing.TB) *oc.Component_SoftwareModule {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/software-module with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_SoftwareModulePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_SoftwareModule {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_SoftwareModule
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_SoftwareModule{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_SoftwareModule", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_SoftwareModule{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/software-module with a ONCE subscription.
func (n *Component_SoftwareModulePathAny) Get(t testing.TB) []*oc.Component_SoftwareModule {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_SoftwareModule
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/software-module with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_SoftwareModulePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_SoftwareModule {
	t.Helper()
	c := &oc.CollectionComponent_SoftwareModule{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_SoftwareModule) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_SoftwareModule{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_SoftwareModule)))
		return false
	})
	return c
}

func watch_Component_SoftwareModulePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_SoftwareModule) bool) *oc.Component_SoftwareModuleWatcher {
	t.Helper()
	w := &oc.Component_SoftwareModuleWatcher{}
	gs := &oc.Component_SoftwareModule{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_SoftwareModule", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_SoftwareModule{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_SoftwareModule)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/software-module with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_SoftwareModulePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_SoftwareModule) bool) *oc.Component_SoftwareModuleWatcher {
	t.Helper()
	return watch_Component_SoftwareModulePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/software-module with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_SoftwareModulePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_SoftwareModule) *oc.QualifiedComponent_SoftwareModule {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_SoftwareModule) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/software-module failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/software-module to the batch object.
func (n *Component_SoftwareModulePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/software-module with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_SoftwareModulePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_SoftwareModule {
	t.Helper()
	c := &oc.CollectionComponent_SoftwareModule{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_SoftwareModule) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/software-module with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_SoftwareModulePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_SoftwareModule) bool) *oc.Component_SoftwareModuleWatcher {
	t.Helper()
	return watch_Component_SoftwareModulePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/software-module to the batch object.
func (n *Component_SoftwareModulePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/software-module/state/module-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_SoftwareModule_ModuleTypePath) Lookup(t testing.TB) *oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	t.Helper()
	goStruct := &oc.Component_SoftwareModule{}
	md, ok := oc.Lookup(t, n, "Component_SoftwareModule", goStruct, true, false)
	if ok {
		return convertComponent_SoftwareModule_ModuleTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/software-module/state/module-type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_SoftwareModule_ModuleTypePath) Get(t testing.TB) oc.E_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/software-module/state/module-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_SoftwareModule_ModuleTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_SoftwareModule{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_SoftwareModule", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_SoftwareModule_ModuleTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/software-module/state/module-type with a ONCE subscription.
func (n *Component_SoftwareModule_ModuleTypePathAny) Get(t testing.TB) []oc.E_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PlatformSoftware_SOFTWARE_MODULE_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/software-module/state/module-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_SoftwareModule_ModuleTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	t.Helper()
	c := &oc.CollectionE_PlatformSoftware_SOFTWARE_MODULE_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_SoftwareModule_ModuleTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE) bool) *oc.E_PlatformSoftware_SOFTWARE_MODULE_TYPEWatcher {
	t.Helper()
	w := &oc.E_PlatformSoftware_SOFTWARE_MODULE_TYPEWatcher{}
	gs := &oc.Component_SoftwareModule{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_SoftwareModule", gs, queryPath, true, false)
		return convertComponent_SoftwareModule_ModuleTypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/software-module/state/module-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_SoftwareModule_ModuleTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE) bool) *oc.E_PlatformSoftware_SOFTWARE_MODULE_TYPEWatcher {
	t.Helper()
	return watch_Component_SoftwareModule_ModuleTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/software-module/state/module-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_SoftwareModule_ModuleTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_PlatformSoftware_SOFTWARE_MODULE_TYPE) *oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/software-module/state/module-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/software-module/state/module-type to the batch object.
func (n *Component_SoftwareModule_ModuleTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/software-module/state/module-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_SoftwareModule_ModuleTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	t.Helper()
	c := &oc.CollectionE_PlatformSoftware_SOFTWARE_MODULE_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/software-module/state/module-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_SoftwareModule_ModuleTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE) bool) *oc.E_PlatformSoftware_SOFTWARE_MODULE_TYPEWatcher {
	t.Helper()
	return watch_Component_SoftwareModule_ModuleTypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/software-module/state/module-type to the batch object.
func (n *Component_SoftwareModule_ModuleTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_SoftwareModule_ModuleTypePath extracts the value of the leaf ModuleType from its parent oc.Component_SoftwareModule
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE.
func convertComponent_SoftwareModule_ModuleTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_SoftwareModule) *oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE{
		Metadata: md,
	}
	val := parent.ModuleType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/software-version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_SoftwareVersionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_SoftwareVersionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/software-version with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_SoftwareVersionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/software-version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_SoftwareVersionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertComponent_SoftwareVersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/software-version with a ONCE subscription.
func (n *Component_SoftwareVersionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/software-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_SoftwareVersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_SoftwareVersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return convertComponent_SoftwareVersionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/software-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_SoftwareVersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_SoftwareVersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/software-version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_SoftwareVersionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/software-version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/software-version to the batch object.
func (n *Component_SoftwareVersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/software-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_SoftwareVersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/software-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_SoftwareVersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_SoftwareVersionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/software-version to the batch object.
func (n *Component_SoftwareVersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_SoftwareVersionPath extracts the value of the leaf SoftwareVersion from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_SoftwareVersionPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SoftwareVersion
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
