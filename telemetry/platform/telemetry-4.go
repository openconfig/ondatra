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

// Lookup fetches the value at /openconfig-platform/components/component/state/redundant-role with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_RedundantRolePath) Lookup(t testing.TB) *oc.QualifiedE_PlatformTypes_ComponentRedundantRole {
	t.Helper()
	goStruct := &oc.Component{}
	md, ok := oc.Lookup(t, n, "Component", goStruct, true, false)
	if ok {
		return convertComponent_RedundantRolePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/redundant-role with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_RedundantRolePath) Get(t testing.TB) oc.E_PlatformTypes_ComponentRedundantRole {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/redundant-role with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_RedundantRolePathAny) Lookup(t testing.TB) []*oc.QualifiedE_PlatformTypes_ComponentRedundantRole {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PlatformTypes_ComponentRedundantRole
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_RedundantRolePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/redundant-role with a ONCE subscription.
func (n *Component_RedundantRolePathAny) Get(t testing.TB) []oc.E_PlatformTypes_ComponentRedundantRole {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PlatformTypes_ComponentRedundantRole
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/redundant-role with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_RedundantRolePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_ComponentRedundantRole {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_ComponentRedundantRole{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_ComponentRedundantRole) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_RedundantRolePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_ComponentRedundantRole) bool) *oc.E_PlatformTypes_ComponentRedundantRoleWatcher {
	t.Helper()
	w := &oc.E_PlatformTypes_ComponentRedundantRoleWatcher{}
	gs := &oc.Component{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component", gs, queryPath, true, false)
		return convertComponent_RedundantRolePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PlatformTypes_ComponentRedundantRole)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/redundant-role with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_RedundantRolePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_ComponentRedundantRole) bool) *oc.E_PlatformTypes_ComponentRedundantRoleWatcher {
	t.Helper()
	return watch_Component_RedundantRolePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/redundant-role with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_RedundantRolePath) Await(t testing.TB, timeout time.Duration, val oc.E_PlatformTypes_ComponentRedundantRole) *oc.QualifiedE_PlatformTypes_ComponentRedundantRole {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PlatformTypes_ComponentRedundantRole) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/redundant-role failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/redundant-role to the batch object.
func (n *Component_RedundantRolePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/redundant-role with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_RedundantRolePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PlatformTypes_ComponentRedundantRole {
	t.Helper()
	c := &oc.CollectionE_PlatformTypes_ComponentRedundantRole{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PlatformTypes_ComponentRedundantRole) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/redundant-role with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_RedundantRolePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PlatformTypes_ComponentRedundantRole) bool) *oc.E_PlatformTypes_ComponentRedundantRoleWatcher {
	t.Helper()
	return watch_Component_RedundantRolePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/redundant-role to the batch object.
func (n *Component_RedundantRolePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_RedundantRolePath extracts the value of the leaf RedundantRole from its parent oc.Component
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PlatformTypes_ComponentRedundantRole.
func convertComponent_RedundantRolePath(t testing.TB, md *genutil.Metadata, parent *oc.Component) *oc.QualifiedE_PlatformTypes_ComponentRedundantRole {
	t.Helper()
	qv := &oc.QualifiedE_PlatformTypes_ComponentRedundantRole{
		Metadata: md,
	}
	val := parent.RedundantRole
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

// Lookup fetches the value at /openconfig-platform/components/component/storage with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_StoragePath) Lookup(t testing.TB) *oc.QualifiedComponent_Storage {
	t.Helper()
	goStruct := &oc.Component_Storage{}
	md, ok := oc.Lookup(t, n, "Component_Storage", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Storage{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/storage with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_StoragePath) Get(t testing.TB) *oc.Component_Storage {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/storage with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_StoragePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Storage {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Storage
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Storage{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Storage", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Storage{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/storage with a ONCE subscription.
func (n *Component_StoragePathAny) Get(t testing.TB) []*oc.Component_Storage {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Storage
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/storage with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_StoragePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Storage {
	t.Helper()
	c := &oc.CollectionComponent_Storage{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Storage) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Storage{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Storage)))
		return false
	})
	return c
}

func watch_Component_StoragePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Storage) bool) *oc.Component_StorageWatcher {
	t.Helper()
	w := &oc.Component_StorageWatcher{}
	gs := &oc.Component_Storage{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Storage", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Storage{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Storage)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/storage with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_StoragePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Storage) bool) *oc.Component_StorageWatcher {
	t.Helper()
	return watch_Component_StoragePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/storage with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_StoragePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Storage) *oc.QualifiedComponent_Storage {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Storage) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/storage failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/storage to the batch object.
func (n *Component_StoragePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/storage with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_StoragePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Storage {
	t.Helper()
	c := &oc.CollectionComponent_Storage{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Storage) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/storage with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_StoragePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Storage) bool) *oc.Component_StorageWatcher {
	t.Helper()
	return watch_Component_StoragePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/storage to the batch object.
func (n *Component_StoragePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_SubcomponentPath) Lookup(t testing.TB) *oc.QualifiedComponent_Subcomponent {
	t.Helper()
	goStruct := &oc.Component_Subcomponent{}
	md, ok := oc.Lookup(t, n, "Component_Subcomponent", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Subcomponent{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_SubcomponentPath) Get(t testing.TB) *oc.Component_Subcomponent {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_SubcomponentPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Subcomponent {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Subcomponent
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Subcomponent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Subcomponent", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Subcomponent{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent with a ONCE subscription.
func (n *Component_SubcomponentPathAny) Get(t testing.TB) []*oc.Component_Subcomponent {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Subcomponent
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/subcomponents/subcomponent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_SubcomponentPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Subcomponent {
	t.Helper()
	c := &oc.CollectionComponent_Subcomponent{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Subcomponent) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Subcomponent{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Subcomponent)))
		return false
	})
	return c
}

func watch_Component_SubcomponentPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Subcomponent) bool) *oc.Component_SubcomponentWatcher {
	t.Helper()
	w := &oc.Component_SubcomponentWatcher{}
	gs := &oc.Component_Subcomponent{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Subcomponent", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Subcomponent{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Subcomponent)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/subcomponents/subcomponent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_SubcomponentPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Subcomponent) bool) *oc.Component_SubcomponentWatcher {
	t.Helper()
	return watch_Component_SubcomponentPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/subcomponents/subcomponent with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_SubcomponentPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Subcomponent) *oc.QualifiedComponent_Subcomponent {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Subcomponent) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/subcomponents/subcomponent failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/subcomponents/subcomponent to the batch object.
func (n *Component_SubcomponentPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/subcomponents/subcomponent with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_SubcomponentPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Subcomponent {
	t.Helper()
	c := &oc.CollectionComponent_Subcomponent{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Subcomponent) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/subcomponents/subcomponent with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_SubcomponentPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Subcomponent) bool) *oc.Component_SubcomponentWatcher {
	t.Helper()
	return watch_Component_SubcomponentPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/subcomponents/subcomponent to the batch object.
func (n *Component_SubcomponentPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Subcomponent_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Subcomponent{}
	md, ok := oc.Lookup(t, n, "Component_Subcomponent", goStruct, true, false)
	if ok {
		return convertComponent_Subcomponent_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Subcomponent_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Subcomponent_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Subcomponent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Subcomponent", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Subcomponent_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a ONCE subscription.
func (n *Component_Subcomponent_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Subcomponent_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Subcomponent_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Component_Subcomponent{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Subcomponent", gs, queryPath, true, false)
		return convertComponent_Subcomponent_NamePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Subcomponent_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Subcomponent_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Subcomponent_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/subcomponents/subcomponent/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/subcomponents/subcomponent/state/name to the batch object.
func (n *Component_Subcomponent_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Subcomponent_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/subcomponents/subcomponent/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Subcomponent_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Component_Subcomponent_NamePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/subcomponents/subcomponent/state/name to the batch object.
func (n *Component_Subcomponent_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Subcomponent_NamePath extracts the value of the leaf Name from its parent oc.Component_Subcomponent
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Subcomponent_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Subcomponent) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_TemperaturePath) Lookup(t testing.TB) *oc.QualifiedComponent_Temperature {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Temperature{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_TemperaturePath) Get(t testing.TB) *oc.Component_Temperature {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_TemperaturePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Temperature {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Temperature
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Temperature{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature with a ONCE subscription.
func (n *Component_TemperaturePathAny) Get(t testing.TB) []*oc.Component_Temperature {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Temperature
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_TemperaturePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Temperature {
	t.Helper()
	c := &oc.CollectionComponent_Temperature{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Temperature) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Temperature{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Temperature)))
		return false
	})
	return c
}

func watch_Component_TemperaturePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Temperature) bool) *oc.Component_TemperatureWatcher {
	t.Helper()
	w := &oc.Component_TemperatureWatcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Temperature{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Temperature)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_TemperaturePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Temperature) bool) *oc.Component_TemperatureWatcher {
	t.Helper()
	return watch_Component_TemperaturePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_TemperaturePath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Temperature) *oc.QualifiedComponent_Temperature {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Temperature) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature to the batch object.
func (n *Component_TemperaturePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_TemperaturePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Temperature {
	t.Helper()
	c := &oc.CollectionComponent_Temperature{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Temperature) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_TemperaturePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Temperature) bool) *oc.Component_TemperatureWatcher {
	t.Helper()
	return watch_Component_TemperaturePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature to the batch object.
func (n *Component_TemperaturePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/alarm-severity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_AlarmSeverityPath) Lookup(t testing.TB) *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_AlarmSeverityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/alarm-severity with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_AlarmSeverityPath) Get(t testing.TB) oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_AlarmSeverityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_AlarmSeverityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a ONCE subscription.
func (n *Component_Temperature_AlarmSeverityPathAny) Get(t testing.TB) []oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmSeverityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	c := &oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_AlarmSeverityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	w := &oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_AlarmSeverityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmSeverityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	return watch_Component_Temperature_AlarmSeverityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/alarm-severity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_AlarmSeverityPath) Await(t testing.TB, timeout time.Duration, val oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/alarm-severity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-severity to the batch object.
func (n *Component_Temperature_AlarmSeverityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmSeverityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	c := &oc.CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmSeverityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) bool) *oc.E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher {
	t.Helper()
	return watch_Component_Temperature_AlarmSeverityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-severity to the batch object.
func (n *Component_Temperature_AlarmSeverityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_AlarmSeverityPath extracts the value of the leaf AlarmSeverity from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY.
func convertComponent_Temperature_AlarmSeverityPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	qv := &oc.QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY{
		Metadata: md,
	}
	val := parent.AlarmSeverity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/alarm-status with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_AlarmStatusPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_AlarmStatusPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/alarm-status with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_AlarmStatusPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/alarm-status with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_AlarmStatusPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_AlarmStatusPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/alarm-status with a ONCE subscription.
func (n *Component_Temperature_AlarmStatusPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmStatusPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_AlarmStatusPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_AlarmStatusPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmStatusPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Component_Temperature_AlarmStatusPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/alarm-status with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_AlarmStatusPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/alarm-status failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-status to the batch object.
func (n *Component_Temperature_AlarmStatusPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmStatusPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmStatusPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Component_Temperature_AlarmStatusPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-status to the batch object.
func (n *Component_Temperature_AlarmStatusPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_AlarmStatusPath extracts the value of the leaf AlarmStatus from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertComponent_Temperature_AlarmStatusPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.AlarmStatus
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/alarm-threshold with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_AlarmThresholdPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_AlarmThresholdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/alarm-threshold with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_AlarmThresholdPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_AlarmThresholdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_AlarmThresholdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a ONCE subscription.
func (n *Component_Temperature_AlarmThresholdPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmThresholdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_AlarmThresholdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_AlarmThresholdPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmThresholdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Component_Temperature_AlarmThresholdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_AlarmThresholdPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/alarm-threshold failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-threshold to the batch object.
func (n *Component_Temperature_AlarmThresholdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AlarmThresholdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/alarm-threshold with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AlarmThresholdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Component_Temperature_AlarmThresholdPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/alarm-threshold to the batch object.
func (n *Component_Temperature_AlarmThresholdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_AlarmThresholdPath extracts the value of the leaf AlarmThreshold from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertComponent_Temperature_AlarmThresholdPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.AlarmThreshold
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/avg with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/avg with a ONCE subscription.
func (n *Component_Temperature_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_AvgPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/avg to the batch object.
func (n *Component_Temperature_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_AvgPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/avg to the batch object.
func (n *Component_Temperature_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_AvgPath extracts the value of the leaf Avg from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Temperature_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/instant with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/instant with a ONCE subscription.
func (n *Component_Temperature_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_InstantPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/instant to the batch object.
func (n *Component_Temperature_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_InstantPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/instant to the batch object.
func (n *Component_Temperature_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_InstantPath extracts the value of the leaf Instant from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Temperature_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/interval with a ONCE subscription.
func (n *Component_Temperature_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_IntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/interval to the batch object.
func (n *Component_Temperature_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_IntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/interval to the batch object.
func (n *Component_Temperature_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_IntervalPath extracts the value of the leaf Interval from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Temperature_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/max with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/max with a ONCE subscription.
func (n *Component_Temperature_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_MaxPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/max to the batch object.
func (n *Component_Temperature_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_MaxPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/max to the batch object.
func (n *Component_Temperature_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_MaxPath extracts the value of the leaf Max from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Temperature_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/max-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/max-time with a ONCE subscription.
func (n *Component_Temperature_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_MaxTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/max-time to the batch object.
func (n *Component_Temperature_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_MaxTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/max-time to the batch object.
func (n *Component_Temperature_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Temperature_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_MinPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/min with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_MinPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/min with a ONCE subscription.
func (n *Component_Temperature_MinPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MinPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_MinPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_MinPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MinPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_MinPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/min with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_MinPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/min failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/min to the batch object.
func (n *Component_Temperature_MinPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/min with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MinPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/min with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MinPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_Component_Temperature_MinPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/min to the batch object.
func (n *Component_Temperature_MinPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_MinPath extracts the value of the leaf Min from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Temperature_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/state/temperature/min-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Temperature_MinTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Component_Temperature{}
	md, ok := oc.Lookup(t, n, "Component_Temperature", goStruct, true, false)
	if ok {
		return convertComponent_Temperature_MinTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/state/temperature/min-time with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Temperature_MinTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/state/temperature/min-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Temperature_MinTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Temperature{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Temperature", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertComponent_Temperature_MinTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/state/temperature/min-time with a ONCE subscription.
func (n *Component_Temperature_MinTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MinTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Component_Temperature_MinTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Component_Temperature{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Temperature", gs, queryPath, true, false)
		return convertComponent_Temperature_MinTimePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MinTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_MinTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/state/temperature/min-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_Temperature_MinTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/state/temperature/min-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/state/temperature/min-time to the batch object.
func (n *Component_Temperature_MinTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/state/temperature/min-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_Temperature_MinTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/state/temperature/min-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_Temperature_MinTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Component_Temperature_MinTimePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/state/temperature/min-time to the batch object.
func (n *Component_Temperature_MinTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertComponent_Temperature_MinTimePath extracts the value of the leaf MinTime from its parent oc.Component_Temperature
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertComponent_Temperature_MinTimePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Temperature) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_TransceiverPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, false, false)
	if ok {
		return (&oc.QualifiedComponent_Transceiver{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_TransceiverPath) Get(t testing.TB) *oc.Component_Transceiver {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_TransceiverPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver with a ONCE subscription.
func (n *Component_TransceiverPathAny) Get(t testing.TB) []*oc.Component_Transceiver {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_TransceiverPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedComponent_Transceiver{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Component_Transceiver)))
		return false
	})
	return c
}

func watch_Component_TransceiverPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver) bool) *oc.Component_TransceiverWatcher {
	t.Helper()
	w := &oc.Component_TransceiverWatcher{}
	gs := &oc.Component_Transceiver{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Component_Transceiver", gs, queryPath, false, false)
		return (&oc.QualifiedComponent_Transceiver{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedComponent_Transceiver)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_TransceiverPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver) bool) *oc.Component_TransceiverWatcher {
	t.Helper()
	return watch_Component_TransceiverPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-platform/components/component/transceiver with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Component_TransceiverPath) Await(t testing.TB, timeout time.Duration, val *oc.Component_Transceiver) *oc.QualifiedComponent_Transceiver {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedComponent_Transceiver) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-platform/components/component/transceiver failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-platform/components/component/transceiver to the batch object.
func (n *Component_TransceiverPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-platform/components/component/transceiver with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Component_TransceiverPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionComponent_Transceiver {
	t.Helper()
	c := &oc.CollectionComponent_Transceiver{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedComponent_Transceiver) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-platform/components/component/transceiver with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Component_TransceiverPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedComponent_Transceiver) bool) *oc.Component_TransceiverWatcher {
	t.Helper()
	return watch_Component_TransceiverPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-platform/components/component/transceiver to the batch object.
func (n *Component_TransceiverPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
