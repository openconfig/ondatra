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

// Lookup fetches the value at /openconfig-platform/components/component with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *ComponentPath) Lookup(t testing.TB) *oc.QualifiedComponent {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *ComponentPath) Get(t testing.TB) *oc.Component {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *ComponentPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component with a ONCE subscription.
func (n *ComponentPathAny) Get(t testing.TB) []*oc.Component {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *ComponentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent {
	t.Helper()
	c := &oc.CollectionComponent{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component)))
		return false
	})
	return c
}

func watch_ComponentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent) bool) *oc.ComponentWatcher {
	t.Helper()
	w := &oc.ComponentWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, false, false)
		return (&oc.QualifiedComponent{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *ComponentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent) bool) *oc.ComponentWatcher {
	t.Helper()
	return watch_ComponentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *ComponentPath) Await(t testing.TB, timeout time.Duration, val *oc.Component) *oc.QualifiedComponent {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component to the batch object.
func (n *ComponentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *ComponentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent {
	t.Helper()
	c := &oc.CollectionComponent{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *ComponentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent) bool) *oc.ComponentWatcher {
	t.Helper()
	return watch_ComponentPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component to the batch object.
func (n *ComponentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/state/allocated-power with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_AllocatedPowerPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_AllocatedPowerPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/allocated-power with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_AllocatedPowerPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/allocated-power with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_AllocatedPowerPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
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
		qv := convertComponent_AllocatedPowerPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/allocated-power with a ONCE subscription.
func (n *Component_AllocatedPowerPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/allocated-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_AllocatedPowerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_AllocatedPowerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return convertComponent_AllocatedPowerPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/allocated-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_AllocatedPowerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Component_AllocatedPowerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/allocated-power with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_AllocatedPowerPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/allocated-power failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/allocated-power to the batch object.
func (n *Component_AllocatedPowerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/allocated-power with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_AllocatedPowerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/allocated-power with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_AllocatedPowerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Component_AllocatedPowerPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/allocated-power to the batch object.
func (n *Component_AllocatedPowerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_AllocatedPowerPath extracts the value of the leaf AllocatedPower from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertComponent_AllocatedPowerPath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.AllocatedPower
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/backplane with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_BackplanePath) Lookup(t testing.TB) *oc.QualifiedComponent_Backplane {
	t.Helper()
	goStruct := &oc.Component_Backplane{}
	md, ok := oc.Lookup(t, n, "Component_Backplane", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Backplane{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/backplane with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_BackplanePath) Get(t testing.TB) *oc.Component_Backplane {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/backplane with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_BackplanePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Backplane {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Backplane
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Backplane{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Backplane", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Backplane{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/backplane with a ONCE subscription.
func (n *Component_BackplanePathAny) Get(t testing.TB) []*oc.Component_Backplane {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Backplane
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/backplane with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_BackplanePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Backplane {
	t.Helper()
	c := &oc.CollectionComponent_Backplane{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Backplane) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Backplane{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Backplane)))
		return false
	})
	return c
}

func watch_Component_BackplanePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Backplane) bool) *oc.Component_BackplaneWatcher {
	t.Helper()
	w := &oc.Component_BackplaneWatcher{}
	gs := &oc.Component_Backplane{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Backplane", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Backplane{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Backplane)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/backplane with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_BackplanePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Backplane) bool) *oc.Component_BackplaneWatcher {
	t.Helper()
	return watch_Component_BackplanePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/backplane with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_BackplanePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Backplane) *oc.QualifiedComponent_Backplane {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Backplane) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/backplane failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/backplane to the batch object.
func (n *Component_BackplanePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/backplane with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_BackplanePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Backplane {
	t.Helper()
	c := &oc.CollectionComponent_Backplane{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Backplane) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/backplane with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_BackplanePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Backplane) bool) *oc.Component_BackplaneWatcher {
	t.Helper()
	return watch_Component_BackplanePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/backplane to the batch object.
func (n *Component_BackplanePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/chassis with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_ChassisPath) Lookup(t testing.TB) *oc.QualifiedComponent_Chassis {
	t.Helper()
	goStruct := &oc.Component_Chassis{}
	md, ok := oc.Lookup(t, n, "Component_Chassis", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Chassis{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/chassis with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_ChassisPath) Get(t testing.TB) *oc.Component_Chassis {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/chassis with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_ChassisPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Chassis {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Chassis
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Chassis{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Chassis", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Chassis{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/chassis with a ONCE subscription.
func (n *Component_ChassisPathAny) Get(t testing.TB) []*oc.Component_Chassis {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Chassis
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/chassis with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_ChassisPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Chassis {
	t.Helper()
	c := &oc.CollectionComponent_Chassis{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Chassis) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Chassis{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Chassis)))
		return false
	})
	return c
}

func watch_Component_ChassisPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Chassis) bool) *oc.Component_ChassisWatcher {
	t.Helper()
	w := &oc.Component_ChassisWatcher{}
	gs := &oc.Component_Chassis{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Chassis", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Chassis{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Chassis)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/chassis with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_ChassisPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Chassis) bool) *oc.Component_ChassisWatcher {
	t.Helper()
	return watch_Component_ChassisPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/chassis with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_ChassisPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Chassis) *oc.QualifiedComponent_Chassis {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Chassis) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/chassis failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/chassis to the batch object.
func (n *Component_ChassisPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/chassis with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_ChassisPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Chassis {
	t.Helper()
	c := &oc.CollectionComponent_Chassis{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Chassis) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/chassis with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_ChassisPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Chassis) bool) *oc.Component_ChassisWatcher {
	t.Helper()
	return watch_Component_ChassisPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/chassis to the batch object.
func (n *Component_ChassisPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/cpu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_CpuPath) Lookup(t testing.TB) *oc.QualifiedComponent_Cpu {
	t.Helper()
	goStruct := &oc.Component_Cpu{}
	md, ok := oc.Lookup(t, n, "Component_Cpu", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Cpu{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/cpu with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_CpuPath) Get(t testing.TB) *oc.Component_Cpu {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/cpu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_CpuPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Cpu {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Cpu
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Cpu{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Cpu", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Cpu{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/cpu with a ONCE subscription.
func (n *Component_CpuPathAny) Get(t testing.TB) []*oc.Component_Cpu {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Cpu
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_CpuPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Cpu {
	t.Helper()
	c := &oc.CollectionComponent_Cpu{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Cpu) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Cpu{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Cpu)))
		return false
	})
	return c
}

func watch_Component_CpuPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Cpu) bool) *oc.Component_CpuWatcher {
	t.Helper()
	w := &oc.Component_CpuWatcher{}
	gs := &oc.Component_Cpu{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Cpu", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Cpu{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Cpu)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_CpuPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Cpu) bool) *oc.Component_CpuWatcher {
	t.Helper()
	return watch_Component_CpuPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/cpu with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_CpuPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Cpu) *oc.QualifiedComponent_Cpu {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Cpu) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/cpu failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/cpu to the batch object.
func (n *Component_CpuPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_CpuPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Cpu {
	t.Helper()
	c := &oc.CollectionComponent_Cpu{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Cpu) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_CpuPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Cpu) bool) *oc.Component_CpuWatcher {
	t.Helper()
	return watch_Component_CpuPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/cpu to the batch object.
func (n *Component_CpuPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/cpu/utilization with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Cpu_UtilizationPath) Lookup(t testing.TB) *oc.QualifiedComponent_Cpu_Utilization {
	t.Helper()
	goStruct := &oc.Component_Cpu_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_Cpu_Utilization", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Cpu_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/cpu/utilization with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Cpu_UtilizationPath) Get(t testing.TB) *oc.Component_Cpu_Utilization {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/cpu/utilization with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Cpu_UtilizationPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Cpu_Utilization {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Cpu_Utilization
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Cpu_Utilization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Cpu_Utilization", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Cpu_Utilization{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/cpu/utilization with a ONCE subscription.
func (n *Component_Cpu_UtilizationPathAny) Get(t testing.TB) []*oc.Component_Cpu_Utilization {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Cpu_Utilization
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_UtilizationPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Cpu_Utilization {
	t.Helper()
	c := &oc.CollectionComponent_Cpu_Utilization{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Cpu_Utilization) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Cpu_Utilization{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Cpu_Utilization)))
		return false
	})
	return c
}

func watch_Component_Cpu_UtilizationPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Cpu_Utilization) bool) *oc.Component_Cpu_UtilizationWatcher {
	t.Helper()
	w := &oc.Component_Cpu_UtilizationWatcher{}
	gs := &oc.Component_Cpu_Utilization{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Cpu_Utilization", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Cpu_Utilization{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Cpu_Utilization)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_UtilizationPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Cpu_Utilization) bool) *oc.Component_Cpu_UtilizationWatcher {
	t.Helper()
	return watch_Component_Cpu_UtilizationPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/cpu/utilization with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Cpu_UtilizationPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Cpu_Utilization) *oc.QualifiedComponent_Cpu_Utilization {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Cpu_Utilization) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/cpu/utilization failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/cpu/utilization to the batch object.
func (n *Component_Cpu_UtilizationPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_UtilizationPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Cpu_Utilization {
	t.Helper()
	c := &oc.CollectionComponent_Cpu_Utilization{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Cpu_Utilization) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_UtilizationPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Cpu_Utilization) bool) *oc.Component_Cpu_UtilizationWatcher {
	t.Helper()
	return watch_Component_Cpu_UtilizationPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/cpu/utilization to the batch object.
func (n *Component_Cpu_UtilizationPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/cpu/utilization/state/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Cpu_Utilization_AvgPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Component_Cpu_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_Cpu_Utilization", goStruct, true, false)
	if ok {
		return convertComponent_Cpu_Utilization_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/cpu/utilization/state/avg with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Cpu_Utilization_AvgPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/cpu/utilization/state/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Cpu_Utilization_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Cpu_Utilization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Cpu_Utilization", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Cpu_Utilization_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/cpu/utilization/state/avg with a ONCE subscription.
func (n *Component_Cpu_Utilization_AvgPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization/state/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_Utilization_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Cpu_Utilization_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Component_Cpu_Utilization{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Cpu_Utilization", gs, queryPath, true, false)
		return convertComponent_Cpu_Utilization_AvgPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization/state/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_Utilization_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Cpu_Utilization_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/cpu/utilization/state/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Cpu_Utilization_AvgPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/cpu/utilization/state/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/cpu/utilization/state/avg to the batch object.
func (n *Component_Cpu_Utilization_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization/state/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_Utilization_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization/state/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_Utilization_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Cpu_Utilization_AvgPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/cpu/utilization/state/avg to the batch object.
func (n *Component_Cpu_Utilization_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Cpu_Utilization_AvgPath extracts the value of the leaf Avg from its parent oc.Component_Cpu_Utilization
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertComponent_Cpu_Utilization_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Cpu_Utilization) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/cpu/utilization/state/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Cpu_Utilization_InstantPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Component_Cpu_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_Cpu_Utilization", goStruct, true, false)
	if ok {
		return convertComponent_Cpu_Utilization_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/cpu/utilization/state/instant with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Cpu_Utilization_InstantPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/cpu/utilization/state/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Cpu_Utilization_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Cpu_Utilization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Cpu_Utilization", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Cpu_Utilization_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/cpu/utilization/state/instant with a ONCE subscription.
func (n *Component_Cpu_Utilization_InstantPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization/state/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_Utilization_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Cpu_Utilization_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Component_Cpu_Utilization{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Cpu_Utilization", gs, queryPath, true, false)
		return convertComponent_Cpu_Utilization_InstantPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization/state/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_Utilization_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Cpu_Utilization_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/cpu/utilization/state/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Cpu_Utilization_InstantPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/cpu/utilization/state/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/cpu/utilization/state/instant to the batch object.
func (n *Component_Cpu_Utilization_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization/state/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_Utilization_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization/state/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_Utilization_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Cpu_Utilization_InstantPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/cpu/utilization/state/instant to the batch object.
func (n *Component_Cpu_Utilization_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Cpu_Utilization_InstantPath extracts the value of the leaf Instant from its parent oc.Component_Cpu_Utilization
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertComponent_Cpu_Utilization_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Cpu_Utilization) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Instant
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/cpu/utilization/state/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Cpu_Utilization_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Cpu_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_Cpu_Utilization", goStruct, true, false)
	if ok {
		return convertComponent_Cpu_Utilization_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/cpu/utilization/state/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Cpu_Utilization_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/cpu/utilization/state/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Cpu_Utilization_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Cpu_Utilization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Cpu_Utilization", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Cpu_Utilization_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/cpu/utilization/state/interval with a ONCE subscription.
func (n *Component_Cpu_Utilization_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization/state/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_Utilization_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Cpu_Utilization_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Cpu_Utilization{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Cpu_Utilization", gs, queryPath, true, false)
		return convertComponent_Cpu_Utilization_IntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization/state/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_Utilization_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Cpu_Utilization_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/cpu/utilization/state/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Cpu_Utilization_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/cpu/utilization/state/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/cpu/utilization/state/interval to the batch object.
func (n *Component_Cpu_Utilization_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization/state/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_Utilization_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization/state/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_Utilization_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Cpu_Utilization_IntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/cpu/utilization/state/interval to the batch object.
func (n *Component_Cpu_Utilization_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Cpu_Utilization_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_Cpu_Utilization
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Cpu_Utilization_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Cpu_Utilization) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/cpu/utilization/state/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Cpu_Utilization_MaxPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Component_Cpu_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_Cpu_Utilization", goStruct, true, false)
	if ok {
		return convertComponent_Cpu_Utilization_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/cpu/utilization/state/max with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Cpu_Utilization_MaxPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/cpu/utilization/state/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Cpu_Utilization_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Cpu_Utilization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Cpu_Utilization", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Cpu_Utilization_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/cpu/utilization/state/max with a ONCE subscription.
func (n *Component_Cpu_Utilization_MaxPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization/state/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_Utilization_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Cpu_Utilization_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Component_Cpu_Utilization{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Cpu_Utilization", gs, queryPath, true, false)
		return convertComponent_Cpu_Utilization_MaxPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization/state/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_Utilization_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Cpu_Utilization_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/cpu/utilization/state/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Cpu_Utilization_MaxPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/cpu/utilization/state/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/cpu/utilization/state/max to the batch object.
func (n *Component_Cpu_Utilization_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization/state/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_Utilization_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization/state/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_Utilization_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Component_Cpu_Utilization_MaxPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/cpu/utilization/state/max to the batch object.
func (n *Component_Cpu_Utilization_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Cpu_Utilization_MaxPath extracts the value of the leaf Max from its parent oc.Component_Cpu_Utilization
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertComponent_Cpu_Utilization_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Cpu_Utilization) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Max
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/cpu/utilization/state/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Cpu_Utilization_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Cpu_Utilization{}
	md, ok := oc.Lookup(t, n, "Component_Cpu_Utilization", goStruct, true, false)
	if ok {
		return convertComponent_Cpu_Utilization_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/cpu/utilization/state/max-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Cpu_Utilization_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/cpu/utilization/state/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Cpu_Utilization_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Cpu_Utilization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Cpu_Utilization", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Cpu_Utilization_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/cpu/utilization/state/max-time with a ONCE subscription.
func (n *Component_Cpu_Utilization_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization/state/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_Utilization_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Cpu_Utilization_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Cpu_Utilization{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Cpu_Utilization", gs, queryPath, true, false)
		return convertComponent_Cpu_Utilization_MaxTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization/state/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_Utilization_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Cpu_Utilization_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/cpu/utilization/state/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Cpu_Utilization_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/cpu/utilization/state/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/cpu/utilization/state/max-time to the batch object.
func (n *Component_Cpu_Utilization_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/cpu/utilization/state/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Cpu_Utilization_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/cpu/utilization/state/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Cpu_Utilization_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Cpu_Utilization_MaxTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/cpu/utilization/state/max-time to the batch object.
func (n *Component_Cpu_Utilization_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Cpu_Utilization_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_Cpu_Utilization
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Cpu_Utilization_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Cpu_Utilization) *oc.QualifiedUint64 {
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
